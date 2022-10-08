package main

import (
	"golang-api/internal/api/constants"
	"golang-api/internal/api/contracts"
	"golang-api/internal/api/handlers"
	"golang-api/internal/api/routers"
	"golang-api/pkg/alog"
	"net/http"
	"os"
	"time"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

var app *contracts.App

func main() {

	os.Setenv("TZ", "Asia/Jakarta")

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Accept", "content-type", "X-Requested-With", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Screen"},
		AllowCredentials: true,
	})

	irisApp := iris.New()
	irisApp.Use(crs)
	irisApp.AllowMethods(iris.MethodOptions)

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
		Addr:         ":" + os.Getenv(constants.AppPort),
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
	}

	irisApp.Run(iris.Server(srv))
}
