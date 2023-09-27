package config

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	opt, err := redis.ParseURL("redis://default:rootpw@localhost:6379/0")
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(opt)
	return client
}
