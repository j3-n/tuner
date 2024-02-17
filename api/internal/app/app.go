package app

import (
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/j3-n/tuner/api/internal/endpoints"
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
	auth = spotifyauth.New(spotifyauth.WithRedirectURL("http://localhost:4444/auth"), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopeUserTopRead))
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

	a.Fiber.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return os.Getenv("ENVIRONMENT") == "development"
		},
	}))

	a.Fiber.Get("/api/v1/questions", endpoints.GetQuestions)
	a.Fiber.Post("/api/v1/user_answer", endpoints.PostUserAnswer)
	a.Fiber.Get("/create", websocket.New(HandleCreationRequest))
	a.Fiber.Get("/play/:lobby", websocket.New(HandleAddPlayerRequest))

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
