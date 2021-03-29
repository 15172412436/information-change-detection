package users

import (
	"github.com/jinzhu/gorm"
	"godeliver/models"
	"godeliver/pkg/e"
	"godeliver/pkg/logging"
	"godeliver/pkg/util"
)

type User struct {
	models.Model
	Nickname string
}

// 获取用户信息
func GetUser(uid uint) (*User, util.Error) {
	var user User
	err:=models.DB.Where("uid = ? ", uid).First(&user).Error
	if err != nil {
		logging.Error(err)
		return nil, util.ErrNewCode(e.ErrorUserGetInfo)
	}
	return &user, nil
}

// AddUser 新增用户信息
func addUser(tx *gorm.DB,userLogin map[string]interface{}) (uint, error) {
	var user User
	if userLogin["username"] != nil {
		user.Nickname = userLogin["username"].(string)
	}
	if userLogin["phone"] != nil {
		user.Nickname = userLogin["phone"].(string)
	}
	if err := tx.Create(&user).Error; err != nil {
		logging.Error(err)
		return 0, err
	}
	return user.UID, nil
}


// ExistUserByID 检查是否存在此用户
func ExistUserByID(id uint) (bool, util.Error) {
	var user User
	err := models.DB.Select("uid").Where("uid = ? ", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Error(err)
		return false, util.ErrNewCode(e.ErrorUserGetInfo)
	}
	if user.UID > 0 {
		return true, nil
	}
	return false, nil
}