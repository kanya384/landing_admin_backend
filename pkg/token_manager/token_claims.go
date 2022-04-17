package token_manager

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	jwt.StandardClaims
	Payload map[string]interface{} `json:"payload"`
}
