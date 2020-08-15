package deps

import (
	lib "go_scaffold/library/redis"

	"github.com/go-redis/redis"
)

var (
	CacheRedis *redis.Client
	QueueRedis *redis.Client
)

func InitRedis() {

	CacheRedis = lib.New(&AppConfig.Redis)

	QueueRedis = lib.New(&AppConfig.QueueRedis)

}
