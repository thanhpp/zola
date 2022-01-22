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
        "/account": {
            "get": {
                "description": "Get account if id is not specified, then get all account with page and size",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Get account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "account id",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Get success",
                        "schema": {
                            "$ref": "#/definitions/dto.GetMailAccountResp"
                        }
                    }
                }
            },
            "put": {
                "description": "Update account info (Need to specified account ID)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Update account info",
                "parameters": [
                    {
                        "description": "Update info",
                        "name": "updateReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateMailAccountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update OK",
                        "schema": {
                            "$ref": "#/definitions/dto.RespErr"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Create new account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "createReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateMailAccountReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Create OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateMailAccountResp"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete account by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Delete account by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "account id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "delete OK",
                        "schema": {
                            "$ref": "#/definitions/dto.RespErr"
                        }
                    }
                }
            }
        },
        "/delete": {
            "delete": {
                "description": "Delete mail by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mail"
                ],
                "summary": "Delete mail",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "mail ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "delete OK",
                        "schema": {
                            "$ref": "#/definitions/dto.RespErr"
                        }
                    }
                }
            }
        },
        "/mail": {
            "get": {
                "description": "Get mail, if id is not specified, get all by page and size",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mail"
                ],
                "summary": "Get mail",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "mail id",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Get OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.GetMailResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "mail": {
                                            "type": "array",
                                            "items": {
                                                "allOf": [
                                                    {
                                                        "$ref": "#/definitions/entity.Mail"
                                                    },
                                                    {
                                                        "type": "object",
                                                        "properties": {
                                                            "send_to": {
                                                                "type": "array",
                                                                "items": {
                                                                    "type": "string"
                                                                }
                                                            }
                                                        }
                                                    }
                                                ]
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "update email, need to specify id, update variables will update entire mail variables",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mail"
                ],
                "summary": "update email",
                "parameters": [
                    {
                        "description": "Update info",
                        "name": "updateReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateMailReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateMailResp"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new mail, will generate html to preview",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mail"
                ],
                "summary": "Create new mail",
                "parameters": [
                    {
                        "description": "create req",
                        "name": "createReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateNewMailReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Create OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateNewMailResp"
                        }
                    }
                }
            }
        },
        "/mail/send": {
            "post": {
                "description": "Send email",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mail"
                ],
                "summary": "Send email",
                "parameters": [
                    {
                        "description": "mail info",
                        "name": "SendReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SendEmailReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Send OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SendEmailResp"
                        }
                    }
                }
            }
        },
        "/template": {
            "get": {
                "description": "Get template, if id is not specified, then get all template with page and size",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "template"
                ],
                "summary": "Get template",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "template id",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Get OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetTemplateResp"
                        }
                    }
                }
            },
            "put": {
                "description": "Update template, need to specified templateID, will not validate template",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "template"
                ],
                "summary": "Update template",
                "parameters": [
                    {
                        "description": "update req",
                        "name": "updateReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateTemplateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update OK",
                        "schema": {
                            "$ref": "#/definitions/dto.RespErr"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new template",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "template"
                ],
                "summary": "Create new template",
                "parameters": [
                    {
                        "description": "Create info",
                        "name": "createReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateNewTemplateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Create OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateNewTemplateResp"
                        }
                    }
                }
            },
            "delete": {
                "description": "DeleteTemplate by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "template"
                ],
                "summary": "DeleteTemplate",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "templateID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "DeleteOK",
                        "schema": {
                            "$ref": "#/definitions/dto.RespErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateMailAccountReq": {
            "type": "object",
            "properties": {
                "account": {
                    "$ref": "#/definitions/entity.MailAccount"
                }
            }
        },
        "dto.CreateMailAccountResp": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "error": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "type": "integer"
                        },
                        "message": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "dto.CreateNewMailReq": {
            "type": "object",
            "properties": {
                "from_id": {
                    "type": "integer"
                },
                "subject": {
                    "type": "string"
                },
                "template_id": {
                    "type": "integer"
                },
                "to": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "variables": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Variable"
                    }
                }
            }
        },
        "dto.CreateNewMailResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "type": "integer"
                        },
                        "message": {
                            "type": "string"
                        }
                    }
                },
                "html": {
                    "type": "string"
                },
                "mail_id": {
                    "type": "integer"
                }
            }
        },
        "dto.CreateNewTemplateReq": {
            "type": "object",
            "properties": {
                "template": {
                    "$ref": "#/definitions/entity.Template"
                },
                "variables": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Variable"
                    }
                }
            }
        },
        "dto.CreateNewTemplateResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "type": "integer"
                        },
                        "message": {
                            "type": "string"
                        }
                    }
                },
                "template_id": {
                    "type": "integer"
                }
            }
        },
        "dto.GetMailAccountResp": {
            "type": "object",
            "properties": {
                "accounts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.MailAccount"
                    }
                },
                "error": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "type": "integer"
                        },
                        "message": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "dto.GetMailResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "type": "integer"
                        },
                        "message": {
                            "type": "string"
                        }
                    }
                },
                "mail": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Mail"
                    }
                }
            }
        },
        "dto.GetTemplateResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "type": "integer"
                        },
                        "message": {
                            "type": "string"
                        }
                    }
                },
                "template": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Template"
                    }
                }
            }
        },
        "dto.RespErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "type": "integer"
                        },
                        "message": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "dto.SendEmailReq": {
            "type": "object",
            "properties": {
                "mailID": {
                    "type": "integer"
                }
            }
        },
        "dto.SendEmailResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "type": "integer"
                        },
                        "message": {
                            "type": "string"
                        }
                    }
                },
                "html": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateMailAccountReq": {
            "type": "object",
            "properties": {
                "account": {
                    "$ref": "#/definitions/entity.MailAccount"
                },
                "account_id": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdateMailReq": {
            "type": "object",
            "properties": {
                "mail": {
                    "$ref": "#/definitions/entity.Mail"
                },
                "mail_id": {
                    "type": "integer"
                },
                "variables": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Variable"
                    }
                }
            }
        },
        "dto.UpdateMailResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "object",
                    "properties": {
                        "code": {
                            "type": "integer"
                        },
                        "message": {
                            "type": "string"
                        }
                    }
                },
                "html": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateTemplateReq": {
            "type": "object",
            "properties": {
                "template": {
                    "$ref": "#/definitions/entity.Template"
                },
                "templateID": {
                    "type": "integer"
                },
                "variables": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Variable"
                    }
                }
            }
        },
        "entity.Mail": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "send_at": {
                    "type": "string"
                },
                "send_to": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                },
                "templateID": {
                    "type": "integer"
                },
                "variables": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "entity.MailAccount": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "smtp_host": {
                    "type": "string"
                },
                "smtp_port": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.Template": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Variable": {
            "type": "object",
            "properties": {
                "default": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "require": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
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
