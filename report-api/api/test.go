package api

import "github.com/kataras/iris/v12/context"

func HelloWorld(context *context.Context) {
	context.HTML("<h1>hello world</h1>")
}
