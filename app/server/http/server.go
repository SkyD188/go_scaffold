package http

import (
	"github.com/kataras/iris"
)

func RunHTTPServer(address string) {
	app := iris.New()

	// 中间件处理

	initRoute(app)

	_ = app.Run(iris.Addr(address))

}
