package mails

import (
	"context"
	"github.com/redis/go-redis/v9"
	"temp-mail/internal/domain"
	"temp-mail/internal/storage"
)

type RedisMails struct {
	Ctx    context.Context
	Client *redis.Client
	Dom    *domain.DomainsGetter
}

func InitStorageMails() StorageMails {
	return &RedisMails{
		Ctx:    context.Background(),
		Client: storage.ConnectToClient(1),
		Dom:    domain.InitDomainsGetter(),
	}
}
