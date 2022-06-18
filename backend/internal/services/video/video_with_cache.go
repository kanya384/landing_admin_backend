package video

// Code generated by gowrap. DO NOT EDIT.
// template: ../../generated/wrappers/cache
// gowrap: http://github.com/hexdigest/gowrap

//go:generate gowrap gen -p landing_admin_backend/internal/services/video -i Video -t ../../generated/wrappers/cache -o video_with_cache.go -l ""

import (
	"context"
	"io"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/services/memcache"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VideoWithCache implements Video that is instrumented with logging
type VideoWithCache struct {
	cache memcache.Cache
	_base Video
}

// NewVideoWithCache instruments an implementation of the Video with simple logging
func NewVideoWithCache(base Video, cache memcache.Cache) VideoWithCache {
	return VideoWithCache{
		_base: base,
		cache: cache,
	}
}

// Create implements Video
func (_d VideoWithCache) Create(ctx context.Context, video domain.Video, file io.ReadCloser, fileName string) (v1 domain.Video, err error) {
	defer _d.cache.Delete("Video")

	return _d._base.Create(ctx, video, file, fileName)
}

// Delete implements Video
func (_d VideoWithCache) Delete(ctx context.Context, videoID primitive.ObjectID) (err error) {
	defer _d.cache.Delete("Video")

	return _d._base.Delete(ctx, videoID)
}

// Get implements Video
func (_d VideoWithCache) Get(ctx context.Context) (videos []*domain.Video, err error) {
	cached, ok := _d.cache.Get("Video")
	if ok {

		return cached.Value.([]*domain.Video), err

	}
	res, err := _d._base.Get(ctx)
	if err != nil {
		return
	}
	_d.cache.Set("Video", res)
	return res, err
}

// Update implements Video
func (_d VideoWithCache) Update(ctx context.Context, video domain.Video, file interface{}, fileName string) (v1 domain.Video, err error) {
	defer _d.cache.Delete("Video")

	return _d._base.Update(ctx, video, file, fileName)
}
