package cache

import "banner-server/pkg/cache/cacheInMemory/model"

type Cache interface {
	Set(key model.KeyCache, value interface{})
	Get(key model.KeyCache) (interface{}, bool)
}
