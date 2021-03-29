package e

const (
	SUCCESS                      = 200
	ERROR                        = 500
	INVALID_PARAMS               = 400
	INVALID_PARAMS_WITHOUT_TOKEN = 401

	// 认证错误
	ERROR_AUTH_CHECK_TOKEN_FAIL = 20001 + iota
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT
	ERROR_AUTH_TOKEN
	ERROR_AUTH

	//todo 获取订阅错误从30000开始
	ErrorSubsGetInfo = iota + 30001
	ErrorSubsAdd
	ErrorSubsEdit
	ErrorSubsFcdyNotExisted
	ErrorDeleteSubsFailed
	ErrorSubsDtdyNotExisted
	ErrorSubsZtdyNotExisted
	ErrorResultNotExisted
	ErrorResultGetInfo
	ErrorResultMarkEdit

	// --- 客户端错误
	ErrorUserGetInfo = iota + 40001
	ErrorUserGetLogin
	ErrorUserRegName
	ErrorUserNameExist
	ErrorPhoneNotValid
	ErrorPhoneCodeSend
	ErrorPhoneCodeExpired
	ErrorPhoneCodeNotValid
	ErrorUserName
	ErrorUserNameNotExist
	ErrorUserInfoEmpty
	ErrorUserLoginEmpty
	ErrorUserPwd

	// --- 服务器错误
	ErrorExecSql = iota + 50001
	ErrorPasswordEncrypt
	ErrorUserInfoCreate
	ErrorUserLoginCreate
	ErrorOauthState
	ErrorOauthCode
	ErrorOauthInfo
)

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	INVALID_PARAMS_WITHOUT_TOKEN:   "不存在Token参数",

	// 订阅错误
	ErrorSubsGetInfo:        "获取订阅信息失败",
	ErrorSubsAdd:            "增加订阅失败",
	ErrorSubsEdit:           "修改订阅失败",
	ErrorSubsFcdyNotExisted: "分词订阅不存在",
	ErrorDeleteSubsFailed:   "删除订阅失败",
	ErrorSubsDtdyNotExisted: "动态订阅不存在",
	ErrorSubsZtdyNotExisted: "主题订阅不存在",
	ErrorResultNotExisted:   "不存在订阅结果",
	ErrorResultGetInfo:      "不存在订阅结果",
	ErrorResultMarkEdit:     "修改阅读标签失败",

	// 用户错误
	ErrorUserGetInfo:       "获取到用户失败.",
	ErrorUserGetLogin:      "获取到帐户失败.",
	ErrorUserRegName:       "用户名输入格式错误.",
	ErrorUserNameExist:     "用户名已存在.",
	ErrorPhoneNotValid:     "手机号验证失败.",
	ErrorPhoneCodeSend:     "验证码发送失败.",
	ErrorPhoneCodeExpired:  "验证码已过期.",
	ErrorPhoneCodeNotValid: "验证码验证失败.",
	ErrorUserNameNotExist:  "用户名不存在.",
	ErrorUserInfoEmpty:     "用户不存在.",
	ErrorUserLoginEmpty:    "帐户不存在.",
	ErrorUserPwd:           "密码错误.",
	ErrorUserName:          "获取用户名失败.",

	// --- 服务器错误
	ErrorPasswordEncrypt: "密码加密失败.",
	ErrorUserInfoCreate:  "用户信息创建失败.",
	ErrorUserLoginCreate: "用户帐户创建失败.",
	ErrorOauthState:      "三方登录状态码错误.",
	ErrorOauthCode:       "三方登录获取token失败.",
	ErrorOauthInfo:       "三方登录获取信息失败.",
	// --- end
}

// 由错误代码获取错误消息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
