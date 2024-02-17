package main

import (
	"github.com/j3-n/tuner/api/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	a := app.New()
	a.Run()
	a.Shutdown()
}
