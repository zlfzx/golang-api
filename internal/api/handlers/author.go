package handlers

import (
	"golang-api/internal/api/entities"
	"golang-api/internal/api/models"
	"golang-api/pkg/ahttp"
	"strconv"

	"github.com/kataras/iris/v12"
)

func GetAuthors(c iris.Context) {
	authors, err := app.Service.Author.Get()
	if err != nil {
		httpError(c, ahttp.ErrInternalServer, err)
		return
	}

	httpSuccess(c, authors, "")
}

func GetAuthor(c iris.Context) {
	id, _ := strconv.Atoi(c.Params().Get("id"))

	author, err := app.Service.Author.GetById(id)
	if err != nil {
		httpError(c, ahttp.ErrInternalServer, err)
		return
	}

	httpSuccess(c, author, "")
}

func PostAuthor(c iris.Context) {
	params := models.Author{}
	err := c.ReadJSON(&params)
	if err != nil {
		httpError(c, ahttp.ErrBadRequest, err)
		return
	}

	// validate
	err = params.Validate()
	if err != nil {
		httpError(c, ahttp.ErrBadRequest, err)
		return
	}

	data := entities.Author{
		Name: params.Name,
	}

	// insert
	err = app.Service.Author.Insert(data)
	if err != nil {
		httpError(c, ahttp.ErrInternalServer, err)
		return
	}

	httpSuccess(c, nil, "")
}

func PutAuthor(c iris.Context) {
	id, _ := strconv.Atoi(c.Params().Get("id"))

	params := models.Author{}
	err := c.ReadJSON(&params)
	if err != nil {
		httpError(c, ahttp.ErrBadRequest, err)
		return
	}

	// validasi
	err = params.Validate()
	if err != nil {
		httpError(c, ahttp.ErrBadRequest, err)
		return
	}

	data := entities.Author{
		ID:   id,
		Name: params.Name,
	}

	// insert
	err = app.Service.Author.Update(data)
	if err != nil {
		httpError(c, ahttp.ErrInternalServer, err)
		return
	}

	httpSuccess(c, nil, "")
}

func DeleteAuthor(c iris.Context) {
	id, _ := strconv.Atoi(c.Params().Get("id"))

	err := app.Service.Author.Delete(id)
	if err != nil {
		httpError(c, ahttp.ErrInternalServer, err)
		return
	}

	httpSuccess(c, nil, "")
}
