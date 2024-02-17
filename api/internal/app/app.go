package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

var (
	auth  *spotifyauth.Authenticator
	state = "joemama"
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
	auth = spotifyauth.New(spotifyauth.WithRedirectURL("http://localhost:4444/auth"), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))

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

	a.Fiber.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/login": auth.AuthURL(state),
		},
	}))

	a.Fiber.Get("/auth", Auth)

	log.Fatal(a.Fiber.Listen(":4444"))
}

func (a *App) Shutdown() {
	log.Println("app shutting down")
}
