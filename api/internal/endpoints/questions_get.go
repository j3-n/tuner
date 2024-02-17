package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/j3-n/tuner/api/internal/quiz"
)

func GetQuestions(c *fiber.Ctx) error {
	eer := quiz.QuestionsSet[0]
	return c.JSON(eer)
}
