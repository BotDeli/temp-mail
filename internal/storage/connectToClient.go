package storage

import "github.com/redis/go-redis/v9"

func ConnectToClient(nDB int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   nDB,
	})
}
