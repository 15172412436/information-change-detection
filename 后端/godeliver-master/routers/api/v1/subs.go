package v1

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"godeliver/pkg/app"
	"godeliver/pkg/e"
	"godeliver/service/subs_service"
	"net/http"
	"strconv"
	"strings"
)

// 获取订阅
// @Summary 获取订阅
// @Tags subscription
// @Produce  json
// @Param uid path int true "ID"
// @Router /api/v1/subs_service/{uid} [get]
func GetSubs(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	uid := c.Param("uid")
	fmt.Println(uid)
	var id int
	if uid != "" {
		id, _ = strconv.Atoi(uid)

		valid.Min(id, 1, "uid")
	}
	if valid.HasErrors() {
		app.MarkError(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	var idu uint
	idu = uint(id)
	fmt.Println(idu)
	subsService := subs_service.Subs{
		UID: idu,
	}
	// todo 获取数量 fcdy多少 dtdy多少个
	// count()

	subs, err := subsService.GetAll()
	if err != nil {
		appG.ResponseFailErrCode(e.ErrorSubsGetInfo)
	}

	appG.ResponseSuc(subs)

}

type AddSubForm struct {
	// 订阅类型 两种 dtdy或fcdy
	SubType string `json:"sub_type" form:"sub_type" valid:"Required" example:"dtdy"`
	// 用户id
	UID string `json:"uid" form:"uid" valid:"Required" example:"13"`
	// 订阅链接-主题或者链接
	Content string `json:"content" form:"content" valid:"Required" example:"url or Subject"`
	// set_time
	SetTime int `json:"set_time" form:"set_time" valid:"Required" example:"12"`
	// 邮箱
	Email string `json:"email" form:"email" valid:"Required" example:"zhuge666@qq.com"`
	// 频率
	Frequency string `json:"frequency" form:"frequency" valid:"Required" example:"everyday"`
	// 关键词 仅在type 为 ztdy 存在
	Keyword string `json:"keyword" form:"keyword"`
	// 返回类型 PDF Email
	ReturnType string `json:"return_type" form:"return_type"`
	//
	Subject string `json:"subject" form:"subject"`
	// 范围
	Fanwei string `json:"fanwei" form:"fanwei"`
}

// 增加订阅
// @Summary 增加订阅
// @Tags subscription
// @Produce  json
// @Param subs body v1.AddSubForm true "订阅信息"
// @Router /api/v1/subs_service [post]
func AddSub(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddSubForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	// uid string 转换成 uint
	temp, _ := strconv.Atoi(form.UID)
	uid := uint(temp)
	subsService := subs_service.Subs{
		SubType:    form.SubType,
		UID:        uid,
		Content:    form.Content,
		SetTime:    form.SetTime,
		Email:      form.Email,
		Frequency:  form.Frequency,
		ReturnType: form.ReturnType,
		Subject:    form.Subject,
		Fanwei:     form.Fanwei,
	}
	fmt.Println(form.Keyword)
	if form.Keyword != "" {
		subsService.Keyword = form.Keyword
	}
	subId, err := subsService.Add()
	if err != nil {
		appG.ResponseFailErrCode(e.ErrorSubsAdd)
		return
	}
	// todo 返回订阅ID
	retMap := make(map[string]interface{})
	retMap["subscription_id"] = subId
	id, err := strconv.Atoi(strings.Split(subId, "_")[0])
	retMap["id"] = id
	appG.ResponseSuc(retMap)
}

// 删除订阅
// @Summary 删除订阅
// @Tags subscription
// @Produce  json
// @Param subscription_id path int true "ID"
// @Router /api/v1/subs_service/{subscription_id} [delete]
func DeleteSub(c *gin.Context) {
	appG := app.Gin{C: c}
	//valid := validation.Validation{}
	sid := c.Param("sid")
	//valid.MinSize(sid, 4, "subscription_id")
	//if valid.HasErrors() {
	//	app.MarkError(valid.Errors)
	//	appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	//	return
	//}
	if sid == "" {
		appG.ResponseFailErrCode(e.INVALID_PARAMS)
		return
	}
	fmt.Println("sid", sid)
	subsService := subs_service.Subs{SubscriptionId: sid}
	err := subsService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorDeleteSubsFailed, nil)
		return
	}

	appG.ResponseSuc("删除成功")
}

type EditSubForm struct {
	// 订阅id 用于修改
	SubscriptionID string `json:"subscription_id" form:"subscription_id" valid:"Required" example:"12_dtdy"`
	// 订阅类型 两种 dtdy或fcdy
	SubType string `json:"sub_type" form:"sub_type" example:"dtdy"`
	// 用户id
	UID string `json:"uid" form:"uid" example:"13"`
	// 订阅链接-主题或者链接
	Content string `json:"content" form:"content" example:"url or subject"`
	// set_time
	SetTime int `json:"set_time" form:"set_time" example:"12"`
	// 邮箱
	Email string `json:"email" form:"email" example:"zhuge666@qq.com"`
	// 频率
	Frequency string `json:"frequency" form:"frequency" example:"everyday"`
	// 关键词 仅存在ztdy
	Keyword string `json:"keyword" form:"keyword"`
	// 返回类型 PDF，Email
	ReturnType string `json:"return_type" form:"return_type"`
	Subject    string `json:"subject" form:"subject"`
	Fanwei     string `json:"fanwei" form:"fanwei"`
}

// 修改订阅
// @Summary 修改订阅
// @Tags subscription
// @Produce  json
// @Param subs body v1.EditSubForm true "订阅信息"
// @Router /api/v1/subs_service/ [put]
func EditSub(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = EditSubForm{SubscriptionID: c.Param("subscription_id")}
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	// uid string 转换成 uint
	temp, _ := strconv.Atoi(form.UID)
	uid := uint(temp)

	subsService := subs_service.Subs{
		SubscriptionId: form.SubscriptionID,
		SubType:        form.SubType,
		UID:            uid,
		Content:        form.Content,
		SetTime:        form.SetTime,
		Email:          form.Email,
		Frequency:      form.Frequency,
		ReturnType:     form.ReturnType,
		Subject:        form.Subject,
		Fanwei:         form.Fanwei,
	}
	if form.Keyword != "" {
		subsService.Keyword = form.Keyword
	}
	fmt.Println(subsService)
	err := subsService.Edit()
	if err != nil {
		appG.ResponseFailErrCode(e.ErrorSubsAdd)
		return
	}

	appG.ResponseSuc("修改成功")
}
