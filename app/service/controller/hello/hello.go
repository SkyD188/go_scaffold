package hello

import "github.com/kataras/iris"

func Hello(ctx iris.Context) {
	_, _ = ctx.JSON("hello world")
}
