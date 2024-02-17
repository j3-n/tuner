package main

import (
	"github.com/j3-n/tuner/api/internal/app"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	a := app.New()
	a.Run()
	a.Shutdown()
}
