// Code generated by go-swagger; DO NOT EDIT.

package generated

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Landing Backend REST",
    "version": "1.0.0"
  },
  "host": "backend-service",
  "paths": {
    "/login": {
      "post": {
        "summary": "authentication path",
        "parameters": [
          {
            "name": "params",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/AuthenticateRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "регистрация успешна",
            "schema": {
              "$ref": "#/definitions/AuthenticateResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/AuthenticateResponse"
            }
          },
          "403": {
            "description": "Authentication Fail",
            "schema": {
              "$ref": "#/definitions/AuthenticateResponse"
            }
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/users": {
      "put": {
        "summary": "create user",
        "parameters": [
          {
            "name": "params",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/UserCreateRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "пользователь успешно создан",
            "schema": {
              "$ref": "#/definitions/ResultResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ResultResponse"
            }
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    }
  },
  "definitions": {
    "AuthenticateRequest": {
      "type": "object",
      "properties": {
        "login": {
          "type": "string",
          "example": "login"
        },
        "pass": {
          "type": "string",
          "example": "password"
        }
      }
    },
    "AuthenticateResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string",
          "example": "error"
        },
        "refresh_token": {
          "type": "string",
          "example": "fewerHHsasqw122231"
        },
        "token": {
          "type": "string",
          "example": "asdkjkzxcqw1290090"
        }
      }
    },
    "ResultResponse": {
      "type": "object",
      "properties": {
        "msg": {
          "type": "string",
          "example": "пользователь успешно создан"
        }
      }
    },
    "UserCreateRequest": {
      "type": "object",
      "properties": {
        "login": {
          "type": "string",
          "example": "login"
        },
        "name": {
          "type": "string",
          "example": "пользователь"
        },
        "pass": {
          "type": "string",
          "example": "password"
        },
        "role": {
          "type": "string",
          "example": "роль"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Landing Backend REST",
    "version": "1.0.0"
  },
  "host": "backend-service",
  "paths": {
    "/login": {
      "post": {
        "summary": "authentication path",
        "parameters": [
          {
            "name": "params",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/AuthenticateRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "регистрация успешна",
            "schema": {
              "$ref": "#/definitions/AuthenticateResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/AuthenticateResponse"
            }
          },
          "403": {
            "description": "Authentication Fail",
            "schema": {
              "$ref": "#/definitions/AuthenticateResponse"
            }
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/users": {
      "put": {
        "summary": "create user",
        "parameters": [
          {
            "name": "params",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/UserCreateRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "пользователь успешно создан",
            "schema": {
              "$ref": "#/definitions/ResultResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ResultResponse"
            }
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    }
  },
  "definitions": {
    "AuthenticateRequest": {
      "type": "object",
      "properties": {
        "login": {
          "type": "string",
          "example": "login"
        },
        "pass": {
          "type": "string",
          "example": "password"
        }
      }
    },
    "AuthenticateResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string",
          "example": "error"
        },
        "refresh_token": {
          "type": "string",
          "example": "fewerHHsasqw122231"
        },
        "token": {
          "type": "string",
          "example": "asdkjkzxcqw1290090"
        }
      }
    },
    "ResultResponse": {
      "type": "object",
      "properties": {
        "msg": {
          "type": "string",
          "example": "пользователь успешно создан"
        }
      }
    },
    "UserCreateRequest": {
      "type": "object",
      "properties": {
        "login": {
          "type": "string",
          "example": "login"
        },
        "name": {
          "type": "string",
          "example": "пользователь"
        },
        "pass": {
          "type": "string",
          "example": "password"
        },
        "role": {
          "type": "string",
          "example": "роль"
        }
      }
    }
  }
}`))
}
