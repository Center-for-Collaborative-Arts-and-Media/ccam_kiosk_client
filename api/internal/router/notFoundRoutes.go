package router

import (
	"github.com/alvinashiatey/kiosk_client/internal/entities/web"
	"github.com/gofiber/fiber/v2"
)

func notFound(router fiber.Router) {
	router.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(web.Response{
				StatusCode: fiber.StatusNotFound,
				Error:      true,
				Data:       nil,
				Message:    "Sorry, endpoint not found",
			})
		},
	)
}
