package app

import (
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/j3-n/tuner/api/internal/endpoints"
	"github.com/j3-n/tuner/api/internal/models"
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

	a.Fiber.Get("/api/v1/questions", func(c *fiber.Ctx) error {
		eer := models.QuestionsSet[1]

		return c.JSON(eer)
	})

	a.Fiber.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return os.Getenv("ENVIRONMENT") == "development"
		},
	}))

	a.Fiber.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	go endpoints.SocketListener()

	a.Fiber.Get("/ws/:id", websocket.New(endpoints.GetSocket))
	a.Fiber.Get("/api/v1/questions", endpoints.GetQuestions)
	a.Fiber.Post("/api/v1/user_answer", endpoints.PostUserAnswer)

	a.Fiber.Use("/create_lobby", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	a.Fiber.Get("/create", websocket.New(HandleCreationRequest))

	a.Fiber.Use("/play", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	a.Fiber.Get("/play/:lobby", websocket.New(HandleAddPlayerRequest))

	a.Fiber.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/login": auth.AuthURL(state),
		},
	}))

	a.Fiber.Get("/auth", Auth)

	a.Fiber.Use("/answer/response", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	a.Fiber.Get("/answer/response", websocket.New(HandlePlayerGuess))

	log.Fatal(a.Fiber.Listen(":4444"))
}

func (a *App) Shutdown() {
	log.Println("app shutting down")
}
