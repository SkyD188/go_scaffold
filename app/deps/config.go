package deps

import (
	"go_scaffold/app/model/conf"
	"go_scaffold/library/config"
	"os"
	"strings"
)

var (
	AppConfig conf.AppConfigs
)

const (
	projectName = "go_scaffold"
	configPath  = "app" + string(os.PathSeparator) + "config"
	configName  = "config.toml"
)

// 找到配置文件所在目录
func getConfPath() string {
	pwd, _ := os.Getwd()
	index := strings.LastIndex(pwd, projectName)
	newPwd := pwd[:index+len(projectName)]

	ps := string(os.PathSeparator)
	return newPwd + ps + configPath + ps + configName
}

func LoadConfig() {
	config.Load(getConfPath(), &AppConfig)
}
