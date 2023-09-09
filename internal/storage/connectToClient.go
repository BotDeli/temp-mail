package storage

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

func ConnectToClient(nDB int, ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("redis_address") + ":" + os.Getenv("redis_port"),
		DB:   nDB,
	})
	pingClient(client, ctx)
	return client
}
