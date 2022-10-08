package main

import (
	"golang-api/internal/api/contracts"
	"golang-api/internal/api/handlers"
	"golang-api/internal/api/routers"
	"golang-api/pkg/alog"
	"net/http"
	"os"
	"time"

	"github.com/kataras/iris/v12"
)

var app *contracts.App

func main() {

	os.Setenv("TZ", "Asia/Jakarta")

	// cors := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 	AllowCredentials: true,
	// })

	irisApp := iris.Default()

	app = &contracts.App{
		Iris: irisApp,
	}

	initConfig()
	alog.Init()

	initDatasources()

	initServices()

	handlers.Init(app)
	routers.Init(app)

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
	}

	irisApp.Run(iris.Server(srv))
}
