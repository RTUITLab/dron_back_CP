// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/auth/login": {
            "post": {
                "description": "login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "userInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.LoginResp"
                        }
                    },
                    "400": {
                        "description": "some user error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "401": {
                        "description": "not auth",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "500": {
                        "description": "internal",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            }
        },
        "/v1/auth/refresh": {
            "post": {
                "description": "refresh",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "refresh",
                "parameters": [
                    {
                        "description": "token",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RefreshReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.LoginResp"
                        }
                    },
                    "400": {
                        "description": "some user error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "401": {
                        "description": "not auth",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "500": {
                        "description": "internal",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            }
        },
        "/v1/module": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create module\nthat can do only admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "module"
                ],
                "summary": "create module",
                "parameters": [
                    {
                        "description": "Module info",
                        "name": "module",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/module.CreateModuleReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/module.CreateModuleResp"
                        }
                    },
                    "400": {
                        "description": "some user error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "401": {
                        "description": "not auth",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "500": {
                        "description": "internal",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            }
        },
        "/v1/module/submodule/{id}/test": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "add submodule test\nthat can do only admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "module"
                ],
                "summary": "add submodule test",
                "parameters": [
                    {
                        "description": "SubModuleTest info",
                        "name": "submoduletest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/module.AddSubModuleTestReq"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "id of submodule",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/module.AddSubModuleTestResp"
                        }
                    },
                    "400": {
                        "description": "some user error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "401": {
                        "description": "not auth",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "404": {
                        "description": "module not found",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "500": {
                        "description": "internal",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            }
        },
        "/v1/module/{id}/submodule": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "add submodule\nthat can do only admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "module"
                ],
                "summary": "add submodule",
                "parameters": [
                    {
                        "description": "SubModule info",
                        "name": "submodule",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/module.AddSubModuleReq"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "id of module",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/module.AddSubModuleResp"
                        }
                    },
                    "400": {
                        "description": "some user error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "401": {
                        "description": "not auth",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "404": {
                        "description": "module not found",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "500": {
                        "description": "internal",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            }
        },
        "/v1/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "get users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "match user on login field",
                        "name": "login",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "match user on role",
                        "name": "role",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "users",
                        "schema": {
                            "$ref": "#/definitions/user.GetUsersResponce"
                        }
                    },
                    "400": {
                        "description": "some user error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "401": {
                        "description": "not auth",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "404": {
                        "description": "user not found",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "500": {
                        "description": "internal",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create user with role\nthat can do only admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "create user",
                "parameters": [
                    {
                        "description": "User info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateUserReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/user.CreateUserResp"
                        }
                    },
                    "400": {
                        "description": "some user error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "401": {
                        "description": "not auth",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "500": {
                        "description": "internal",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            }
        },
        "/v1/user/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update user\nit can do this user or admin",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "update user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "some user error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "401": {
                        "description": "not auth",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "404": {
                        "description": "user not found",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "500": {
                        "description": "internal",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete user\nthat can do only admin",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "delete user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "some user error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "401": {
                        "description": "not auth",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "404": {
                        "description": "user not found",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "500": {
                        "description": "internal",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.LoginReq": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "auth.LoginResp": {
            "type": "object",
            "properties": {
                "accesstoken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "auth.RefreshReq": {
            "type": "object",
            "properties": {
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "err.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "module.AddAnswerResp": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string"
                },
                "correct": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "module.AddQuestionResp": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/module.AddAnswerResp"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "question": {
                    "type": "string"
                }
            }
        },
        "module.AddSubModuleReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "module.AddSubModuleResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "moduleID": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "module.AddSubModuleTestReq": {
            "type": "object",
            "properties": {
                "theoreticalTest": {
                    "$ref": "#/definitions/module.CreateTheoreticalTestReq"
                }
            }
        },
        "module.AddSubModuleTestResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "theoreticalTest": {
                    "$ref": "#/definitions/module.AddTheoreticalTestResp"
                }
            }
        },
        "module.AddTheoreticalTestResp": {
            "type": "object",
            "properties": {
                "questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/module.AddQuestionResp"
                    }
                }
            }
        },
        "module.CreateAnswerReq": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string"
                },
                "correct": {
                    "type": "boolean"
                }
            }
        },
        "module.CreateModuleReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "module.CreateModuleResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "module.CreateQuestionReq": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/module.CreateAnswerReq"
                    }
                },
                "question": {
                    "type": "string"
                }
            }
        },
        "module.CreateTheoreticalTestReq": {
            "type": "object",
            "properties": {
                "questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/module.CreateQuestionReq"
                    }
                }
            }
        },
        "user.CreateUserReq": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "student",
                        "teacher",
                        "admin"
                    ]
                }
            }
        },
        "user.CreateUserResp": {
            "type": "object",
            "properties": {
                "int": {
                    "type": "integer"
                },
                "login": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "user.GetUserResponce": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "login": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "user.GetUsersResponce": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.GetUserResponce"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/dron",
	Schemes:     []string{},
	Title:       "ROSSETI-DRON API",
	Description: "This is a server to get projects from github",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
