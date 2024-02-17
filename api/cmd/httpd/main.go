package main

import (
	"fmt"

	"github.com/j3-n/tuner/api/internal/app"
)

func main() {
	fmt.Println("hello")
	a := app.New()
	a.Run()
	a.Shutdown()
}
