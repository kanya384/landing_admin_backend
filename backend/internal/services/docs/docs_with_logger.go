package docs

// Code generated by gowrap. DO NOT EDIT.
// template: https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/logrus
// gowrap: http://github.com/hexdigest/gowrap

//go:generate gowrap gen -p landing_admin_backend/internal/services/docs -i Docs -t https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/logrus -o docs_with_logger.go -l ""

import (
	"context"
	"io"
	"landing_admin_backend/internal/domain"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DocsWithLogrus implements Docs that is instrumented with logrus logger
type DocsWithLogrus struct {
	_log  *logrus.Entry
	_base Docs
}

// NewDocsWithLogrus instruments an implementation of the Docs with simple logging
func NewDocsWithLogrus(base Docs, log *logrus.Entry) DocsWithLogrus {
	return DocsWithLogrus{
		_base: base,
		_log:  log,
	}
}

// Create implements Docs
func (_d DocsWithLogrus) Create(ctx context.Context, doc domain.Doc, file io.ReadCloser, fileName string) (d1 domain.Doc, err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":      ctx,
		"doc":      doc,
		"file":     file,
		"fileName": fileName})).Debug("DocsWithLogrus: calling Create")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"d1":  d1,
				"err": err})).Error("DocsWithLogrus: method Create returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"d1":  d1,
				"err": err})).Debug("DocsWithLogrus: method Create finished")
		}
	}()
	return _d._base.Create(ctx, doc, file, fileName)
}

// Delete implements Docs
func (_d DocsWithLogrus) Delete(ctx context.Context, docID primitive.ObjectID) (err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":   ctx,
		"docID": docID})).Debug("DocsWithLogrus: calling Delete")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Error("DocsWithLogrus: method Delete returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Debug("DocsWithLogrus: method Delete finished")
		}
	}()
	return _d._base.Delete(ctx, docID)
}

// Get implements Docs
func (_d DocsWithLogrus) Get(ctx context.Context) (docs []*domain.Doc, err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx": ctx})).Debug("DocsWithLogrus: calling Get")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"docs": docs,
				"err":  err})).Error("DocsWithLogrus: method Get returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"docs": docs,
				"err":  err})).Debug("DocsWithLogrus: method Get finished")
		}
	}()
	return _d._base.Get(ctx)
}

// Update implements Docs
func (_d DocsWithLogrus) Update(ctx context.Context, doc domain.Doc, file interface{}, fileName string) (d1 domain.Doc, err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":      ctx,
		"doc":      doc,
		"file":     file,
		"fileName": fileName})).Debug("DocsWithLogrus: calling Update")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"d1":  d1,
				"err": err})).Error("DocsWithLogrus: method Update returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"d1":  d1,
				"err": err})).Debug("DocsWithLogrus: method Update finished")
		}
	}()
	return _d._base.Update(ctx, doc, file, fileName)
}