package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

// 服务
type App struct {
	JwtSecret string
	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string

	// send sms
	SecretID string
	SecretKey string
}

var AppSetting = &App{}

// 数据库
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	LogMode     bool
}

var DatabaseSetting = &Database{}

// 服务器配置
type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

// redis
type Redis struct {
	Host     string
	Password string
	// 最大空闲连接数
	MaxIdle int
	// 在给定时间内，允许分配的最大连接数（当为零时，没有限制）
	MaxActive int
	// 在给定时间内将会保持空闲状态，若到达时间限制则关闭连接（当为零时，没有限制）
	IdleTimeout time.Duration
}

// RedisSetting Redis 缓存
var RedisSetting = &Redis{}

var cfg *ini.File


func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis",RedisSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// mapTo函数将配置映射到 section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
