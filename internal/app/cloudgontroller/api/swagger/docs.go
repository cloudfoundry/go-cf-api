// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package swagger

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
        "/buildpacks": {
            "get": {
                "description": "Retrieve all buildpacks the user has access to.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Buildpacks"
                ],
                "summary": "Buildpacks List buildpacks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/presenter.BuildpackResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.HTTPError"
                        }
                    }
                }
            }
        },
        "/buildpacks/{guid}": {
            "get": {
                "description": "Retrieve all buildpacks the user has access to.",
                "tags": [
                    "Buildpacks"
                ],
                "summary": "Show a buildpack",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Buildpack GUID",
                        "name": "guid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.BuildpackResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "presenter.BuildpackResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "enabled": {
                    "type": "boolean"
                },
                "filename": {
                    "type": "string"
                },
                "guid": {
                    "type": "string"
                },
                "links": {
                    "type": "object",
                    "properties": {
                        "self": {
                            "type": "object",
                            "properties": {
                                "href": {
                                    "type": "string"
                                }
                            }
                        },
                        "upload": {
                            "type": "object",
                            "properties": {
                                "href": {
                                    "type": "string"
                                },
                                "method": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                },
                "locked": {
                    "type": "boolean"
                },
                "metadata": {
                    "$ref": "#/definitions/presenter.Metadata"
                },
                "name": {
                    "type": "string"
                },
                "position": {
                    "type": "integer"
                },
                "stack": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "presenter.Metadata": {
            "type": "object",
            "properties": {
                "annotations": {
                    "type": "object"
                },
                "labels": {
                    "type": "object"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    },
    "x-extension-openapi": {
        "example": "value on a json format"
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
	Version:     "3.0.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v3",
	Schemes:     []string{},
	Title:       "CloudGontroller API",
	Description: "CAPI V3 Compatible API with blazing fast backend.",
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
