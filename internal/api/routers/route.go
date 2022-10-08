package routers

import (
	"golang-api/internal/api/contracts"
	"golang-api/internal/api/handlers"

	"github.com/kataras/iris/v12"
)

func Init(app *contracts.App) {

	route := app.Iris

	route.Get("/", func(c iris.Context) {
		c.JSON(iris.Map{
			"status": true,
		})
	})

	v1 := route.Party("v1")

	author := v1.Party("/author")
	{
		author.Get("/", handlers.GetAuthors)
		author.Get("/{id:int}", handlers.GetAuthor)
		author.Post("/", handlers.PostAuthor)
		author.Put("/:id", handlers.PutAuthor)
		author.Delete("/:id", handlers.DeleteAuthor)
	}

	book := v1.Party("/book")
	{
		book.Get("/", handlers.GetBooks)
	}
}
