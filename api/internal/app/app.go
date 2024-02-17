package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Port string
	Log  bool
}

type App struct {
	Config Config
	Fiber  *fiber.App
}

func New(args ...Config) App {

	c := Config{
		Port: ":4444",
		Log:  false,
	}

	f := fiber.New()

	if len(args) > 0 {
		c = args[0]
	}

	return App{
		Config: c,
		Fiber:  f,
	}
}

func (a *App) Run() {

	a.Fiber.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(a.Fiber.Listen(":4444"))
}

func (a *App) Shutdown() {
	log.Println("app shutting down")
}
