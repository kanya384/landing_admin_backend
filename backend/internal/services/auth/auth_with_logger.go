package auth

// Code generated by gowrap. DO NOT EDIT.
// template: https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/logrus
// gowrap: http://github.com/hexdigest/gowrap

//go:generate gowrap gen -p landing_admin_backend/internal/services/auth -i Auth -t https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/logrus -o auth_with_logger.go -l ""

import (
	"context"
	"landing_admin_backend/pkg/token_manager"

	"github.com/sirupsen/logrus"
)

// AuthWithLogrus implements Auth that is instrumented with logrus logger
type AuthWithLogrus struct {
	_log  *logrus.Entry
	_base Auth
}

// NewAuthWithLogrus instruments an implementation of the Auth with simple logging
func NewAuthWithLogrus(base Auth, log *logrus.Entry) AuthWithLogrus {
	return AuthWithLogrus{
		_base: base,
		_log:  log,
	}
}

// Authenticate implements Auth
func (_d AuthWithLogrus) Authenticate(ctx context.Context, login string, pass string) (tokens token_manager.AuthTokens, err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":   ctx,
		"login": login,
		"pass":  pass})).Debug("AuthWithLogrus: calling Authenticate")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"tokens": tokens,
				"err":    err})).Error("AuthWithLogrus: method Authenticate returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"tokens": tokens,
				"err":    err})).Debug("AuthWithLogrus: method Authenticate finished")
		}
	}()
	return _d._base.Authenticate(ctx, login, pass)
}