package handlers

import (
	"golang-api/pkg/ahttp"

	"github.com/kataras/iris/v12"
)

func GetBooks(c iris.Context) {
	books, err := app.Service.Book.Get()
	if err != nil {
		httpError(c, ahttp.ErrInternalServer, err)
		return
	}

	httpSuccess(c, books, "")
}
