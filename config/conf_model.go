package config

type Config struct {
	Mongo mongoConf
	Http  httpConf
	Mysql mysqlConf
	Redis redisConf
}

// Mongo
type mongoConf struct {
	Address     string
	MaxPoolSize uint64
}

// Http
type httpConf struct {
	Listen string
}

// Mysql
type mysqlConf struct {
	Address     string
	MaxOpenConn int
	MaxIdleConn int
}

// Redis
type redisConf struct {
	Address     string
	MaxPoolSize int
	Password    string
}
