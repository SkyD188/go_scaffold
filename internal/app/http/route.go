package http

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func InitHttpRoute(app *iris.Application) {
	app.Get("/hello", hello)
}

func hello(ctx context.Context) {
	ctx.JSON("hello world")
}
