package token_manager

import "sync"

type AuthTokens interface {
	GetToken() string
	GetRefreshToken() string
	SetToken(token string)
	SetRefreshToken(token string)
	GetPayload() map[string]interface{}
	SetPayload(payload map[string]interface{})
}
type authTokens struct {
	Token        string                 `json:"token"`
	Payload      map[string]interface{} `json:"payload"`
	RefreshToken string                 `json:"refresh_token"`
	mu           sync.RWMutex
}

func NewAuthTokens(token string, refreshToken string, payload map[string]interface{}) AuthTokens {
	return &authTokens{
		Token:        token,
		RefreshToken: refreshToken,
		Payload:      payload,
	}
}

func (a *authTokens) GetToken() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.Token
}

func (a *authTokens) GetRefreshToken() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.RefreshToken
}

func (a *authTokens) SetToken(token string) {
	a.mu.Lock()
	a.Token = token
	a.mu.Unlock()
}

func (a *authTokens) SetRefreshToken(token string) {
	a.mu.Lock()
	a.RefreshToken = token
	a.mu.Unlock()
}

func (a *authTokens) GetPayload() map[string]interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.Payload
}

func (a *authTokens) SetPayload(payload map[string]interface{}) {
	a.mu.Lock()
	a.Payload = payload
	a.mu.Unlock()
}
