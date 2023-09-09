package storage

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

func pingClient(client *redis.Client, ctx context.Context) {
	if cmd := client.Ping(ctx); cmd.Err() != nil {
		log.Fatalf("ping client error: %v", cmd.Err())
	}
}
