// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Justine Bachelard.",
            "email": "justine.bachelard@ext.uca.fr"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/musics": {
            "get": {
                "description": "Get musics.",
                "tags": [
                    "musics"
                ],
                "summary": "Get musics.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Music"
                            }
                        }
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            },
            "post": {
                "description": "Create a music.",
                "tags": [
                    "musics"
                ],
                "summary": "Create a music.",
                "parameters": [
                    {
                        "description": "Music object",
                        "name": "body",
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Music"
                            }
                        }
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            }
        },
        "/musics/{id}": {
            "get": {
                "description": "Get a music.",
                "tags": [
                    "musics"
                ],
                "summary": "Get a music.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Music UUID formatted ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Music"
                        }
                    },
                    "422": {
                        "description": "Cannot parse id"
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Music": {
            "type": "object",
            "properties": {
                "albumId": {
                    "type": "string"
                },
                "artistId": {
                    "type": "string"
                },
                "genreId": {
                    "type": "string"
                },
                "id": {
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
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "middleware/example",
	Description:      "API to manage collections.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
