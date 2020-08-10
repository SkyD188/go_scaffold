package config

import "github.com/spf13/viper"

var conf Config

func init() {
	v := viper.New()

	v.SetConfigFile("./config/conf.toml") // 配置文件的路径

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
