package http

import (
	"go_scaffold/app/service/controller/hello"

	"github.com/kataras/iris"
)

func initRoute(app *iris.Application) {
	app.Get("/hello", hello.Hello)
}
