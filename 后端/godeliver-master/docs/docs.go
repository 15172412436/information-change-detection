// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-07-21 18:28:16.588983 +0800 CST m=+0.074378807

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/results": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "result"
                ],
                "summary": "标记已读的内容",
                "parameters": [
                    {
                        "description": "标记已读的内容",
                        "name": "result",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.result"
                        }
                    }
                ]
            }
        },
        "/api/v1/results/{subscription_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "result"
                ],
                "summary": "获取结果",
                "parameters": [
                    {
                        "type": "string",
                        "description": "放在链接中get请求：如 http://127.0.0.1:8000/api/v1/results/25_ztdy",
                        "name": "subscription_id",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/api/v1/subs_service": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscription"
                ],
                "summary": "增加订阅",
                "parameters": [
                    {
                        "description": "订阅信息",
                        "name": "subs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.AddSubForm"
                        }
                    }
                ]
            }
        },
        "/api/v1/subs_service/": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscription"
                ],
                "summary": "修改订阅",
                "parameters": [
                    {
                        "description": "订阅信息",
                        "name": "subs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.EditSubForm"
                        }
                    }
                ]
            }
        },
        "/api/v1/subs_service/{subscription_id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscription"
                ],
                "summary": "删除订阅",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "subscription_id",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/api/v1/subs_service/{uid}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscription"
                ],
                "summary": "获取订阅",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/auth/code": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "发送手机验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "账号密码登录",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "auth",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.auth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "$ref": "#/definitions/api.token"
                        }
                    }
                }
            }
        },
        "/auth/phoneLogin": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "手机号快速登陆/注册",
                "parameters": [
                    {
                        "description": "手机号快速登录/注册",
                        "name": "auth",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.phone"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{\"token\":token},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "账号密码注册",
                "parameters": [
                    {
                        "description": "账号密码登录/注册",
                        "name": "auth",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.auth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{\"token\":token},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.auth": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "PassWord 密码",
                    "type": "string",
                    "example": "zhangsan"
                },
                "username": {
                    "description": "UserName 用户名",
                    "type": "string",
                    "example": "zhangsan"
                }
            }
        },
        "api.phone": {
            "type": "object",
            "required": [
                "code",
                "phone"
            ],
            "properties": {
                "code": {
                    "description": "Code 手机号验证码",
                    "type": "string",
                    "example": "123456"
                },
                "phone": {
                    "description": "Phone 手机号",
                    "type": "string",
                    "example": "13938738804"
                }
            }
        },
        "api.token": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "v1.AddSubForm": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "订阅链接-主题或者链接",
                    "type": "string",
                    "example": "url or Subject"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "zhuge666@qq.com"
                },
                "frequency": {
                    "description": "频率",
                    "type": "string",
                    "example": "everyday"
                },
                "keyword": {
                    "description": "关键词 仅在type 为 ztdy 存在",
                    "type": "string"
                },
                "set_time": {
                    "description": "set_time",
                    "type": "integer",
                    "example": 12
                },
                "sub_type": {
                    "description": "订阅类型 两种 dtdy或fcdy",
                    "type": "string",
                    "example": "dtdy"
                },
                "uid": {
                    "description": "用户id",
                    "type": "integer",
                    "example": 13
                }
            }
        },
        "v1.EditSubForm": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "订阅链接-主题或者链接",
                    "type": "string",
                    "example": "url or subject"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "zhuge666@qq.com"
                },
                "frequency": {
                    "description": "频率",
                    "type": "string",
                    "example": "everyday"
                },
                "keyword": {
                    "description": "关键词 仅存在ztdy",
                    "type": "string"
                },
                "set_time": {
                    "description": "set_time",
                    "type": "integer",
                    "example": 12
                },
                "sub_type": {
                    "description": "订阅类型 两种 dtdy或fcdy",
                    "type": "string",
                    "example": "dtdy"
                },
                "subscription_id": {
                    "description": "订阅id 用于修改",
                    "type": "string",
                    "example": "12_dtdy"
                },
                "uid": {
                    "description": "用户id",
                    "type": "integer",
                    "example": 13
                }
            }
        },
        "v1.result": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "结果id",
                    "type": "integer",
                    "example": 10
                },
                "sub_type": {
                    "description": "订阅类型 三种：fcdy dtdy ztdy",
                    "type": "string",
                    "example": "dtdy"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
