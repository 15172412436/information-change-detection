package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"syscall"

	"godeliver/models"
	"godeliver/pkg/gredis"
	"godeliver/pkg/logging"
	"godeliver/pkg/setting"
	"godeliver/routers"
	"log"
)

func main() {
	setting.Setup()
	logging.Setup()
	models.SetUp()

	err := gredis.Setup()
	if err != nil {
		logging.Error("redis setup failed", err)
	}

	router := routers.InitRouter()

	// 若 使用endless 保持过Godeliver 长时间运行
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, router)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

	// 正常模式
	//server := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ServerSetting.ReadTimeout,
	//	WriteTimeout:   setting.ServerSetting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//err = server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
