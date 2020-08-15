package conf

import (
	"go_scaffold/library/mongo"
	"go_scaffold/library/mysql"
	"go_scaffold/library/redis"
)

type AppConfigs struct {
	HttpListen string
	Mongo      mongo.Conf
	Mysql      mysql.Conf
	Redis      redis.Conf
	QueueRedis redis.Conf
}
