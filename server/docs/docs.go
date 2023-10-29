// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/question/": {
            "put": {
                "description": "更新题目",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "题目"
                ],
                "summary": "更新",
                "parameters": [
                    {
                        "description": "题目",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Question"
                        }
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "添加题目",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "题目"
                ],
                "summary": "添加",
                "parameters": [
                    {
                        "description": "题目",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Question"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/question/list/{id}": {
            "get": {
                "description": "获取题目列表列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "题目"
                ],
                "summary": "题目列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "关键字",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页大小",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "题目id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/question/{id}": {
            "get": {
                "description": "获取题目",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "题目"
                ],
                "summary": "获取",
                "responses": {}
            },
            "delete": {
                "description": "删除题目",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "题目"
                ],
                "summary": "删除",
                "responses": {}
            }
        },
        "/survey/": {
            "put": {
                "description": "更新问卷",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "问卷"
                ],
                "summary": "更新",
                "parameters": [
                    {
                        "description": "问卷",
                        "name": "survey",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Survey"
                        }
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "添加问卷",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "问卷"
                ],
                "summary": "添加",
                "parameters": [
                    {
                        "description": "问卷",
                        "name": "survey",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Survey"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/survey/list": {
            "get": {
                "description": "获取题目列表列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "问卷"
                ],
                "summary": "问卷列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "关键字",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页大小",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/survey/{id}": {
            "get": {
                "description": "获取获取",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "问卷"
                ],
                "summary": "获取",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "题目id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "删除问卷",
                "tags": [
                    "问卷"
                ],
                "summary": "删除",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "问卷id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "entity.Option": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "label": {
                    "type": "string"
                },
                "questionId": {
                    "type": "integer"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "entity.Question": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "options": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Option"
                    }
                },
                "order": {
                    "type": "integer"
                },
                "surveyID": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "entity.Survey": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "questions": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "问卷调查SurveyX API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
