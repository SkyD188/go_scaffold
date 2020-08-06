package config

type Iface interface {
	InitConfig(confName, confType, confPath string, conf interface{})
}
