package config

import (
	"github.com/spf13/viper"
)

// 传入配置文件类型、所在路径和需要反序列化结构体地址
// 初始化配置只是在启动会初始化。不用担心性能速度问题
func Load(confPath string, conf interface{}) {
	v := viper.New()

	v.SetConfigFile(confPath) // 配置文件的路径

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(conf); err != nil {
		panic(err)
	}
}
