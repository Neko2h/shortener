package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	*redis.Client
}

func NewRedisClient(redisHost string) (*Cache, error) {
	//"redis://<user>:<pass>@localhost:6379/<db>"

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost, //"localhost:6379", // host:port of the redis server
		Password: "",        // no password set
		DB:       0,         // use default DB
	})

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	return &Cache{client}, err
}
