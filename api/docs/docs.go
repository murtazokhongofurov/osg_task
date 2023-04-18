// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "2 term OSG",
        "contact": {
            "name": "Murtazoxon",
            "url": "https://t.me/murtazokhon_gofurov",
            "email": "gofurovmurtazoxon@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/employee": {
            "post": {
                "description": "Through this api, can create an employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Create new employee",
                "parameters": [
                    {
                        "description": "CreateEmployee",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EmployeeReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.FailureInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.FailureInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailureInfo"
                        }
                    }
                }
            }
        },
        "/employee/{id}": {
            "get": {
                "description": "Through this api, can GET an employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "GET an employee",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.FailureInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.FailureInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailureInfo"
                        }
                    }
                }
            }
        },
        "/employees": {
            "get": {
                "description": "Through this api, can GET employees",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "GET employees",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.FailureInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.FailureInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailureInfo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.EmployeeReq": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "profile_photo": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "models.FailureInfo": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {},
                "status_code": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7777",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Online Service Group Task api",
	Description:      "This is Task api server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
