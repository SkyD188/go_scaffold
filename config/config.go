package config

import "github.com/spf13/viper"

func (c *configStruct) InitConfig(confName, confType, confPath string, conf interface{}) {
	v := viper.New()

	v.SetConfigName(confName) // 配置文件的名字
	v.SetConfigType(confType) // 配置文件的类型
	v.AddConfigPath(confPath) // 配置文件的路径

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&conf); err != nil {
		panic(err)
	}

}

type configStruct struct {
}

func Conf() Iface {
	return &configStruct{}
}
