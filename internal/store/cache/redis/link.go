package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Neko2h/shortener/internal/entity"
	"github.com/go-redis/redis/v8"
)

type LinkCache struct {
	cache *redis.Client
}

// NewFileMetaRepo ...
func NewLinkCache(cache *redis.Client) *LinkCache {
	return &LinkCache{cache: cache}
}

func (c LinkCache) Get(ctx context.Context, hash string) (*entity.LinkDB, error) {
	result, err := c.cache.Get(ctx, hash).Result()

	if err == redis.Nil {
		fmt.Println("link was not found in cache")
		return nil, nil
	}

	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	data := &entity.LinkDB{}
	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (c LinkCache) Set(ctx context.Context, link *entity.LinkDB) error {
	hashKey := link.Hash

	b, err := json.Marshal(link)
	if err != nil {
		return err
	}

	err = c.cache.Set(ctx, hashKey, b, 0).Err()

	if err != nil {
		return err
	}
	return nil
}
