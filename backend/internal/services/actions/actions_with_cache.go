package actions

// Code generated by gowrap. DO NOT EDIT.
// template: ../../generated/wrappers/cache
// gowrap: http://github.com/hexdigest/gowrap

//go:generate gowrap gen -p landing_admin_backend/internal/services/actions -i Actions -t ../../generated/wrappers/cache -o actions_with_cache.go -l ""

import (
	"context"
	"io"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/services/memcache"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ActionsWithCache implements Actions that is instrumented with logging
type ActionsWithCache struct {
	cache memcache.Cache
	_base Actions
}

// NewActionsWithCache instruments an implementation of the Actions with simple logging
func NewActionsWithCache(base Actions, cache memcache.Cache) ActionsWithCache {
	return ActionsWithCache{
		_base: base,
		cache: cache,
	}
}

// Create implements Actions
func (_d ActionsWithCache) Create(ctx context.Context, action domain.Action, file io.ReadCloser) (actionRes domain.Action, err error) {
	defer _d.cache.Delete("Actions")

	return _d._base.Create(ctx, action, file)
}

// Delete implements Actions
func (_d ActionsWithCache) Delete(ctx context.Context, actionID primitive.ObjectID) (err error) {
	defer _d.cache.Delete("Actions")

	return _d._base.Delete(ctx, actionID)
}

// Get implements Actions
func (_d ActionsWithCache) Get(ctx context.Context, filter map[string]interface{}) (actions []*domain.Action, err error) {
	cached, ok := _d.cache.Get("Actions")
	if ok {

		return cached.Value.([]*domain.Action), err

	}
	res, err := _d._base.Get(ctx, filter)
	if err != nil {
		return
	}
	_d.cache.Set("Actions", res)
	return res, err
}

// GetByID implements Actions
func (_d ActionsWithCache) GetByID(ctx context.Context, id primitive.ObjectID) (action domain.Action, err error) {
	return _d._base.GetByID(ctx, id)
}

// Update implements Actions
func (_d ActionsWithCache) Update(ctx context.Context, action domain.Action, file interface{}) (actionRes domain.Action, err error) {
	defer _d.cache.Delete("Actions")

	return _d._base.Update(ctx, action, file)
}

// UpdateOrder implements Actions
func (_d ActionsWithCache) UpdateOrder(ctx context.Context, first domain.UpdateOrder, second domain.UpdateOrder) (err error) {
	defer _d.cache.Delete("Actions")

	return _d._base.UpdateOrder(ctx, first, second)
}