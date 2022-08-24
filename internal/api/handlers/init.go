package handlers

import "golang-api/internal/api/contracts"

var app *contracts.App

func Init(a *contracts.App) {
	app = a
}
