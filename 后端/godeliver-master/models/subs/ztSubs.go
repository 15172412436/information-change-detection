package subs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"godeliver/models"
	"godeliver/pkg/e"
	"godeliver/pkg/logging"
	"godeliver/pkg/util"
	"strconv"
	"strings"
)

// ztdy订阅
type Ztdy struct {
	SubscriptionID string
	UID            uint
	Subject        string
	Frequency      string
	SetTime        int
	Email          string
	Keyword        string
	Id             uint
	ReturnType     string
}

// 是否存在订阅id
func ExistZtdyBySid(subscriptionId string) (bool, error) {
	var ztdy Ztdy
	//subscriptionId += "_ztdy"
	err := models.DB.Select("subscription_id").Where("subscription_id = ?", subscriptionId).First(&ztdy).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	sid := strings.Split(ztdy.SubscriptionID, "_")[0]
	subId, _ := strconv.Atoi(sid)
	if subId > 0 {
		return true, nil
	}
	return false, nil
}

// 用户是否存在订阅
func ExistZtdyByUid(uid uint) (bool, error) {
	var ztdy Ztdy
	err := models.DB.Select("uid").Where("uid = ?", uid).First(&ztdy).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if ztdy.UID > 0 {
		return true, nil
	}
	return false, nil
}

// 获取用户所有的动态订阅
func GetZtdy(uid uint) ([]*Ztdy, util.Error) {
	var ztdy []*Ztdy
	// 不存在返回空列表
	exist, err := ExistZtdyByUid(uid)
	if err != nil {
		return ztdy, util.ErrNewErr(err)
	}
	if !exist {
		return ztdy, util.ErrNewCode(e.ErrorSubsZtdyNotExisted)
	}
	err = models.DB.Where("uid = ? ", uid).Find(&ztdy).Error
	if err != nil {
		logging.Error(err)
		return nil, util.ErrNewCode(e.ErrorSubsGetInfo)
	}
	return ztdy, nil
}

// 添加动态订阅
func AddZtdy(ZtdyInfo map[string]interface{}) (string, error) {
	tx := models.DB.Begin()

	ztdy := Ztdy{
		UID:        ZtdyInfo["uid"].(uint),
		Subject:    ZtdyInfo["content"].(string),
		Frequency:  ZtdyInfo["frequency"].(string),
		SetTime:    ZtdyInfo["set_time"].(int),
		Email:      ZtdyInfo["email"].(string),
		Keyword:    ZtdyInfo["keyword"].(string),
		ReturnType: ZtdyInfo["return_type"].(string),
	}
	// todo 获取分词订阅id `1_ztdy`
	//total := 0
	//
	//if err := tx.Model(Ztdy{}).Count(&total).Error; err != nil {
	//	tx.Rollback()
	//	logging.Error(err)
	//	return "0", err
	//}
	var lastZtdy []Ztdy
	if err := tx.Model(Ztdy{}).Find(&lastZtdy).Error; err != nil {
		logging.Error(err)
		tx.Rollback()
		return "0", err
	}
	var current string
	var currentNum int
	if len(lastZtdy) == 0 {
		current = "0"
		currentNum, _ = strconv.Atoi(current)
	} else {
		fmt.Println(lastZtdy[len(lastZtdy)-1].SubscriptionID)
		currentNum = int(lastZtdy[len(lastZtdy)-1].Id)
		fmt.Println("111", currentNum)
	}
	subID := strconv.Itoa(currentNum+1) + "_ztdy"
	ztdy.SubscriptionID = subID

	if err := tx.Create(&ztdy).Error; err != nil {
		tx.Rollback()
		logging.Error(err)
		return "0", err
	}

	tx.Commit()
	return ztdy.SubscriptionID, nil
}

// 删除动态订阅
func DeleteZtdy(sid string) error {
	tx := models.DB.Begin()
	exist, err := ExistZtdyBySid(sid)
	if err != nil {
		return util.ErrNewErr(err)
	}
	if !exist {
		return util.ErrNewCode(e.ErrorSubsZtdyNotExisted)
	}

	if err := tx.Where("subscription_id = ?", sid).Delete(Ztdy{}).Error; err != nil {
		tx.Rollback()
		logging.Error(err)
		return err
	}

	tx.Commit()
	return nil
}

// 修改动态订阅
func EditZtdy(sid string, ZtdyInfo map[string]interface{}) error {
	tx := models.DB.Begin()

	// 检查Ztdy是否存在
	exist, err := ExistZtdyBySid(sid)
	if err != nil {
		return util.ErrNewErr(err)
	}
	if !exist {
		return util.ErrNewCode(e.ErrorSubsZtdyNotExisted)
	}

	if err := tx.Model(&Ztdy{}).Where("subscription_id = ?", sid).Updates(ZtdyInfo).Error; err != nil {
		tx.Rollback()
		logging.Error(err)
		return err
	}

	tx.Commit()
	return nil
}
