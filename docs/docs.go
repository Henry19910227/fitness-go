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
        "/migrate/down": {
            "put": {
                "security": [
                    {
                        "icebaby_admin_token": []
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
                        "icebaby_admin_token": []
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
                        "icebaby_admin_token": []
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
                        "icebaby_admin_token": []
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
                        "icebaby_admin_token": []
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
                        "icebaby_admin_token": []
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
                        "icebaby_admin_token": []
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
        }
    },
    "definitions": {
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
                    "type": "object"
                },
                "msg": {
                    "description": "錯誤訊息",
                    "type": "string",
                    "example": "system error!"
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
                    "type": "object"
                },
                "msg": {
                    "description": "成功訊息",
                    "type": "string",
                    "example": "success!"
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
