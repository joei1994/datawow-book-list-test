// Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/books": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get books",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books Actions"
                ],
                "summary": "Get books",
                "operationId": "books-get",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "Books IDs",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Book"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
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
                "description": "Create book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books Actions"
                ],
                "summary": "Create book",
                "operationId": "books-create",
                "parameters": [
                    {
                        "description": "Book title and author",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.BookPayload"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/api/books/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books Actions"
                ],
                "summary": "Get book",
                "operationId": "book-get",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books Actions"
                ],
                "summary": "Update book",
                "operationId": "books-update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Book title and author",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.BookPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
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
                "description": "Delete book",
                "tags": [
                    "Books Actions"
                ],
                "summary": "Delete book",
                "operationId": "books-delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Perform user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth Actions"
                ],
                "summary": "Authenticate a user",
                "operationId": "user-login",
                "parameters": [
                    {
                        "description": "User's credentials",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.LoginResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth Actions"
                ],
                "summary": "Register a user",
                "operationId": "user-register",
                "parameters": [
                    {
                        "description": "User's credentials",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "createdBy": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "genre": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "tag": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "year": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "requests.BookPayload": {
            "type": "object",
            "required": [
                "author",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "genre": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "year": {
                    "type": "string"
                }
            }
        },
        "requests.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "requests.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "responses.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "responses.LoginResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "exp": {
                    "type": "integer"
                },
                "refreshToken": {
                    "type": "string"
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Book List",
	Description:      "Book List",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
