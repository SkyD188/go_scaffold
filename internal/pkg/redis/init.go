package redis

import (
	"go_scaffold/config"
	"log"

	"github.com/go-redis/redis"
)

var (
	redisCache *redis.Client
)

func InitRedis() {

	redisOptions := &redis.Options{
		ReadTimeout: -1,
	}
	redisOptions.Addr = config.GetConf().Redis.Address
	if config.GetConf().Redis.Password != "" {
		redisOptions.Password = config.GetConf().Redis.Password
	}
	redisOptions.DB = 0
	if config.GetConf().Redis.MaxPoolSize > 0 {
		redisOptions.PoolSize = config.GetConf().Redis.MaxPoolSize
	}
	redisCache = redis.NewClient(redisOptions)

	pong, err := redisCache.Ping().Result()
	if err != nil {
		panic(err)
	}

	log.Println("Redis is Collection!!!", pong)
}
