package token_manager_test

import (
	"landing_admin_backend/pkg/token_manager"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const TOKEN_SECRET_FOR_TEST = "secret"

func TestTokenManagerTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, &TokenManagerTestSuite{})
}

type TokenManagerTestSuite struct {
	suite.Suite
	tokenManager token_manager.TokenManager
}

func (s *TokenManagerTestSuite) SetupSuite() {
	s.tokenManager = token_manager.NewTokenManager(TOKEN_SECRET_FOR_TEST, time.Minute, time.Hour)
}

func (s *TokenManagerTestSuite) TestGenerateAuthTokens() {
	authTokens, err := s.tokenManager.GenerateAuthTokens(primitive.NewObjectID().Hex(), "test", "admin")
	if err != nil {
		s.FailNow("error creating auth tokens", err)
		return
	}

	s.NotEmpty(authTokens.GetToken())
	s.NotEmpty(authTokens.GetRefreshToken())
}

func (s *TokenManagerTestSuite) TestParseAuthTokens() {
	//generate auth tokens, for test purposes
	authTokens, err := s.tokenManager.GenerateAuthTokens(primitive.NewObjectID().Hex(), "test", "admin")
	if err != nil {
		s.FailNow("error creating auth tokens", err)
		return
	}

	//data for expired
	id := primitive.NewObjectID().Hex()
	authTokensExpiredToken, err := s.tokenManager.GenerateAuthTokens(id, "test", "admin")
	if err != nil {
		s.FailNow("error creating auth tokens for expired", err)
		return
	}
	tkn, err := token_manager.GenerateToken(TOKEN_SECRET_FOR_TEST, time.Duration(-time.Minute), map[string]interface{}{"id": id, "name": "test", "role": "admin"})
	if err != nil {
		s.FailNow("error generation token", err)
		return
	}
	authTokensExpiredToken.SetToken(tkn)

	cases := map[string]struct {
		input   token_manager.AuthTokens
		payload map[string]interface{}
		expired bool
		err     interface{}
	}{
		"success": {
			input:   authTokens,
			payload: authTokens.GetPayload(),
			expired: false,
			err:     nil,
		},
		"expired only token": {
			input:   authTokensExpiredToken,
			payload: authTokensExpiredToken.GetPayload(),
			expired: true,
			err:     nil,
		},
	}

	for name, cs := range cases {
		s.Run(name, func() {
			auth, expired, err := s.tokenManager.ParseAuthTokens(cs.input.GetToken(), cs.input.GetRefreshToken())
			if auth != nil {
				s.Equal(cs.payload, auth.GetPayload())
			}
			s.Equal(cs.expired, expired)
			s.Equal(cs.err, err)
		})
	}
}

func (s *TokenManagerTestSuite) TestRefreshTokens() {
	//generate auth tokens, for test purposes
	authTokens, err := s.tokenManager.GenerateAuthTokens(primitive.NewObjectID().Hex(), "test", "admin")
	if err != nil {
		s.FailNow("error creating auth tokens", err)
		return
	}

	cases := map[string]struct {
		input token_manager.AuthTokens
		err   interface{}
	}{
		"success": {
			input: authTokens,
			err:   nil,
		},
	}

	for name, cs := range cases {
		s.Run(name, func() {
			_, err := s.tokenManager.RefreshTokens(cs.input)
			s.Equal(cs.err, err)
		})
	}
}
