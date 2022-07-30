package landing_content

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/services/memcache"
)

type LandingContentWithCache struct {
	cache memcache.Cache
	_base LandingContent
}

// NewPlansWithCache instruments an implementation of the Plans with simple logging
func NewLandingContentWithCache(base LandingContent, cache memcache.Cache) LandingContentWithCache {
	return LandingContentWithCache{
		_base: base,
		cache: cache,
	}
}

// GetActivePlans implements Plans
func (_d LandingContentWithCache) Get(ctx context.Context) (content domain.LandingContent, err error) {
	contentCache, err := _d.cache.GetAppContent()
	if err == nil {
		content = domain.LandingContent{
			Advantages:      contentCache["Adavanatges"].([]*domain.Advantage),
			AdvantagePhotos: contentCache["AdavanatgePhotos"].([]domain.AdvantagePhoto),
			Docs:            contentCache["Docs"].([]*domain.Doc),
			Editables:       contentCache["Editables"].([]*domain.Editable),
			Years:           contentCache["Years"].([]domain.Year),
			Months:          contentCache["Months"].([]domain.Month),
			Photos:          contentCache["Photos"].([]domain.HodPhoto),
			Plans:           contentCache["Plans"].([]*domain.Plan),
			Posters:         contentCache["Posters"].([]*domain.Poster),
			ProjectInfo:     contentCache["ProjectInfo"].([]*domain.ProjectInfo),
			Video:           contentCache["Video"].([]*domain.Video),
			Setting:         contentCache["Settings"].([]*domain.Setting),
			Title:           contentCache["Titles"].([]*domain.Title),
			Action:          contentCache["Action"].([]*domain.Action),
		}
		return
	}
	content, err = _d._base.Get(ctx)
	if err != nil {
		return
	}
	cacheContent := map[string]interface {
	}{
		"Adavanatges":      content.Advantages,
		"AdavanatgePhotos": content.AdvantagePhotos,
		"Docs":             content.Docs,
		"Editables":        content.Editables,
		"Years":            content.Years,
		"Months":           content.Months,
		"Photos":           content.Photos,
		"Plans":            content.Plans,
		"Posters":          content.Posters,
		"ProjectInfo":      content.ProjectInfo,
		"Video":            content.Video,
		"Settings":         content.Setting,
		"Titles":           content.Title,
		"Action":           content.Action,
	}
	_d.cache.SetAppContent(cacheContent)
	return
}
