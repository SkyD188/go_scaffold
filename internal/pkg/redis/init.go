package redis

import (
	"github.com/go-redis/redis"
	"log"
)

var (
	redisCache *redis.Client
)

func InitRedis(address, password string, poolSize int) {

	redisOptions := &redis.Options{
		ReadTimeout: -1,
		Addr:        address,
		Password:    password,
		PoolSize:    poolSize,
	}
	redisCache = redis.NewClient(redisOptions)

	pong, err := redisCache.Ping().Result()
	if err != nil {
		panic(err)
	}

	log.Println("Redis is Collection!!!", pong)
}
