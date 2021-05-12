// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login/admin/email": {
            "post": {
                "description": "管理者使用信箱登入",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "管理者使用信箱登入",
                "parameters": [
                    {
                        "description": "輸入參數",
                        "name": "json_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validator.LoginByEmailBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登入成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.SuccessLoginResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/logindto.Admin"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "登入失敗",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/login/user/email": {
            "post": {
                "description": "用戶使用信箱登入",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "用戶使用信箱登入",
                "parameters": [
                    {
                        "description": "輸入參數",
                        "name": "json_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validator.LoginByEmailBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登入成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.SuccessLoginResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/logindto.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "登入失敗",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/logout/admin": {
            "post": {
                "security": [
                    {
                        "fitness_admin_token": []
                    }
                ],
                "description": "管理員登出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "管理員登出",
                "responses": {
                    "200": {
                        "description": "登出成功",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResult"
                        }
                    },
                    "400": {
                        "description": "登出失敗",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/logout/user": {
            "post": {
                "security": [
                    {
                        "fitness_user_token": []
                    }
                ],
                "description": "用戶登出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "用戶登出",
                "responses": {
                    "200": {
                        "description": "登出成功",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResult"
                        }
                    },
                    "400": {
                        "description": "登出失敗",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/migrate/down": {
            "put": {
                "security": [
                    {
                        "fitness_admin_token": []
                    }
                ],
                "description": "回滾 Schema 至初始版本",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Migrate"
                ],
                "summary": "回滾 Schema 至初始版本",
                "responses": {
                    "200": {
                        "description": "成功!",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResult"
                        }
                    },
                    "400": {
                        "description": "失敗!",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/migrate/down/{step}": {
            "put": {
                "security": [
                    {
                        "fitness_admin_token": []
                    }
                ],
                "description": "將 Schema 回滾N個版本",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Migrate"
                ],
                "summary": "將 Schema 回滾N個版本",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "版本跨度",
                        "name": "step",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功!",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResult"
                        }
                    },
                    "400": {
                        "description": "失敗!",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/migrate/force/{version}": {
            "put": {
                "security": [
                    {
                        "fitness_admin_token": []
                    }
                ],
                "description": "Schema 升級時遇到錯誤時的操作模式",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Migrate"
                ],
                "summary": "修正 Schema 版本並解除錯誤狀態",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "版本號",
                        "name": "version",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功!",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResult"
                        }
                    },
                    "400": {
                        "description": "失敗!",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/migrate/up": {
            "put": {
                "security": [
                    {
                        "fitness_admin_token": []
                    }
                ],
                "description": "將 Schema 升至最新版本",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Migrate"
                ],
                "summary": "將 Schema 升至最新版本",
                "responses": {
                    "200": {
                        "description": "成功!",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResult"
                        }
                    },
                    "400": {
                        "description": "失敗!",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/migrate/up/{step}": {
            "put": {
                "security": [
                    {
                        "fitness_admin_token": []
                    }
                ],
                "description": "將 Schema 升級N個版本",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Migrate"
                ],
                "summary": "將 Schema 升級N個版本",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "版本跨度",
                        "name": "step",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功!",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResult"
                        }
                    },
                    "400": {
                        "description": "失敗!",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/migrate/version": {
            "get": {
                "security": [
                    {
                        "fitness_admin_token": []
                    }
                ],
                "description": "獲取當前 Schema 版本",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Migrate"
                ],
                "summary": "獲取當前 Schema 版本",
                "responses": {
                    "200": {
                        "description": "成功!",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResult"
                        }
                    },
                    "400": {
                        "description": "失敗!",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/migrate/version/{version}": {
            "put": {
                "security": [
                    {
                        "fitness_admin_token": []
                    }
                ],
                "description": "升級至指定 Schema 版本",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Migrate"
                ],
                "summary": "升級至指定 Schema 版本",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "版本號",
                        "name": "version",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功!",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResult"
                        }
                    },
                    "400": {
                        "description": "失敗!",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/register/email": {
            "post": {
                "description": "使用信箱註冊",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Register"
                ],
                "summary": "使用信箱註冊",
                "parameters": [
                    {
                        "description": "輸入參數",
                        "name": "json_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validator.RegisterForEmailBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "註冊成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/registerdto.Register"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "註冊失敗",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/register/email/otp": {
            "post": {
                "description": "發送 Email OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Register"
                ],
                "summary": "發送 Email OTP",
                "parameters": [
                    {
                        "description": "輸入參數",
                        "name": "Parameter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validator.EmailBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "驗證郵件已寄出",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/registerdto.OTP"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "發送失敗",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/register/email/validate": {
            "post": {
                "description": "驗證信箱是否可使用",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Register"
                ],
                "summary": "驗證信箱是否可使用",
                "parameters": [
                    {
                        "description": "輸入參數",
                        "name": "json_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validator.ValidateEmailDupBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "此暱稱可使用",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResult"
                        }
                    },
                    "400": {
                        "description": "該資料已存在",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/register/nickname/validate": {
            "post": {
                "description": "驗證暱稱是否可使用",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Register"
                ],
                "summary": "驗證暱稱是否可使用",
                "parameters": [
                    {
                        "description": "輸入參數",
                        "name": "json_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validator.ValidateNicknameDupBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "此暱稱可使用",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResult"
                        }
                    },
                    "400": {
                        "description": "該資料已存在",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        },
        "/user/my/info": {
            "patch": {
                "security": [
                    {
                        "fitness_user_token": []
                    }
                ],
                "description": "更新個人資料",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "更新個人資料",
                "parameters": [
                    {
                        "description": "更新欄位",
                        "name": "json_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/validator.UpdateMyUserInfoBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功!",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.SuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/userdto.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "失敗!",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "logindto.Admin": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "信箱",
                    "type": "string",
                    "example": "henry@gmail.com"
                },
                "id": {
                    "description": "帳戶id",
                    "type": "integer",
                    "example": 1
                },
                "lv": {
                    "description": "身份 (1:一般管理員)",
                    "type": "integer",
                    "example": 1
                },
                "nickname": {
                    "description": "暱稱",
                    "type": "string",
                    "example": "henry"
                }
            }
        },
        "logindto.User": {
            "type": "object",
            "properties": {
                "account": {
                    "description": "帳號",
                    "type": "string",
                    "example": "henry@gmail.com"
                },
                "account_type": {
                    "description": "帳號類型 (1:Email註冊/2:FB註冊/3:Google註冊/4:Line註冊)",
                    "type": "integer",
                    "example": 1
                },
                "birthday": {
                    "description": "生日",
                    "type": "string",
                    "example": "1991-02-27"
                },
                "create_at": {
                    "description": "創建日期",
                    "type": "string",
                    "example": "2021-06-01 12:00:00"
                },
                "device_token": {
                    "description": "推播 Token",
                    "type": "string",
                    "example": "f144b48d9695..."
                },
                "email": {
                    "description": "信箱",
                    "type": "string",
                    "example": "henry@gmail.com"
                },
                "experience": {
                    "description": "經驗 (0:未指定/1:初學/2:中級/3:中高/4:專業)",
                    "type": "integer",
                    "example": 2
                },
                "height": {
                    "description": "身高",
                    "type": "number",
                    "example": 176.5
                },
                "id": {
                    "description": "帳戶id",
                    "type": "integer",
                    "example": 10001
                },
                "nickname": {
                    "description": "暱稱",
                    "type": "string",
                    "example": "Henry"
                },
                "sex": {
                    "description": "性別 (m:男/f:女)",
                    "type": "string",
                    "example": "m"
                },
                "target": {
                    "description": "目標 (0:未指定/1:減重/2:維持健康/3:增肌)",
                    "type": "integer",
                    "example": 3
                },
                "update_at": {
                    "description": "修改日期",
                    "type": "string",
                    "example": "2021-06-01 12:00:00"
                },
                "user_status": {
                    "description": "用戶狀態 (1:正常/2:違規/3:刪除)",
                    "type": "integer",
                    "example": 1
                },
                "user_type": {
                    "description": "用戶狀態 (1:一般用戶/2:訂閱用戶)",
                    "type": "integer",
                    "example": 1
                },
                "weight": {
                    "description": "體重",
                    "type": "number",
                    "example": 72.5
                }
            }
        },
        "model.Data": {
            "type": "object"
        },
        "model.ErrorResult": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "錯誤碼",
                    "type": "integer",
                    "example": 9000
                },
                "data": {
                    "description": "回傳資料",
                    "$ref": "#/definitions/model.Data"
                },
                "msg": {
                    "description": "錯誤訊息",
                    "type": "string",
                    "example": "system error!"
                }
            }
        },
        "model.SuccessLoginResult": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "狀態碼",
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "description": "回傳資料",
                    "$ref": "#/definitions/model.Data"
                },
                "msg": {
                    "description": "成功訊息",
                    "type": "string",
                    "example": "success!"
                },
                "token": {
                    "description": "Token",
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTQ0MDc3NjMsInN1YiI6IjEwMDEzIn0.Z5UeEC8ArCVYej9kI1paXD2f5FMFiTfeLpU6e_CZZw0"
                }
            }
        },
        "model.SuccessResult": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "狀態碼",
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "description": "回傳資料",
                    "$ref": "#/definitions/model.Data"
                },
                "msg": {
                    "description": "成功訊息",
                    "type": "string",
                    "example": "success!"
                }
            }
        },
        "registerdto.OTP": {
            "type": "object",
            "properties": {
                "otp_code": {
                    "description": "信箱驗證碼",
                    "type": "string",
                    "example": "254235"
                }
            }
        },
        "registerdto.Register": {
            "type": "object",
            "properties": {
                "user_id": {
                    "description": "用戶ID",
                    "type": "integer",
                    "example": 10001
                }
            }
        },
        "userdto.User": {
            "type": "object",
            "properties": {
                "account": {
                    "description": "帳號",
                    "type": "string",
                    "example": "henry@gmail.com"
                },
                "account_type": {
                    "description": "帳號類型 (1:Email註冊/2:FB註冊/3:Google註冊/4:Line註冊)",
                    "type": "integer",
                    "example": 1
                },
                "birthday": {
                    "description": "生日",
                    "type": "string",
                    "example": "1991-02-27"
                },
                "create_at": {
                    "description": "創建日期",
                    "type": "string",
                    "example": "2021-06-01 12:00:00"
                },
                "device_token": {
                    "description": "推播 Token",
                    "type": "string",
                    "example": "f144b48d9695..."
                },
                "email": {
                    "description": "信箱",
                    "type": "string",
                    "example": "henry@gmail.com"
                },
                "experience": {
                    "description": "經驗 (0:未指定/1:初學/2:中級/3:中高/4:專業)",
                    "type": "integer",
                    "example": 2
                },
                "height": {
                    "description": "身高",
                    "type": "number",
                    "example": 176.5
                },
                "id": {
                    "description": "帳戶id",
                    "type": "integer",
                    "example": 10001
                },
                "nickname": {
                    "description": "暱稱",
                    "type": "string",
                    "example": "Henry"
                },
                "sex": {
                    "description": "性別 (m:男/f:女)",
                    "type": "string",
                    "example": "m"
                },
                "target": {
                    "description": "目標 (0:未指定/1:減重/2:維持健康/3:增肌)",
                    "type": "integer",
                    "example": 3
                },
                "update_at": {
                    "description": "修改日期",
                    "type": "string",
                    "example": "2021-06-01 12:00:00"
                },
                "user_status": {
                    "description": "用戶狀態 (1:正常/2:違規/3:刪除)",
                    "type": "integer",
                    "example": 1
                },
                "user_type": {
                    "description": "用戶狀態 (1:一般用戶/2:訂閱用戶)",
                    "type": "integer",
                    "example": 1
                },
                "weight": {
                    "description": "體重",
                    "type": "number",
                    "example": 72.5
                }
            }
        },
        "validator.EmailBody": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@gmail.com"
                }
            }
        },
        "validator.LoginByEmailBody": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "信箱",
                    "type": "string",
                    "example": "test@gmail.com"
                },
                "password": {
                    "description": "密碼 (8~16字元)",
                    "type": "string",
                    "example": "12345678"
                }
            }
        },
        "validator.RegisterForEmailBody": {
            "type": "object",
            "required": [
                "email",
                "nickname",
                "otp_code",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "信箱",
                    "type": "string",
                    "example": "test@gmail.com"
                },
                "nickname": {
                    "description": "暱稱 (1~16字元)",
                    "type": "string",
                    "example": "henry"
                },
                "otp_code": {
                    "description": "信箱驗證碼",
                    "type": "string",
                    "example": "531476"
                },
                "password": {
                    "description": "密碼 (8~16字元)",
                    "type": "string",
                    "example": "12345678"
                }
            }
        },
        "validator.UpdateMyUserInfoBody": {
            "type": "object",
            "properties": {
                "birthday": {
                    "description": "生日",
                    "type": "string",
                    "example": "1991-02-27"
                },
                "email": {
                    "description": "信箱",
                    "type": "string",
                    "example": "henry@gmail.com"
                },
                "experience": {
                    "description": "經驗 (0:未指定/1:初學/2:中級/3:中高/4:專業)",
                    "type": "string",
                    "example": "2"
                },
                "height": {
                    "description": "身高 (最大230)",
                    "type": "string",
                    "example": "176.5"
                },
                "nickname": {
                    "description": "暱稱 (1~16字元)",
                    "type": "string",
                    "example": "henry"
                },
                "sex": {
                    "description": "性別 (f:女/m:男)",
                    "type": "string",
                    "example": "m"
                },
                "target": {
                    "description": "目標 (0:未指定/1:減重/2:維持健康/3:增肌)",
                    "type": "string",
                    "example": "3"
                },
                "weight": {
                    "description": "體重 (最大230)",
                    "type": "string",
                    "example": "70.5"
                }
            }
        },
        "validator.ValidateEmailDupBody": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "description": "信箱",
                    "type": "string",
                    "example": "henry@gmail.com"
                }
            }
        },
        "validator.ValidateNicknameDupBody": {
            "type": "object",
            "required": [
                "nickname"
            ],
            "properties": {
                "nickname": {
                    "description": "暱稱 (1~16字元)",
                    "type": "string",
                    "example": "henry"
                }
            }
        }
    },
    "securityDefinitions": {
        "fitness_admin_token": {
            "type": "apiKey",
            "name": "Token",
            "in": "header"
        },
        "fitness_trainer_token": {
            "type": "apiKey",
            "name": "Token",
            "in": "header"
        },
        "fitness_user_token": {
            "type": "apiKey",
            "name": "Token",
            "in": "header"
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
	Title:       "fitness api",
	Description: "健身平台 api",
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
