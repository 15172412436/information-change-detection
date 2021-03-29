package users

import (
	"github.com/jinzhu/gorm"
	"godeliver/models"
	"godeliver/pkg/e"
	"godeliver/pkg/logging"
	"godeliver/pkg/util"
)

type UserLogin struct {
	UID      uint `gorm:"primary_key"`
	Username string
	Phone    string
	Password string
	status   uint
}

// 验证码登录
//func codeLogin() {
//
//}

// AddUserLogin 添加用户账号 与 初始化个人信息
func AddUserLogin(userLogin map[string]interface{}) (uint, util.Error) {

	tx := models.DB.Begin()

	// 首先创建 user
	userID, err := addUser(tx, userLogin)
	if err != nil {
		tx.Rollback()
		logging.Error(err)
		return 0, util.ErrNewCode(e.ErrorUserInfoCreate)
	}

	loginInfo := UserLogin{
		UID: userID,
	}

	if userLogin["password"] != nil {
		loginInfo.Password = userLogin["password"].(string)
	}

	if userLogin["username"] != nil {
		loginInfo.Username = userLogin["username"].(string)
		goto InsertLogin
	}
	if userLogin["phone"] != nil {
		loginInfo.Phone = userLogin["phone"].(string)
		goto InsertLogin
	}

InsertLogin:

	if err := tx.Create(&loginInfo).Error; err != nil {
		tx.Rollback()
		logging.Error(err)
		return 0, util.ErrNewCode(e.ErrorUserLoginCreate)
	}
	tx.Commit()
	return userID, nil
}

// LoginUserLogin 采用密码方式登录
func LoginUserLogin(maps map[string]interface{}) (*UserLogin, util.Error) {
	var user UserLogin
	if err := models.DB.Where(maps).First(&user).Error; err != nil {
		logging.Error(err)
		return nil, util.ErrNewSql(err)
	}
	return &user, nil
}

// ExistUserLogin 判断是否存在此用户账号 返回uid
func ExistUserLogin(maps map[string]interface{}) (uint, util.Error) {
	var user UserLogin
	err := models.DB.Select("uid").Where(maps).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Error(err)
		return 0, util.ErrNewCode(e.ErrorUserGetLogin)
	}
	if user.UID > 0 {
		return user.UID, nil
	}
	return 0, nil
}

// UserLoginGetUserID 通过用户名 获取 用户ID
func UserLoginGetUserID(maps map[string]interface{}) (uint, util.Error) {
	var user UserLogin
	err := models.DB.Select("uid").Where(maps).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, util.ErrNewSql(err)
	}
	if user.UID < 1 { // 判断用户账号是否存在
		return 0, util.ErrNewCode(e.ErrorUserNameNotExist)
	}
	if user.UID < 1 { // 判断用户信息是否存在
		return 0, util.ErrNewCode(e.ErrorUserInfoEmpty)
	}

	return user.UID, nil
}
