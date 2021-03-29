package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"godeliver/pkg/e"
	"godeliver/pkg/logging"
	"godeliver/pkg/util"
	"net/http"
	"strconv"
	"strings"
)

type Gin struct {
	C *gin.Context
}

type HTTPSuccess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"msg" example:"ok" `
	Data    string `json:"data" example:"null"`
}

// GetGin 获取Gin
func GetGin(c *gin.Context) Gin {
	return Gin{c}
}

// Response 返回的数据
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})

	return
}

// ResponseSuc 返回成功
func (g *Gin) ResponseSuc(data interface{}) {
	g.C.JSON(http.StatusOK, gin.H{
		"code":   e.SUCCESS,
		"msg":    e.GetMsg(e.SUCCESS),
		"data":   data,
		"status": 1,
	})
	return
}

// ResponseFail 返回失败
func (g *Gin) ResponseFail() {
	g.C.JSON(http.StatusOK, gin.H{
		"code": e.ERROR,
		"msg":  e.GetMsg(e.ERROR),
		"data": nil,
		"status": 0,
	})
	return
}

// ResponseFailErrCode 返回失败
func (g *Gin) ResponseFailErrCode(errCode int) {
	errMsg := "code : " + strconv.Itoa(errCode) + "msg : " + e.GetMsg(errCode)
	MarkError(errMsg)

	g.C.JSON(http.StatusOK, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": nil,
		"status": 0,
	})
	return
}

// ResponseFailError 返回自定义的错误类型
func (g *Gin) ResponseFailError(error util.Error) {
	msg := error.Error()
	MarkError(msg)
	g.C.JSON(http.StatusOK, gin.H{
		"code": error.Code(),
		"msg":  msg,
		"data": nil,
		"status": 0,
	})
	return
}

// ResponseFailValidParam 验证参数错误
func (g *Gin) ResponseFailValidParam(err error) {
	errs := err.(validator.ValidationErrors)

	jsonKey := errs[0].Field()
	fieldName, _ := util.GetTrans().T(jsonKey)
	msg := strings.Replace(errs[0].Translate(util.GetTrans()), jsonKey, fieldName, -1)
	//jsonKey = jsonKey[2 : len(jsonKey)-2]
	//fmt.Println(jsonKey, ":", msg)

	MarkError(msg)
	g.C.JSON(http.StatusOK, gin.H{
		"code": e.INVALID_PARAMS,
		"msg":  msg,
		"data": nil,
		"status": 0,
	})
	return
}

func MarkError(v ...interface{}) {
	logging.Error(v...)
	return
}
