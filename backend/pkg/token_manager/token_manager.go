package token_manager

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenManager interface {
	GenerateAuthTokens(id string, name string, role string) (authTokens AuthTokens, err error)
	ParseAuthTokens(token string, refreshToken string) (authTokens AuthTokens, tokenExpired bool, err error)
	RefreshTokens(authTokens AuthTokens) (authTokensRes AuthTokens, err error)
}

type tokenManager struct {
	secret     string
	tokenTTL   time.Duration
	refreshTTL time.Duration
}

func NewTokenManager(secret string, tokenTTL time.Duration, refreshTTL time.Duration) TokenManager {
	return &tokenManager{
		secret:     secret,
		tokenTTL:   tokenTTL,
		refreshTTL: refreshTTL,
	}
}

func (t *tokenManager) GenerateAuthTokens(id string, name string, role string) (authTokens AuthTokens, err error) {
	payload := map[string]interface{}{
		"id":   id,
		"name": name,
		"role": role,
	}
	token, err := GenerateToken(t.secret, t.tokenTTL, payload)
	if err != nil {
		return
	}
	refreshToken, err := GenerateToken(t.secret, t.tokenTTL, nil)
	if err != nil {
		return
	}
	authTokens = NewAuthTokens(token, refreshToken, payload)
	return
}

func (t *tokenManager) ParseAuthTokens(token string, refreshToken string) (authTokens AuthTokens, tokenExpired bool, err error) {
	// check token
	payload, err := CheckToken(token, t.secret)
	if err != nil && err.Error() != ErrTokenIsExpired {
		return nil, false, errors.New(ErrNotValidToken)
	}
	authTokens = NewAuthTokens(token, "", payload)

	//if token is expired, check refresh
	if err != nil && err.Error() == ErrTokenIsExpired {
		err = nil
		tokenExpired = true
		_, err := CheckToken(refreshToken, t.secret)
		if err != nil {
			return nil, true, errors.New(ErrNotValidRefreshToken)
		}
		//if booth tokens are expired, return error
		if err != nil && err.Error() == ErrTokenIsExpired {
			return nil, true, errors.New(ErrTokensAreExpired)
		}
		authTokens.SetRefreshToken(refreshToken)
	}

	return
}

func (t *tokenManager) RefreshTokens(authTokens AuthTokens) (authTokensRes AuthTokens, err error) {
	token, err := GenerateToken(t.secret, t.tokenTTL, authTokens.GetPayload())
	if err != nil {
		return
	}
	refreshToken, err := GenerateToken(t.secret, t.tokenTTL, nil)
	if err != nil {
		return
	}
	authTokensRes = NewAuthTokens(token, refreshToken, authTokens.GetPayload())
	return
}

func CheckToken(token string, secret string) (payload map[string]interface{}, err error) {
	//проверям шифрование
	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(ErrWrongSingingMethod)
		}
		return []byte(secret), nil
	})
	fmt.Println(err)
	fmt.Println("parsed")

	if err != nil && (err.Error() == ErrWrongSingingMethod || err.Error() == ErrInvalidNumberOfSegments) {
		return nil, err
	}
	claims, ok := parsedToken.Claims.(*TokenClaims)
	payload = claims.Payload
	if !ok {
		return payload, errors.New(ErrNotValidStruct)
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return payload, errors.New(ErrTokenIsExpired)
	}
	return
}

func GenerateToken(secret string, timeToLive time.Duration, payload map[string]interface{}) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(timeToLive).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		payload,
	}).SignedString([]byte(secret))
}
