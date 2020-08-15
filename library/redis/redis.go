package redis

import (
	"log"

	"github.com/go-redis/redis"
)

type Conf struct {
	Address     string
	MaxPoolSize int
	Password    string
}

func New(conf *Conf) *redis.Client {

	redisOptions := &redis.Options{
		Addr:     conf.Address,
		Password: conf.Password,
		PoolSize: conf.MaxPoolSize,
	}

	redisCache := redis.NewClient(redisOptions)

	pong, err := redisCache.Ping().Result()
	if err != nil {
		panic(err)
	}

	log.Println("Redis is Collection!!!", pong)
	return redisCache
}
