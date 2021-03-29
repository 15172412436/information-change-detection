package results

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"godeliver/models"
	"godeliver/pkg/e"
	"godeliver/pkg/logging"
	"godeliver/pkg/util"
	"strconv"
	"strings"
	"time"
)

type ZtdyResult struct {
	Id             uint
	SubscriptionId string    `gorm:"Column:subscriptionid"`
	UId            string    `gorm:"uid"`
	ExecTime       time.Time `gorm:"Column:exectime"`
	Title          string    `gorm:"title"`
	Url            string    `gorm:"url"`
	Content        string    `gorm:"content"`
	ReadMark       string    `gorm:"Column:readmark"`
}

// 判断是否存在ID的结果
func ExistZtdyResultById(id uint) (bool, error) {
	var ztResult ZtdyResult

	err := models.DB.Select("id").Where("id = ?", id).First(&ztResult).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if ztResult.Id > 0 {
		return true, nil
	}
	return false, nil
}

// 判断是否存在该订阅的结果
func ExistZtdyResultBySid(subscriptionId string) (bool, error) {
	var ztResult ZtdyResult

	err := models.DB.Select("subscriptionid").Where("subscriptionid = ?", subscriptionId).First(&ztResult).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	sid := strings.Split(ztResult.SubscriptionId, "_")[0]
	subId, _ := strconv.Atoi(sid)
	fmt.Println(subId, sid, ztResult)
	if subId > 0 {
		return true, nil
	}

	return false, nil
}

// 获取 ：判断是否已读
// 获取所有订阅结果 判断是否已读，返回未读的结果。
func GetUnreadZtdyResult(subscriptionId string) ([]*ZtdyResult, util.Error) {
	var ztdyRet []*ZtdyResult
	// 不存在返回空列表
	exist, err := ExistZtdyResultBySid(subscriptionId)
	if err != nil {
		return ztdyRet, util.ErrNewErr(err)
	}
	if !exist {
		return ztdyRet, util.ErrNewCode(e.ErrorResultNotExisted)
	}
	// 查询 满足的订阅id 且 未阅读的选项
	err = models.DB.Where("subscriptionid = ? AND readmark = ? ", subscriptionId, "False").Find(&ztdyRet).Error
	if err != nil {
		logging.Error(err)
		return nil, util.ErrNewCode(e.ErrorResultGetInfo)
	}

	return ztdyRet, nil

}

// 获取所有订阅结果 判断是否已读，返回未读的结果。
func GetReadZtdyResult(subscriptionId string) ([]*ZtdyResult, util.Error) {
	var ztdyRet []*ZtdyResult
	// 不存在返回空列表
	exist, err := ExistZtdyResultBySid(subscriptionId)
	if err != nil {
		return ztdyRet, util.ErrNewErr(err)
	}
	if !exist {
		return ztdyRet, util.ErrNewCode(e.ErrorResultNotExisted)
	}
	// 查询 满足的订阅id 且 未阅读的选项
	err = models.DB.Where("subscriptionid = ? AND readmark = ? ", subscriptionId, "True").Find(&ztdyRet).Error
	if err != nil {
		logging.Error(err)
		return nil, util.ErrNewCode(e.ErrorResultGetInfo)
	}
	return ztdyRet, nil
}

// 获取所有订阅结果 返回所有订阅的结果
func GetAllZtdyResult(subscriptionId string) ([]*ZtdyResult, util.Error) {
	var ztdyRet []*ZtdyResult
	// 不存在返回空列表
	exist, err := ExistZtdyResultBySid(subscriptionId)
	if err != nil {
		return ztdyRet, util.ErrNewErr(err)
	}
	if !exist {
		return ztdyRet, util.ErrNewCode(e.ErrorResultNotExisted)
	}
	// 查询 满足的订阅id 且 未阅读的选项
	err = models.DB.Where("subscriptionid = ? ", subscriptionId).Find(&ztdyRet).Error
	if err != nil {
		logging.Error(err)
		return nil, util.ErrNewCode(e.ErrorResultGetInfo)
	}
	return ztdyRet, nil
}

// 删除

// 增加 获取到新的内容添加到数据库

// todo 修改未读已读标签 (单个修改）
// 修改 主要修改 已读未读标签
func EditZtdyReadMark(id uint, mark string) error {
	tx := models.DB.Begin()

	// 检查 结果是否存在
	exist, err := ExistZtdyResultById(id)
	if err != nil {
		return util.ErrNewErr(err)
	}
	if !exist {
		return util.ErrNewCode(e.ErrorResultNotExisted)
	}

	if mark == "false" || mark == "False" {
		if err := tx.Model(&ZtdyResult{}).Where("id = ?", id).Update("readmark", "False").Error; err != nil {
			tx.Rollback()
			logging.Error(err)
			return err
		}
	} else {
		if err := tx.Model(&ZtdyResult{}).Where("id = ?", id).Update("readmark", "True").Error; err != nil {
			tx.Rollback()
			logging.Error(err)
			return err
		}
	}

	tx.Commit()
	return nil
}

// 删除 传入订阅id 删除对应的所有结果
func DeleteZtdyResult(sid string) error {
	tx := models.DB.Begin()
	exist, err := ExistZtdyResultBySid(sid)
	if err != nil {
		return util.ErrNewErr(err)
	}
	if !exist {
		return util.ErrNewCode(e.ErrorResultNotExisted)
	}

	if err := tx.Where("subscriptionid = ?", sid).Delete(ZtdyResult{}).Error; err != nil {
		tx.Rollback()
		logging.Error(err)
		return err
	}

	tx.Commit()
	return nil
}
