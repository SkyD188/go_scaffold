package main

import (
	"go_scaffold/app/deps"
	"go_scaffold/app/server/http"
)

func main() {

	deps.InitProject()
	//deps.LoadConfig()
	http.RunHTTPServer(deps.AppConfig.HttpListen)
}
