package http

import (
	"github.com/kataras/iris/v12"
	"report-api/api"
)

func LoadRouter(app *iris.Application) {
	test(app)
}

func test(app *iris.Application) {
	app.Get("/", api.HelloWorld)
}
