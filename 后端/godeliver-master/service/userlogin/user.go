package userlogin

import (
	"encoding/json"
	"fmt"
	"godeliver/models/users"
	"godeliver/pkg/e"
	"godeliver/pkg/gredis"
	"godeliver/pkg/logging"
	"godeliver/pkg/sms"
	"godeliver/pkg/util"
	"godeliver/service/caches"
	"godeliver/service/subs_service"
)

// User 用户
type User struct {
	ID       uint
	UserName string
	Password string
	Code     string
}

// Register 注册用户
func (u *User) Register() (uint, util.Error) {
	maps := make(map[string]interface{})
	maps, err := u.validUserName(maps)
	if err != nil {
		return 0, err
	}

	// 加密
	password, err := util.Encrypt(u.Password)
	if err != nil {
		return 0, err
	}

	maps["password"] = password

	// 创建 用户信息 与 用户密码
	return users.AddUserLogin(maps)
}

// PwdLogin 登录用户
func (u *User) PwdLogin() (map[string]interface{}, util.Error) {

	user, err := u.getUserLoginInfo()
	if err != nil {
		return nil, err
	}

	inputPwd := u.Password
	hashPwd := user.Password

	// 比较 密码
	if err := util.Compare(inputPwd, hashPwd); err != nil {
		return nil, util.ErrNewCode(e.ErrorUserPwd)
	}

	if err := existUserInfo(user.UID); err != nil {
		return nil, err
	}

	// 生成token
	token, merr := util.GenerateToken(u.UserName)
	if merr != nil {
		logging.Error(merr)
		return nil, util.ErrNewCode(e.ERROR_AUTH_TOKEN)
	}
	retMap := make(map[string]interface{})
	retMap["token"] = token
	retMap["uid"] = user.UID

	subsService := subs_service.Subs{
		UID: user.UID,
	}

	subs, errs := subsService.GetAll()
	if errs != nil {
		logging.Error(errs)
	}
	retMap["subs"] = subs

	//resultsService := subs_service.Result{
	//	UID: user.UID,
	//}
	//results, errs := resultsService.GetAll()
	//if errs != nil {
	//	logging.Error(errs)
	//}
	//retMap["results"] = results
	return retMap, nil
}

// PhoneRegister 手机号注册
func (u *User) PhoneRegister() (uint, util.Error) {

	maps := map[string]interface{}{
		"phone": u.UserName,
	}
	fmt.Println(maps)
	// 创建 用户信息 与 用户密码
	return users.AddUserLogin(maps)

}

// GetUserInfo 获取用户信息
func (u *User) GetUserInfo() (*users.User, util.Error) {
	return users.GetUser(u.ID)

}

//todo SendCode 发送手机验证码
func SendCode(phone string) (string, util.Error) {
	var (
		code string
		err  error
	)
	code = GetCacheCode(phone)
	// 如果没有验证码,随机生成
	if code == "" {
		code = util.GetRandomCode()
	}

	cache := caches.Phone{Phone: phone}
	key := cache.GetPhoneCodeKey()

	SMSInfo := make(map[string]interface{}, 2)
	SMSInfo["phone"] = phone
	logging.Info(code)
	SMSInfo["code"] = code

	s := sms.SendSMS(SMSInfo)
	logging.Info(s)
	//	十分钟验证码缓存
	if err = gredis.Set(key, code, 600); err != nil {
		logging.Error(err)
		logging.Warn(caches.ErrorSet, err)
	}
	//if err != nil {
	//	logging.Error(err)
	//	return "", util.ErrNewCode(e.ErrorPhoneCodeNotValid)
	//}
	// 便于测试，code返回出去
	return code, nil
}

func (u *User) getUserLoginInfo() (*users.UserLogin, util.Error) {
	maps := make(map[string]interface{})
	maps, err := u.validUserName(maps)
	if err != nil {
		return nil, err
	}
	// 查询 用户登录信息
	user, err := users.LoginUserLogin(maps)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func existUserInfo(userID uint) util.Error {
	// 匹配成功  根据 userID 查询 用户信息
	if !(userID > 0) {
		return util.ErrNewCode(e.ErrorUserGetInfo)
	}

	exist, err := users.ExistUserByID(userID)
	if err != nil {
		return err
	}

	if !exist {
		return util.ErrNewCode(e.ErrorUserInfoEmpty)
	}
	return nil
}

// GetCacheCode 获取缓存的验证码
func GetCacheCode(phone string) string {
	cache := caches.Phone{Phone: phone}
	key := cache.GetPhoneCodeKey()
	logging.Info("use redis:", key)
	if !gredis.Exists(key) {
		return ""
	}
	var code string
	data, err := gredis.Get(key)
	if err != nil {
		logging.Warn(caches.ErrorGet, err)
		return ""
	}
	err = json.Unmarshal(data, &code)
	if err != nil {
		logging.Warn("GetCache Unmarshal 失败")
	}
	return code
}

// ExistByUserName 是否存在用户账号
func (u *User) ExistByUserName() (uint, util.Error) {
	maps := make(map[string]interface{})

	maps, err := u.validUserName(maps)
	if err != nil {
		return 0, err
	}
	return users.ExistUserLogin(maps)
}

// UserLoginGetUserID 返回用户ID
func (u *User) UserLoginGetUserID() (uint, util.Error) {
	maps := make(map[string]interface{})

	maps, err := u.validUserName(maps)
	if err != nil {
		return 0, err
	}
	return users.UserLoginGetUserID(maps)
}

//todo 验证 用户名类型
func (u *User) validUserName(maps map[string]interface{}) (map[string]interface{}, util.Error) {
	//if util.ValidEmail(u.UserName) {
	//	maps["login_email"] = u.UserName
	//	return maps, nil
	//}
	// 如果是手机号
	if util.RegPhone(u.UserName) {
		maps["phone"] = u.UserName
		return maps, nil
	}
	if util.RegUserName(u.UserName) {
		maps["username"] = u.UserName
		return maps, nil
	}
	return nil, util.ErrNewCode(e.ErrorUserRegName)
}
