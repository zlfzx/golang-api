package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "Hello, World!",
		})
	})

	app.Run(iris.Addr(":8080"))
}
