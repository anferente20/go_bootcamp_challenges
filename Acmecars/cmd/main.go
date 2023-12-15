package main

import (
	"app/cmd/application"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// env
	godotenv.Load()

	// app
	// - config
	cfg := &application.ConfigDefaultInMemory{
		FileLoader: "./docs/db/vehicles_100.json",
		Addr:       os.Getenv("SERVER_ADDR"),
	}

	// - app
	app := application.NewDefaultInMemory(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
