package main

import "github.com/joho/godotenv"

func initConfig() {
	// user .env file
	godotenv.Load()

	env, err := godotenv.Read()
	if err != nil {
		panic(err)
	}

	app.Config = env
}
