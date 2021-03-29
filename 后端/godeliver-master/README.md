# GoDeliver 

爬，推送


### 待完成的事项

~~- 对订阅的增删改查找~~
- 在encrypt模块中对密码的加密进行优化
- 实现爬取 
- 读入一种搜索引擎结果

### 文档结构
```
├── README.md 
├── conf 配置文件
│   └── app_default.ini 配置文件示例
├── docs api 文档
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod 
├── go.sum 
├── main.go 
├── middleware 中间件
│   └── jwt token 验证
│       └── jwt.go
├── models 
│   ├── model.go
│   ├── results
│   │   └── result.go
│   ├── subs
│   │   ├── dtSubs.go
│   │   ├── fcSubs.go
│   │   └── ztSubs.go
│   └── users
│       ├── login.go
│       └── user.go
├── pkg 
│   ├── app
│   │   ├── form.go
│   │   ├── request.go
│   │   └── respones.go
│   ├── e
│   │   └── msg.go
│   ├── file
│   │   └── file.go
│   ├── gredis
│   │   └── redis.go
│   ├── logging
│   │   ├── file.go
│   │   └── log.go
│   ├── setting
│   │   └── setting.go
│   ├── sms
│   │   └── sendSMS.go
│   └── util
│       ├── encrypt.go
│       ├── error.go
│       ├── jwt.go
│       ├── random.go
│       ├── reg.go
│       └── valid.go
├── routers 路由
│   ├── api
│   │   ├── user.go
│   │   └── v1
│   │       └── subs.go
│   └── router.go
├── runtime 运行时
│   └── logs
└── service 服务
    ├── caches
    │   ├── cache.go
    │   └── phone.go
    ├── subs_service
    │   ├── result.go
    │   └── subs.go
    └── userlogin
        └── user.go
```