package config

import "github.com/spf13/viper"

var conf Config

func init() {
	v := viper.New()

	v.SetConfigName("conf")     // 配置文件的名字
	v.SetConfigType("toml")     // 配置文件的类型
	v.AddConfigPath("./config") // 配置文件的路径

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&conf); err != nil {
		panic(err)
	}

}

func GetConf() *Config {
	return &conf
}
