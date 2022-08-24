package main

import (
	"golang-api/internal/api/contracts"
	"golang-api/internal/api/services/author"
	"golang-api/internal/api/services/book"
)

func initServices() {
	app.Service = &contracts.Services{
		Author: author.Init(app),
		Book:   book.Init(app),
	}
}
