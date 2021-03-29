package api

import (
	"github.com/gin-gonic/gin"
	"godeliver/pkg/app"
	"godeliver/pkg/e"
	"godeliver/pkg/logging"
	"godeliver/pkg/util"
	"godeliver/service/subs_service"
	"godeliver/service/userlogin"
)

type auth struct {
	// UserName 用户名
	Username string `json:"username" example:"zhangsan" validate:"required,gte=5,lte=30"`
	// PassWord 密码
	Password string `json:"password" example:"zhangsan" validate:"required,gte=5,lte=30"`
}

type token struct {
	Token string `json:"token"`
}

// Register 账号密码注册
// @Summary 账号密码注册
// @accept application/x-www-form-urlencoded
// @Tags auth
// @Produce  json
// @Param auth body api.auth true "账号密码登录/注册"
// @Success 200 {string} token "{"code":200,"data":{"token":token},"msg":"ok"}"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	appG := app.GetGin(c)
	var mAuth auth
	// 解析 body json 数据到实体类
	if err := c.ShouldBindJSON(&mAuth); err != nil {
		appG.ResponseFailError(util.ErrNewErr(err))
		return
	}
	// 验证
	validate := util.GetValidate()
	err := validate.Struct(mAuth)
	if err != nil {
		logging.Error(err)
		appG.ResponseFailErrCode(e.INVALID_PARAMS)
		return
	}

	userService := userlogin.User{UserName: mAuth.Username, Password: mAuth.Password}
	// 判断是否存在
	uid, err := userService.ExistByUserName()
	if err != nil {
		appG.ResponseFailError(util.ErrNewCode(e.ErrorUserNameNotExist))
		return
	}

	if uid != 0 {
		appG.ResponseFailError(util.ErrNewCode(e.ErrorUserNameExist))
		return
	}

	// 注册
	if id, err := userService.Register(); err != nil {
		uid = id
		appG.ResponseFailError(err)
		return
	}
	var t token
	// 注册成功之后 make token
	t.Token, err = util.GenerateToken(mAuth.Username)
	if err != nil {
		logging.Error(err)
		appG.ResponseFailError(util.ErrNewCode(e.ERROR_AUTH_CHECK_TOKEN_FAIL))
		return
	}

	retMap := make(map[string]interface{})
	retMap["token"] = t.Token
	retMap["uid"] = uid

	appG.ResponseSuc(retMap)

}

// Login 账号密码登录
// @Summary 账号密码登录
// @accept application/x-www-form-urlencoded
// @Tags auth
// @Produce  json
// @Param auth body api.auth true "用户信息"
// @Success 200 {string} token "{"code":200,"data":{"token":token},"msg":"ok"}"
// @Failure 200 {object} token "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	appG := app.GetGin(c)
	var mAuth auth
	// 解析 body json 数据到实体类
	if err := c.ShouldBindJSON(&mAuth); err != nil {
		appG.ResponseFailError(util.ErrNewErr(err))
		return
	}

	// 验证
	validate := util.GetValidate()
	err := validate.Struct(mAuth)
	if err != nil {
		logging.Error(err)
		appG.ResponseFailErrCode(e.INVALID_PARAMS)
		return
	}

	// 传值
	userService := userlogin.User{
		UserName: mAuth.Username,
		Password: mAuth.Password,
	}

	userMap, errA := userService.PwdLogin()
	// 登录查询成功
	if errA != nil {
		appG.ResponseFailError(errA)
		return
	}
	appG.ResponseSuc(userMap)
}

type phone struct {
	// Phone 手机号
	Phone string `json:"phone" example:"13938738804" validate:"required"`
	// Code 手机号验证码
	Code string `json:"code" example:"123456" validate:"required"`
}

// PhoneLogin 手机号快速登陆/注册
// @Summary 手机号快速登陆/注册
// @Document 如果登录手机号未注册,则自动注册再登录
// @accept application/x-www-form-urlencoded
// @Tags auth
// @Produce  json
// @Param auth body api.phone true "手机号快速登录/注册"
// @Success 200 {string} token "{"code":200,"data":{"token":token},"msg":"ok"}"
// @Router /auth/phoneLogin [post]
func PhoneLogin(c *gin.Context) {
	appG := app.GetGin(c)
	var mAuth phone

	// 解析 body json 数据到实体类
	if err := c.ShouldBindJSON(&mAuth); err != nil {
		appG.ResponseFailError(util.ErrNewErr(err))
		return
	}
	// 验证
	validate := util.GetValidate()
	err := validate.Struct(mAuth)
	if err != nil {
		appG.ResponseFailErrCode(e.INVALID_PARAMS)
		return
	}

	if !util.RegPhone(mAuth.Phone) {
		appG.ResponseFailErrCode(e.ErrorPhoneNotValid)
		return
	}

	// 验证验证码
	code := userlogin.GetCacheCode(mAuth.Phone)
	if code == "" {
		appG.ResponseFailErrCode(e.ErrorPhoneCodeExpired)
		return
	}
	if mAuth.Code != code {
		appG.ResponseFailErrCode(e.ErrorPhoneCodeNotValid)
		return
	}

	userService := userlogin.User{UserName: mAuth.Phone, Code: mAuth.Code}

	// 判断是否存在
	uid, merr := userService.ExistByUserName()
	if merr != nil {
		appG.ResponseFailError(merr)
		return
	}

	if uid == 0 { // 注册
		if id, err := userService.PhoneRegister(); err != nil {
			uid = id
			appG.ResponseFailError(err)
			return
		}
	}

	// 登录 make token
	token, err := util.GenerateToken(mAuth.Phone)
	if err != nil {
		logging.Error(err)
		appG.ResponseFailError(util.ErrNewCode(e.ERROR_AUTH_TOKEN))
		return
	}
	retMap := make(map[string]interface{})
	retMap["token"] = token
	retMap["uid"] = uid

	subsService := subs_service.Subs{
		UID: uid,
	}

	subs, errs := subsService.GetAll()
	if errs != nil {
		logging.Error(errs)
	}
	retMap["subs"] = subs
	//
	//resultsService := subs_service.Result{
	//	UID: uid,
	//}
	//results, errs := resultsService.GetAll()
	//if errs != nil {
	//	logging.Error(errs)
	//}
	//retMap["results"] = results

	appG.ResponseSuc(retMap)
}

// SendCode 发送手机验证码
// @Summary 发送手机验证码
// @accept application/x-www-form-urlencoded
// @Tags auth
// @Produce  json
// @Param phone formData string true "手机号"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/code [post]
func SendCode(c *gin.Context) {
	appG := app.GetGin(c)
	phone := c.PostForm("phone")

	if !util.RegPhone(phone) {
		appG.ResponseFailErrCode(e.ErrorPhoneNotValid)
		return
	}
	//  发送验证码
	code, err := userlogin.SendCode(phone)
	if err != nil {
		appG.ResponseFailError(err)
	}
	logging.Info("返回的验证码为", code, " 电话号码为：", phone)
	// 便于测试 此处将 code返回
	appG.ResponseSuc("验证码发送成功")
}
