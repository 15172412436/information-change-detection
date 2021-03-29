package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"godeliver/pkg/e"
	"godeliver/pkg/logging"
	"net/http"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	var err error
	if c.ContentType() == "application/json" {
		err = c.BindJSON(form)
	} else {
		err = c.Bind(form)
	}

	if err != nil {
		logging.Error("参数绑定错误", err)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkError(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
