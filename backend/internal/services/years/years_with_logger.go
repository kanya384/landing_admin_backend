package years

// Code generated by gowrap. DO NOT EDIT.
// template: https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/logrus
// gowrap: http://github.com/hexdigest/gowrap

//go:generate gowrap gen -p landing_admin_backend/internal/services/years -i Years -t https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/logrus -o years_with_logger.go -l ""

import (
	"context"
	"landing_admin_backend/internal/domain"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// YearsWithLogrus implements Years that is instrumented with logrus logger
type YearsWithLogrus struct {
	_log  *logrus.Entry
	_base Years
}

// NewYearsWithLogrus instruments an implementation of the Years with simple logging
func NewYearsWithLogrus(base Years, log *logrus.Entry) YearsWithLogrus {
	return YearsWithLogrus{
		_base: base,
		_log:  log,
	}
}

// Create implements Years
func (_d YearsWithLogrus) Create(ctx context.Context, year domain.Year) (err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":  ctx,
		"year": year})).Debug("YearsWithLogrus: calling Create")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Error("YearsWithLogrus: method Create returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Debug("YearsWithLogrus: method Create finished")
		}
	}()
	return _d._base.Create(ctx, year)
}

// Delete implements Years
func (_d YearsWithLogrus) Delete(ctx context.Context, yearID primitive.ObjectID) (err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":    ctx,
		"yearID": yearID})).Debug("YearsWithLogrus: calling Delete")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Error("YearsWithLogrus: method Delete returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Debug("YearsWithLogrus: method Delete finished")
		}
	}()
	return _d._base.Delete(ctx, yearID)
}

// Get implements Years
func (_d YearsWithLogrus) Get(ctx context.Context) (years []domain.Year, err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx": ctx})).Debug("YearsWithLogrus: calling Get")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"years": years,
				"err":   err})).Error("YearsWithLogrus: method Get returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"years": years,
				"err":   err})).Debug("YearsWithLogrus: method Get finished")
		}
	}()
	return _d._base.Get(ctx)
}

// Update implements Years
func (_d YearsWithLogrus) Update(ctx context.Context, year domain.Year) (err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":  ctx,
		"year": year})).Debug("YearsWithLogrus: calling Update")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Error("YearsWithLogrus: method Update returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Debug("YearsWithLogrus: method Update finished")
		}
	}()
	return _d._base.Update(ctx, year)
}
