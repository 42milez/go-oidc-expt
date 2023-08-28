// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "TBD",
        "contact": {
            "name": "TBD",
            "url": "TBD",
            "email": "TBD"
        },
        "license": {
            "name": "MIT",
            "url": "TBD"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/authenticate": {
            "post": {
                "description": "TBD",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "TBD",
                "operationId": "Authenticate.ServeHTTP",
                "parameters": [
                    {
                        "description": "Username",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AuthenticateResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/authorization": {
            "get": {
                "description": "TBD",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authorization"
                ],
                "summary": "TBD",
                "operationId": "AuthorizeGet.ServeHTTP",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TBD",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "TBD",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "TBD",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authorization"
                ],
                "summary": "TBD",
                "operationId": "AuthorizePost.ServeHTTP",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TBD",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "TBD",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AuthenticateResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "model.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    },
    "tags": [
        {
            "description": "TBD",
            "name": "TBD",
            "externalDocs": {
                "description": "TBD",
                "url": "TBD"
            }
        }
    ],
    "externalDocs": {
        "description": "TBD",
        "url": "TBD"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "TBD",
	BasePath:         "/v1",
	Schemes:          []string{"http", "https"},
	Title:            "go-oidc-server",
	Description:      "TBD",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}