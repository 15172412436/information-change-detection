package results

import (
	"github.com/jinzhu/gorm"
	"godeliver/models"
	"godeliver/pkg/e"
	"godeliver/pkg/logging"
	"godeliver/pkg/util"
	"strconv"
	"strings"
	"time"
)

type DtdyResult struct {
	Id             uint
	SubscriptionID string
	Url            string
	ExecTime       time.Time
	ReadMark       string
}

// 判断是否存在ID的结果
func ExistDtdyResultById(id uint) (bool, error) {
	var dtResult DtdyResult

	err := models.DB.Select("id").Where("id = ?", id).First(&dtResult).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if dtResult.Id > 0 {
		return true, nil
	}
	return false, nil
}

// 判断是否存在该订阅的结果
func ExistDtdyResultBySid(subscriptionId string) (bool, error) {
	var dtResult DtdyResult

	err := models.DB.Select("subscription_id").Where("subscription_id = ?", subscriptionId).First(&dtResult).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	sid := strings.Split(dtResult.SubscriptionID, "_")[0]
	subId, _ := strconv.Atoi(sid)
	if subId > 0 {
		return true, nil
	}
	return false, nil
}

// 获取所有订阅结果 判断是否已读，返回未读的结果。
func GetUnreadDtdyResult(subscriptionId string) ([]*DtdyResult, util.Error) {
	var dtdyRet []*DtdyResult
	// 不存在返回空列表
	exist, err := ExistDtdyResultBySid(subscriptionId)
	if err != nil {
		return dtdyRet, util.ErrNewErr(err)
	}
	if !exist {
		return dtdyRet, util.ErrNewCode(e.ErrorResultNotExisted)
	}
	// 查询 满足的订阅id 且 未阅读的选项
	err = models.DB.Where("subscription_id = ? AND read_mark = ? ", subscriptionId, "False").Find(&dtdyRet).Error
	if err != nil {
		logging.Error(err)
		return nil, util.ErrNewCode(e.ErrorResultGetInfo)
	}

	return dtdyRet, nil

}

// 获取所有订阅结果 判断是否已读，返回未读的结果。
func GetReadDtdyResult(subscriptionId string) ([]*DtdyResult, util.Error) {
	var dtdyRet []*DtdyResult
	// 不存在返回空列表
	exist, err := ExistDtdyResultBySid(subscriptionId)
	if err != nil {
		return dtdyRet, util.ErrNewErr(err)
	}
	if !exist {
		return dtdyRet, util.ErrNewCode(e.ErrorResultNotExisted)
	}
	// 查询 满足的订阅id 且 未阅读的选项
	err = models.DB.Where("subscription_id = ? AND read_mark = ? ", subscriptionId, "True").Find(&dtdyRet).Error
	if err != nil {
		logging.Error(err)
		return nil, util.ErrNewCode(e.ErrorResultGetInfo)
	}

	return dtdyRet, nil

}

// 获取所有订阅结果 返回所有订阅的结果
func GetAllDtdyResult(subscriptionId string) ([]*DtdyResult, util.Error) {
	var dtdyRet []*DtdyResult
	// 不存在返回空列表
	exist, err := ExistDtdyResultBySid(subscriptionId)
	if err != nil {
		return dtdyRet, util.ErrNewErr(err)
	}
	if !exist {
		return dtdyRet, util.ErrNewCode(e.ErrorResultNotExisted)
	}
	// 查询 满足的订阅id 且 未阅读的选项
	err = models.DB.Where("subscription_id = ? ", subscriptionId).Find(&dtdyRet).Error
	if err != nil {
		logging.Error(err)
		return nil, util.ErrNewCode(e.ErrorResultGetInfo)
	}
	return dtdyRet, nil
}

// todo 添加多个
// 增加 添加新的结果到数据库
func AddDtdyResult(resultInfo map[string]interface{}) (string, error) {
	tx := models.DB.Begin()

	dtdyResult := DtdyResult{
		SubscriptionID: resultInfo["subscription_id"].(string),
		Url:            resultInfo["url"].(string),
		ExecTime:       resultInfo["exec_time"].(time.Time),
		ReadMark:       resultInfo["read_mark"].(string),
	}

	// 单个添加
	if err := tx.Create(&dtdyResult).Error; err != nil {
		tx.Rollback()
		logging.Error(err)
		return "0", err
	}
	tx.Commit()
	return "1", nil
}

// todo 修改未读已读标签 (单个修改）
// 修改 主要修改 已读未读标签
func EditDtdyReadMark(id uint, mark string) error {
	tx := models.DB.Begin()

	// 检查 结果是否存在
	exist, err := ExistDtdyResultById(id)
	if err != nil {
		return util.ErrNewErr(err)
	}
	if !exist {
		return util.ErrNewCode(e.ErrorResultNotExisted)
	}
	if mark == "false" || mark == "False" {
		if err := tx.Model(&DtdyResult{}).Where("id = ?", id).Update("read_mark", "False").Error; err != nil {
			tx.Rollback()
			logging.Error(err)
			return err
		}
	} else {
		if err := tx.Model(&DtdyResult{}).Where("id = ?", id).Update("read_mark", "True").Error; err != nil {
			tx.Rollback()
			logging.Error(err)
			return err
		}
	}
	tx.Commit()
	return nil
}

// 删除 传入订阅id 删除对应的所有结果
func DeleteDtdyResult(sid string) error {
	tx := models.DB.Begin()
	exist, err := ExistDtdyResultBySid(sid)
	if err != nil {
		return util.ErrNewErr(err)
	}
	if !exist {
		return util.ErrNewCode(e.ErrorResultNotExisted)
	}

	if err := tx.Where("subscription_id = ?", sid).Delete(DtdyResult{}).Error; err != nil {
		tx.Rollback()
		logging.Error(err)
		return err
	}

	tx.Commit()
	return nil
}
