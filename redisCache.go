package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

var redisCache *cache.Cache

func initRedisCache() {
	_, redisCache = NewRedisCache(map[string]string{"default": redisDefault}, redisDB)
}

func NewRedisCache(hostPort map[string]string, db int) (*redis.Ring, *cache.Cache) {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: hostPort,
		DB:    db,
	})

	cacheDefault := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Hour),
		Marshal: func(v interface{}) ([]byte, error) {
			return json.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return json.Unmarshal(b, v)
		},
	})

	return ring, cacheDefault
}

func CacheSet(key string, obj interface{}) error {
	return redisCache.Set(&cache.Item{
		Key:   key,
		Value: obj,
		TTL:   time.Hour * 24 * 365,
	})
}

func CacheGet(ctx context.Context, key string, obj interface{}) error {
	return redisCache.Get(ctx, key, obj)
}
