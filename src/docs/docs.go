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
        "/v1/health/": {
            "get": {
                "description": "Health Check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/order-matching_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/order-matching_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/orders/cancel-order": {
            "post": {
                "description": "Cancel an order by its ID and market",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Cancel an order",
                "parameters": [
                    {
                        "description": "CancelOrderRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order-matching_api_dto.CancelOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order canceled successfully",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Market or Order not found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/v1/orders/create": {
            "post": {
                "description": "Create a new order of type limit or market",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Create an order",
                "parameters": [
                    {
                        "description": "OrderRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order-matching_api_dto.OrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order processed successfully",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Invalid input or processing error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/v1/orders/get-order": {
            "post": {
                "description": "Retrieve order details by its ID and market",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get an order by ID",
                "parameters": [
                    {
                        "description": "GetOrderRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order-matching_api_dto.GetOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order Find",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Market or Order not found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/v1/orders/orderbook": {
            "post": {
                "description": "Retrieve the depth of the order book for a specific market",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get order book depth",
                "parameters": [
                    {
                        "description": "GetOrderBookRequest",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order-matching_api_dto.GetOrderBookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OrderBook depth retrieved",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Market not found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gin.H": {
            "type": "object",
            "additionalProperties": {}
        },
        "order-matching_api_dto.CancelOrderRequest": {
            "type": "object",
            "required": [
                "market",
                "order_id"
            ],
            "properties": {
                "market": {
                    "type": "string"
                },
                "order_id": {
                    "type": "string"
                }
            }
        },
        "order-matching_api_dto.GetOrderBookRequest": {
            "type": "object",
            "required": [
                "market"
            ],
            "properties": {
                "market": {
                    "type": "string"
                }
            }
        },
        "order-matching_api_dto.GetOrderRequest": {
            "type": "object",
            "required": [
                "market",
                "order_id"
            ],
            "properties": {
                "market": {
                    "type": "string"
                },
                "order_id": {
                    "type": "string"
                }
            }
        },
        "order-matching_api_dto.OrderRequest": {
            "type": "object",
            "required": [
                "amount",
                "market",
                "order_id",
                "order_kind",
                "order_type"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "market": {
                    "type": "string"
                },
                "order_id": {
                    "type": "string"
                },
                "order_kind": {
                    "type": "string"
                },
                "order_type": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "order-matching_api_helper.BaseHttpResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "result": {},
                "resultCode": {
                    "$ref": "#/definitions/order-matching_api_helper.ResultCode"
                },
                "success": {
                    "type": "boolean"
                },
                "validationErrors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/order-matching_api_validation.ValidationError"
                    }
                }
            }
        },
        "order-matching_api_helper.ResultCode": {
            "type": "integer",
            "enum": [
                200,
                400,
                401,
                403,
                404,
                429,
                42901,
                500,
                500
            ],
            "x-enum-comments": {
                "AuthError": "Authentication error",
                "CustomRecovery": "Custom recovery error",
                "ForbiddenError": "Access forbidden",
                "InternalError": "Internal server error",
                "LimiterError": "Too many requests (rate limit)",
                "NotFoundError": "Resource not found",
                "OtpLimiterError": "Too many OTP requests (rate limit for OTP)",
                "Success": "Request successful",
                "ValidationError": "Validation error"
            },
            "x-enum-varnames": [
                "Success",
                "ValidationError",
                "AuthError",
                "ForbiddenError",
                "NotFoundError",
                "LimiterError",
                "OtpLimiterError",
                "CustomRecovery",
                "InternalError"
            ]
        },
        "order-matching_api_validation.ValidationError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "property": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}