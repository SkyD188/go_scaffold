package main

import (
	"go_scaffold/config"
	route "go_scaffold/internal/app/http"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	route.InitHttpRoute(app)
	app.Run(iris.Addr(config.GetConf().Http.Listen))
}
