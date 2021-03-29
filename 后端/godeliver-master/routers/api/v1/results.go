package v1

import (
	"github.com/gin-gonic/gin"
	"godeliver/pkg/app"
	"godeliver/pkg/e"
	"godeliver/pkg/logging"
	"godeliver/pkg/util"
	"godeliver/service/results_service"
	"strings"
)

type result struct {
	// 结果id
	Id uint `json:"id" example:"10"`
	// 订阅类型 三种：fcdy dtdy ztdy
	SubType string `json:"sub_type" example:"dtdy"`
}

// 获取结果
// @Summary 获取结果
// @Tags result
// @Produce  json
// @Param subscription_id path string true "放在链接中get请求：如 http://127.0.0.1:8000/api/v1/results/25_ztdy"
// @Router /api/v1/results/{subscription_id} [get]
func GetResult(c *gin.Context) {
	appG := app.GetGin(c)
	sid := c.Param("sid")
	if sid == "" {
		appG.ResponseFailErrCode(e.INVALID_PARAMS)
		return
	}
	logging.Info("获取订阅结果 sid is " + sid)
	// 拆分sid
	service := strings.Split(sid, "_")[1]
	logging.Info("请求了查询结果接口，订阅类型是：" + service)
	// 调用服务
	switch service {
	case "fcdy":
		resultService := results_service.FcdyResult{SubscriptionId: sid}
		results, err := resultService.Get()
		if err != nil {
			appG.ResponseFailErrCode(e.ErrorResultGetInfo)
		}
		appG.ResponseSuc(results)
	case "dtdy":
		resultService := results_service.DtdyResult{SubscriptionId: sid}
		results, err := resultService.Get()
		if err != nil {
			appG.ResponseFailErrCode(e.ErrorResultGetInfo)
		}
		appG.ResponseSuc(results)
	case "ztdy":
		resultService := results_service.ZtdyResult{SubscriptionId: sid}
		results, err := resultService.Get()
		if err != nil {
			appG.ResponseFailErrCode(e.ErrorResultGetInfo)
		}
		appG.ResponseSuc(results)
	}
}

// todo
// 标记已读的内容结果
// @Summary 标记已读的内容
// @Tags result
// @Produce  json
// @Param result body v1.result true "标记已读的内容"
// @Router /api/v1/results [post]
func EditResults(c *gin.Context) {
	appG := app.GetGin(c)
	var sResult result
	// 解析json 到结构体中
	if err := c.ShouldBindJSON(&sResult); err != nil {
		appG.ResponseFailError(util.ErrNewErr(err))
		return
	}

	resultsService := results_service.Result{
		Id:      sResult.Id,
		SubType: sResult.SubType,
	}

	err := resultsService.Edit()

	if err != nil {
		appG.ResponseFailErrCode(e.ErrorResultMarkEdit)
		return
	}
	appG.ResponseSuc("修改成功")

}

type results struct {
	// 订阅类型
	SubType string `json:"sub_type" example:"dtdy"`
	// 多个id一次性标记
	Ids []uint `json:"ids"`
	// 标记
	Mark string `json:"mark"`
}

// 标记已读的内容结果
// @Summary 标记已读的内容
// @Tags result
// @Produce  json
// @Param result body v1.result true "标记已读\未读的多个内容"
// @Router /api/v1/results [post]
func EditManyResults(c *gin.Context) {
	appG := app.GetGin(c)
	var sResult results
	// 解析json 到结构体中
	if err := c.ShouldBindJSON(&sResult); err != nil {
		appG.ResponseFailError(util.ErrNewErr(err))
		return
	}
	subtype := sResult.SubType
	mark := sResult.Mark
	for _, id := range sResult.Ids {
		resultsService := results_service.Result{
			Id:      id,
			SubType: subtype,
			Mark:    mark,
		}
		err := resultsService.Edit()
		if err != nil {
			logging.Error(err)
			return
		}
	}

	appG.ResponseSuc("修改成功")
}
