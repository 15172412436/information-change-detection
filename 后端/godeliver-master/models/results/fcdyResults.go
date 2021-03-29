package results

import (
	"github.com/jinzhu/gorm"
	"godeliver/models"
	"godeliver/pkg/e"
	"godeliver/pkg/logging"
	"godeliver/pkg/util"
	"strconv"
	"strings"
)

type FcdyResult struct {
	Id             uint
	SubscriptionId string `gorm:"Column:subscriptionid"`
	WordFrequency  string `gorm:"Column:wordfrequency"`
	ReadMark       string `gorm:"Column:readmark"`
}

// 判断是否存在ID的结果
func ExistFcdyResultById(id uint) (bool, error) {
	var fcResult FcdyResult

	err := models.DB.Select("id").Where("id = ?", id).First(&fcResult).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if fcResult.Id > 0 {
		return true, nil
	}
	return false, nil
}

// 判断是否存在该订阅的结果
func ExistFcdyResultBySid(subscriptionId string) (bool, error) {
	var fcResult FcdyResult

	err := models.DB.Select("subscriptionid").Where("subscriptionid = ?", subscriptionId).First(&fcResult).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	sid := strings.Split(fcResult.SubscriptionId, "_")[0]
	subId, _ := strconv.Atoi(sid)
	if subId > 0 {
		return true, nil
	}
	return false, nil
}

// 获取 ：判断是否已读
// 获取所有订阅结果 判断是否已读，返回未读的结果。
func GetUnreadFcdyResult(subscriptionId string) ([]*FcdyResult, util.Error) {
	var fcdyRet []*FcdyResult
	// 不存在返回空列表
	exist, err := ExistFcdyResultBySid(subscriptionId)
	if err != nil {
		return fcdyRet, util.ErrNewErr(err)
	}
	if !exist {
		return fcdyRet, util.ErrNewCode(e.ErrorResultNotExisted)
	}
	// 查询 满足的订阅id 且 未阅读的选项
	err = models.DB.Where("subscriptionid = ? AND readmark = ? ", subscriptionId, "False").Find(&fcdyRet).Error
	if err != nil {
		logging.Error(err)
		return nil, util.ErrNewCode(e.ErrorResultGetInfo)
	}

	return fcdyRet, nil

}

// 获取所有订阅结果 判断是否已读，返回已读的结果。
func GetReadFcdyResult(subscriptionId string) ([]*FcdyResult, util.Error) {
	var fcdyRet []*FcdyResult
	// 不存在返回空列表
	exist, err := ExistFcdyResultBySid(subscriptionId)
	if err != nil {
		return fcdyRet, util.ErrNewErr(err)
	}
	if !exist {
		return fcdyRet, util.ErrNewCode(e.ErrorResultNotExisted)
	}
	// 查询 满足的订阅id 且 未阅读的选项
	err = models.DB.Where("subscriptionid = ? AND readmark = ? ", subscriptionId, "True").Find(&fcdyRet).Error
	if err != nil {
		logging.Error(err)
		return nil, util.ErrNewCode(e.ErrorResultGetInfo)
	}

	return fcdyRet, nil

}

func GetAllFcdyResult(subscriptionId string) ([]*FcdyResult, util.Error) {
	var fcdyRet []*FcdyResult
	// 不存在返回空列表
	exist, err := ExistFcdyResultBySid(subscriptionId)
	if err != nil {
		return fcdyRet, util.ErrNewErr(err)
	}
	if !exist {
		return fcdyRet, util.ErrNewCode(e.ErrorResultNotExisted)
	}
	// 查询 满足的订阅id 且 未阅读的选项
	err = models.DB.Where("subscriptionid = ? ", subscriptionId).Find(&fcdyRet).Error
	if err != nil {
		logging.Error(err)
		return nil, util.ErrNewCode(e.ErrorResultGetInfo)
	}
	return fcdyRet, nil
}

// 删除

// 增加 获取到新的内容添加到数据库

// todo 修改未读已读标签 (单个修改）
// 修改 主要修改 已读未读标签
func EditFcdyReadMark(id uint, mark string) error {
	tx := models.DB.Begin()

	// 检查 结果是否存在
	exist, err := ExistFcdyResultById(id)
	if err != nil {
		return util.ErrNewErr(err)
	}
	if !exist {
		return util.ErrNewCode(e.ErrorResultNotExisted)
	}
	logging.Info(mark)
	if mark == "false" || mark == "False" {
		if err := tx.Model(&FcdyResult{}).Where("id = ?", id).Update("readmark", "False").Error; err != nil {
			tx.Rollback()
			logging.Error(err)
			return err
		}
	} else {
		if err := tx.Model(&FcdyResult{}).Where("id = ?", id).Update("readmark", "True").Error; err != nil {
			tx.Rollback()
			logging.Error(err)
			return err
		}
	}
	tx.Commit()
	return nil
}

// 删除 传入订阅id 删除对应的所有结果
func DeleteFcdyResult(sid string) error {
	tx := models.DB.Begin()
	exist, err := ExistFcdyResultBySid(sid)
	if err != nil {
		return util.ErrNewErr(err)
	}
	if !exist {
		return util.ErrNewCode(e.ErrorResultNotExisted)
	}

	if err := tx.Where("subscriptionid = ?", sid).Delete(FcdyResult{}).Error; err != nil {
		tx.Rollback()
		logging.Error(err)
		return err
	}

	tx.Commit()
	return nil
}
