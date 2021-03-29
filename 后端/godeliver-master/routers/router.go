package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "godeliver/docs"
	"godeliver/middleware/jwt"
	"godeliver/pkg/setting"
	"godeliver/routers/api"
	v1 "godeliver/routers/api/v1"
)

func InitRouter() *gin.Engine {
	// gin 启动模式
	gin.SetMode(setting.ServerSetting.RunMode)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	root := r.Group("")
	{
		// 注册登陆
		auth := root.Group("/auth")
		{
			auth.POST("/register", api.Register)
			auth.POST("/login", api.Login)
			auth.POST("/code", api.SendCode)
			auth.POST("/phoneLogin", api.PhoneLogin)
		}

		// api v1
		apiV1 := root.Group("/api/v1")
		apiV1.Use(jwt.JWT())
		{
			// 获得订阅列表
			apiV1.GET("/subs_service/:uid", v1.GetSubs)
			// 获取订阅单个的订阅内容
			//apiV1.GET("/subs_service/:sid",v1.GetSub)
			// 添加订阅
			apiV1.POST("/subs_service", v1.AddSub)
			// 删除订阅
			apiV1.DELETE("/subs_service/:sid", v1.DeleteSub)
			// 修改订阅
			apiV1.PUT("/subs_service/", v1.EditSub)

			// 获取订阅结果集
			apiV1.GET("/results/:sid", v1.GetResult)
			// 标记已读
			apiV1.POST("/results", v1.EditResults)
			// 标记多个已读
			apiV1.POST("/someresult", v1.EditManyResults)
		}
	}
	// todo swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
