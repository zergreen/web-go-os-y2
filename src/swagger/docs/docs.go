// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
  "swagger": "2.0",
  "info": {
    "description": "",
    "version": "",
    "title": "Alumni-KMITL"
  },
  "basePath": "{{.BasePath}}",
  "host": "localhost:8080",
  "tags": [
    {
      "name": "Alumni-KMITL"
    }
  ],
  "paths": {
    "/api/createUser": {
      "post": {
        "tags": [
          "Alumni-KMITL"
        ],
        "summary": "add my address",
        "description": "ใส่ข้อมูลของคุณ",
        "operationId": "createUser",
        "consumes": [
          "application/json"
          ],
          "parameters": [
            {
            "in": "body",
            "name": "body",
            "description": "input",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                  "id": {
                  "type": "integer",
                  "example": 63050173
                },
                "first_name": {
                  "type": "string",
                  "example": "John"
                },
                "last_name": {
                  "type": "string",
                  "example": "Ocone"
                },
                "district": {
                  "type": "string",
                  "example": "เขตบางเขน"
                }
              }
            }
          }
          ],
          "responses": {
            "200": {
              "description": "Success to create your account",
              "schema": {
              "properties": {
                "statusCode": {
                  "type": "string",
                  "example": 200
                },
                "message": {
                  "type": "string",
                  "example": "Success to create your account"
                }
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "properties": {
                "statusCode": {
                  "type": "string",
                  "example": 400
                },
                "message": {
                  "type": "string",
                  "example": "Bad Request"
                }
              }
            }
          },
          "500": {
            "description": "unAuthenticate",
            "schema": {
              "properties": {
                "statusCode": {
                  "type": "string",
                  "example": 500
                },
                "message": {
                  "type": "string",
                  "example": "fail connect database,"
                }
              }
            }
          }
        }
      }
    },
    "/api/info": {
      "post": {
        "tags": [
          "Alumni-KMITL"
        ],
        "summary": "find your address",
        "description": "ใส่ข้อมูลของคุณ",
        "operationId": "InfoEndpoint",
        "consumes": [
          "application/json"
        ],
       
        "parameters": [
          {
              "in": "body",
              "name": "body",
              "description": "input",
              "required": true,
              "schema": {
        "type": "object",
        "properties": {
          "id": {
          "type": "integer",
          "example": 63050203
          },
          "first_name": {
          "type": "string",
          "example": "Sumate"
          },
          "last_name": {
          "type": "string",
          "example": "Kaewjai"
          },
          "district": {
          "type": "string",
          "example":"เขตคลองสามวา"
          },
        }
        }
          },
        ],
        "responses": {
          "200": {
            "description": "Success to add your address",
            "schema": {
              "properties": {
                "statusCode": {
                  "type": "string",
                  "example": 200
                },
                "message": {
                  "type": "string",
                  "example": "Success to add your address"
                }
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "properties": {
                "statusCode": {
                  "type": "string",
                  "example": 400
                },
                "message": {
                  "type": "string",
                  "example": "Bad Request"
                }
              }
            }
          },
          "500": {
            "description": "unAuthenticate",
            "schema": {
              "properties": {
                "statusCode": {
                  "type": "string",
                  "example": 500
                },
                "message": {
                  "type": "string",
                  "example": "fail connect database"
                }
              }
            }
          }
        }
      }
    },
    "/api/showGeography": {
      "post": {
        "tags": [
          "Alumni-KMITL"
          ],
          "summary": "add my address",
        "description": "ใช้เพื่อเเสดงการกระจายเชิงภูมิภาค",
        "operationId": "showGeography",
        "responses": {
          "200": {
            "description": "Success to show geofraphy",
            "schema": {
              "properties": {
                "statusCode": {
                  "type": "string",
                  "example": 200
                },
                "message": {
                  "type": "string",
                  "example": "Success to show geofraphy"
                }
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "properties": {
                "statusCode": {
                  "type": "string",
                  "example": 400
                },
                "message": {
                  "type": "string",
                  "example": "Bad Request"
                }
              }
            }
          },
          "500": {
            "description": "unAuthenticate",
            "schema": {
              "properties": {
                "statusCode": {
                  "type": "string",
                  "example": 500
                },
                "message": {
                  "type": "string",
                  "example": "fail connect database,"
                }
              }
            }
          }
        }
      }
    }
  }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Find My Address",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}