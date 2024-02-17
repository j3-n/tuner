package app

import (
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/j3-n/tuner/api/internal/endpoints"
	"github.com/j3-n/tuner/api/internal/quiz"
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

	a.Fiber.Get("/api/v1/questions", func(c *fiber.Ctx) error {
		eer := quiz.QuestionsSet[1]

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

	a.Fiber.Post("/user_answer", endpoints.PostUserAnswer)

	log.Fatal(a.Fiber.Listen(":4444"))
}

func (a *App) Shutdown() {
	log.Println("app shutting down")
}
