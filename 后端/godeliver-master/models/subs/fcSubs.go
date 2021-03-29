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

type Fcdy struct {
	SubscriptionID string
	UID            uint
	Url            string
	Frequency      string
	SetTime        int
	Email          string
	Id             uint
	Subject        string
}

// todo 用户是否存在订阅 by uid
func ExistFcdyByUid(uid uint) (bool, error) {
	var fcdy Fcdy
	err := models.DB.Select("uid").Where("uid = ?", uid).First(&fcdy).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if fcdy.UID > 0 {
		return true, nil
	}
	return false, nil
}

//todo 是否存在分词订阅id
func ExistFcdyBySid(subscriptionId string) (bool, error) {
	var fcdy Fcdy
	//subscriptionId += "_fcdy"
	err := models.DB.Select("subscription_id").Where("subscription_id = ?", subscriptionId).First(&fcdy).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	sid := strings.Split(fcdy.SubscriptionID, "_")[0]
	subId, _ := strconv.Atoi(sid)
	if subId > 0 {
		return true, nil
	}
	return false, nil
}

// 获取分词订阅
//todo 返回用户所有的分词订阅
func GetFcdy(uid uint) ([]*Fcdy, util.Error) {
	var fcdy []*Fcdy
	// 不存在返回空列表
	exist, err := ExistFcdyByUid(uid)
	if err != nil {
		return fcdy, util.ErrNewErr(err)
	}
	if !exist {
		return fcdy, util.ErrNewCode(e.ErrorSubsFcdyNotExisted)
	}
	err = models.DB.Where("uid = ? ", uid).Find(&fcdy).Error
	if err != nil {
		logging.Error(err)
		return nil, util.ErrNewCode(e.ErrorSubsGetInfo)
	}
	return fcdy, nil
}

// 添加分词订阅 todo 还需要实现判断是否存在？
func AddFcdy(fcdyInfo map[string]interface{}) (string, error) {
	tx := models.DB.Begin()

	fcdy := Fcdy{
		UID:       fcdyInfo["uid"].(uint),
		Url:       fcdyInfo["content"].(string),
		Frequency: fcdyInfo["frequency"].(string),
		SetTime:   fcdyInfo["set_time"].(int),
		Email:     fcdyInfo["email"].(string),
		Subject:   fcdyInfo["subject"].(string),
	}
	// todo 获取分词订阅id `1_fcdy`
	//total := 0
	//
	//if err := tx.Model(Fcdy{}).Count(&total).Error; err != nil {
	//	tx.Rollback()
	//	logging.Error(err)
	//	return "0", err
	//}
	var lastFcdy []Fcdy
	if err := tx.Model(Fcdy{}).Find(&lastFcdy).Error; err != nil {
		logging.Error(err)
		tx.Rollback()
		return "0", err
	}
	var current string
	var currentNum int
	if len(lastFcdy) == 0 {
		current = "0"
		currentNum, _ = strconv.Atoi(current)
	} else {
		fmt.Println(lastFcdy[len(lastFcdy)-1].SubscriptionID)
		currentNum = int(lastFcdy[len(lastFcdy)-1].Id)
		fmt.Println("111", currentNum)
	}
	subID := strconv.Itoa(currentNum+1) + "_fcdy"
	fcdy.SubscriptionID = subID

	if err := tx.Create(&fcdy).Error; err != nil {
		tx.Rollback()
		logging.Error(err)
		return "0", err
	}

	tx.Commit()
	return fcdy.SubscriptionID, nil
}

//todo 修改分词订阅
func EditFcdy(sid string, fcdyInfo map[string]interface{}) error {
	tx := models.DB.Begin()

	// 检查fcdy是否存在
	exist, err := ExistFcdyBySid(sid)
	if err != nil {
		return util.ErrNewErr(err)
	}
	if !exist {
		return util.ErrNewCode(e.ErrorSubsFcdyNotExisted)
	}

	if err := tx.Model(&Fcdy{}).Where("subscription_id = ?", sid).Updates(fcdyInfo).Error; err != nil {
		tx.Rollback()
		logging.Error(err)
		return err
	}

	tx.Commit()
	return nil
}

//todo 删除分词订阅
func DeleteFcdy(sid string) error {
	tx := models.DB.Begin()
	exist, err := ExistFcdyBySid(sid)
	if err != nil {
		return util.ErrNewErr(err)
	}
	if !exist {
		return util.ErrNewCode(e.ErrorSubsFcdyNotExisted)
	}

	if err := tx.Where("subscription_id = ?", sid).Delete(Fcdy{}).Error; err != nil {
		tx.Rollback()
		logging.Error(err)
		return err
	}

	tx.Commit()
	return nil
}
