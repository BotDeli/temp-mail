package connections

import (
	"context"
	"github.com/redis/go-redis/v9"
	"temp-mail/internal/storage"
)

type RedisConnections struct {
	Ctx    context.Context
	Client *redis.Client
}

func InitStorageConnections() StorageConnections {
	return &RedisConnections{
		Ctx:    context.Background(),
		Client: storage.ConnectToClient(2),
	}
}
