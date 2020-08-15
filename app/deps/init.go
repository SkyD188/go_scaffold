package deps

// 初始化整个项目用到的连接库
func InitProject() {
	// 加载项目配置
	LoadConfig()
	// redis
	InitRedis()
	// mysql
	InitMysql()
	// mongo
	InitMongo()
}
