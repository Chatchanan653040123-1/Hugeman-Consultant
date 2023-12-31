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
        "/create": {
            "post": {
                "description": "ID, Title,Date, and Status fields must be required\nThe Title field must not over 100 characters\nThe Status field must accept only IN_PROGRESS or COMPLETE status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Create a task",
                "parameters": [
                    {
                        "description": "Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.Response"
                        }
                    }
                }
            }
        },
        "/showSearchData": {
            "post": {
                "description": "To search the data by Title or Description fields",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Show searched data",
                "parameters": [
                    {
                        "description": "SearchRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.SearchRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.Response"
                        }
                    }
                }
            }
        },
        "/showSortData": {
            "post": {
                "description": "To sort the data by Title or Date or Status fields",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Show sorted data",
                "parameters": [
                    {
                        "description": "SortRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.SortRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.Response"
                        }
                    }
                }
            }
        },
        "/update/{id}": {
            "post": {
                "description": "To update Title, Description, Date, Image, and Status fields corresponding to the requirements from the CREATE feature",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Update a task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "services.Request": {
            "type": "object",
            "properties": {
                "Description": {
                    "type": "string"
                },
                "Image": {
                    "type": "string"
                },
                "Status": {
                    "type": "string"
                },
                "Title": {
                    "type": "string"
                }
            }
        },
        "services.Response": {
            "type": "object",
            "properties": {
                "Created At Date Time": {
                    "type": "string"
                },
                "Description": {
                    "type": "string"
                },
                "ID": {
                    "type": "string"
                },
                "Image": {
                    "type": "string"
                },
                "Status": {
                    "type": "string"
                },
                "Title": {
                    "type": "string"
                }
            }
        },
        "services.SearchRequest": {
            "type": "object",
            "properties": {
                "SearchKeyWord": {
                    "type": "string"
                }
            }
        },
        "services.SortRequest": {
            "type": "object",
            "properties": {
                "SortBy": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1",
	Host:             "localhost:5000",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "Backend Test(Golang) API",
	Description:      "This is a sample server for Backend Test API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
