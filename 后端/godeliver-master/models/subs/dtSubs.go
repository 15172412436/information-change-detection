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

// 动态订阅
type Dtdy struct {
	Id             uint
	SubscriptionID string
	UID            uint
	Url            string
	Frequency      string
	SetTime        int
	Email          string
	Subject        string
	Fanwei         string
}

// 是否存在订阅id
func ExistDtdyBySid(subscriptionId string) (bool, error) {
	var dtdy Dtdy
	//subscriptionId += "_dtdy"
	err := models.DB.Select("subscription_id").Where("subscription_id = ?", subscriptionId).First(&dtdy).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	sid := strings.Split(dtdy.SubscriptionID, "_")[0]
	subId, _ := strconv.Atoi(sid)
	if subId > 0 {
		return true, nil
	}
	return false, nil
}

// 用户是否存在订阅
func ExistDtdyByUid(uid uint) (bool, error) {
	var dtdy Dtdy
	err := models.DB.Select("uid").Where("uid = ?", uid).First(&dtdy).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if dtdy.UID > 0 {
		return true, nil
	}
	return false, nil
}

// 获取用户所有的动态订阅
func GetDtdy(uid uint) ([]*Dtdy, util.Error) {
	var dtdy []*Dtdy
	// 不存在返回空列表
	exist, err := ExistDtdyByUid(uid)
	if err != nil {
		return dtdy, util.ErrNewErr(err)
	}
	if !exist {
		return dtdy, util.ErrNewCode(e.ErrorSubsDtdyNotExisted)
	}
	err = models.DB.Where("uid = ? ", uid).Find(&dtdy).Error
	if err != nil {
		logging.Error(err)
		return nil, util.ErrNewCode(e.ErrorSubsGetInfo)
	}
	return dtdy, nil
}

// 添加动态订阅
func AddDtdy(DtdyInfo map[string]interface{}) (string, error) {
	tx := models.DB.Begin()

	dtdy := Dtdy{
		UID:       DtdyInfo["uid"].(uint),
		Url:       DtdyInfo["content"].(string),
		Frequency: DtdyInfo["frequency"].(string),
		SetTime:   DtdyInfo["set_time"].(int),
		Email:     DtdyInfo["email"].(string),
		Subject:   DtdyInfo["subject"].(string),
		Fanwei:    DtdyInfo["fanwei"].(string),
	}
	// todo 获取分词订阅id `1_dtdy`
	//total := 0
	//
	//if err := tx.Model(Dtdy{}).Count(&total).Error; err != nil {
	//	tx.Rollback()
	//	logging.Error(err)
	//	return "0", err
	//}
	var lastDtdy []Dtdy
	if err := tx.Model(Dtdy{}).Find(&lastDtdy).Error; err != nil {
		logging.Error(err)
		tx.Rollback()
		return "0", err
	}
	var current string
	var currentNum int
	if len(lastDtdy) == 0 {
		current = "0"
		currentNum, _ = strconv.Atoi(current)
	} else {
		fmt.Println(lastDtdy[len(lastDtdy)-1].SubscriptionID)
		currentNum = int(lastDtdy[len(lastDtdy)-1].Id)
		fmt.Println("111", currentNum)
	}
	subID := strconv.Itoa(currentNum+1) + "_dtdy"
	dtdy.SubscriptionID = subID

	if err := tx.Create(&dtdy).Error; err != nil {
		tx.Rollback()
		logging.Error(err)
		return "0", err
	}

	tx.Commit()
	return dtdy.SubscriptionID, nil
}

// 删除动态订阅
func DeleteDtdy(sid string) error {
	tx := models.DB.Begin()
	exist, err := ExistDtdyBySid(sid)
	if err != nil {
		return util.ErrNewErr(err)
	}
	if !exist {
		return util.ErrNewCode(e.ErrorSubsDtdyNotExisted)
	}

	if err := tx.Where("subscription_id = ?", sid).Delete(Dtdy{}).Error; err != nil {
		tx.Rollback()
		logging.Error(err)
		return err
	}

	tx.Commit()
	return nil
}

// 修改动态订阅
func EditDtdy(sid string, DtdyInfo map[string]interface{}) error {
	tx := models.DB.Begin()

	// 检查Dtdy是否存在
	exist, err := ExistDtdyBySid(sid)
	if err != nil {
		return util.ErrNewErr(err)
	}
	if !exist {
		return util.ErrNewCode(e.ErrorSubsDtdyNotExisted)
	}

	if err := tx.Model(&Dtdy{}).Where("subscription_id = ?", sid).Updates(DtdyInfo).Error; err != nil {
		tx.Rollback()
		logging.Error(err)
		return err
	}

	tx.Commit()
	return nil
}
