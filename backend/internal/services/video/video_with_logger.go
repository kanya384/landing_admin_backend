package video

// Code generated by gowrap. DO NOT EDIT.
// template: https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/logrus
// gowrap: http://github.com/hexdigest/gowrap

//go:generate gowrap gen -p landing_admin_backend/internal/services/video -i Video -t https://raw.githubusercontent.com/hexdigest/gowrap/bd05dcaf6963696b62ac150a98a59674456c6c53/templates/logrus -o video_with_logger.go -l ""

import (
	"context"
	"io"
	"landing_admin_backend/internal/domain"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VideoWithLogrus implements Video that is instrumented with logrus logger
type VideoWithLogrus struct {
	_log  *logrus.Entry
	_base Video
}

// NewVideoWithLogrus instruments an implementation of the Video with simple logging
func NewVideoWithLogrus(base Video, log *logrus.Entry) VideoWithLogrus {
	return VideoWithLogrus{
		_base: base,
		_log:  log,
	}
}

// Create implements Video
func (_d VideoWithLogrus) Create(ctx context.Context, video domain.Video, file io.ReadCloser, fileName string) (v1 domain.Video, err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":      ctx,
		"video":    video,
		"file":     file,
		"fileName": fileName})).Debug("VideoWithLogrus: calling Create")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"v1":  v1,
				"err": err})).Error("VideoWithLogrus: method Create returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"v1":  v1,
				"err": err})).Debug("VideoWithLogrus: method Create finished")
		}
	}()
	return _d._base.Create(ctx, video, file, fileName)
}

// Delete implements Video
func (_d VideoWithLogrus) Delete(ctx context.Context, videoID primitive.ObjectID) (err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":     ctx,
		"videoID": videoID})).Debug("VideoWithLogrus: calling Delete")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Error("VideoWithLogrus: method Delete returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"err": err})).Debug("VideoWithLogrus: method Delete finished")
		}
	}()
	return _d._base.Delete(ctx, videoID)
}

// Get implements Video
func (_d VideoWithLogrus) Get(ctx context.Context) (videos []*domain.Video, err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx": ctx})).Debug("VideoWithLogrus: calling Get")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"videos": videos,
				"err":    err})).Error("VideoWithLogrus: method Get returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"videos": videos,
				"err":    err})).Debug("VideoWithLogrus: method Get finished")
		}
	}()
	return _d._base.Get(ctx)
}

// Update implements Video
func (_d VideoWithLogrus) Update(ctx context.Context, video domain.Video, file interface{}, fileName string) (v1 domain.Video, err error) {
	_d._log.WithFields(logrus.Fields(map[string]interface{}{
		"ctx":      ctx,
		"video":    video,
		"file":     file,
		"fileName": fileName})).Debug("VideoWithLogrus: calling Update")
	defer func() {
		if err != nil {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"v1":  v1,
				"err": err})).Error("VideoWithLogrus: method Update returned an error")
		} else {
			_d._log.WithFields(logrus.Fields(map[string]interface{}{
				"v1":  v1,
				"err": err})).Debug("VideoWithLogrus: method Update finished")
		}
	}()
	return _d._base.Update(ctx, video, file, fileName)
}
