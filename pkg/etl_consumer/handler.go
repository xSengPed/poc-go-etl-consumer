package etlhandler

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	App *fiber.App
}

func NewHandler() *Handler {
	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"hello": "world"})
	})
	return &Handler{
		App: app,
	}
}
