package users

// Code generated by gowrap. DO NOT EDIT.
// template: https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/logrus
// gowrap: http://github.com/hexdigest/gowrap

//go:generate gowrap gen -p landing_admin_backend/internal/services/users -i Users -t https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/logrus -o users_with_logger.go -l ""

import (
	"context"
	"landing_admin_backend/internal/domain"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UsersWithLogrus implements Users that is instrumented with logrus logger
type UsersWithLogrus struct {
	_log  *logrus.Entry
	_base Users
}

// NewUsersWithLogrus instruments an implementation of the Users with simple logging
func NewUsersWithLogrus(base Users, log *logrus.Entry) UsersWithLogrus {
	return UsersWithLogrus{
		_base: base,
		_log:  log,
	}
}

// Create implements Users
func (_d UsersWithLogrus) Create(ctx context.Context, user domain.User) (err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":  ctx,
		"user": user})).Debug("UsersWithLogrus: calling Create")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Error("UsersWithLogrus: method Create returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Debug("UsersWithLogrus: method Create finished")
		}
	}()
	return _d._base.Create(ctx, user)
}

// Delete implements Users
func (_d UsersWithLogrus) Delete(ctx context.Context, userID primitive.ObjectID) (err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":    ctx,
		"userID": userID})).Debug("UsersWithLogrus: calling Delete")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Error("UsersWithLogrus: method Delete returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Debug("UsersWithLogrus: method Delete finished")
		}
	}()
	return _d._base.Delete(ctx, userID)
}

// Get implements Users
func (_d UsersWithLogrus) Get(ctx context.Context) (users []*domain.UserInfo, err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx": ctx})).Debug("UsersWithLogrus: calling Get")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"users": users,
				"err":   err})).Error("UsersWithLogrus: method Get returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"users": users,
				"err":   err})).Debug("UsersWithLogrus: method Get finished")
		}
	}()
	return _d._base.Get(ctx)
}

// Update implements Users
func (_d UsersWithLogrus) Update(ctx context.Context, user domain.User) (err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":  ctx,
		"user": user})).Debug("UsersWithLogrus: calling Update")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Error("UsersWithLogrus: method Update returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Debug("UsersWithLogrus: method Update finished")
		}
	}()
	return _d._base.Update(ctx, user)
}
