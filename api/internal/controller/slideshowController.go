package controller

import (
	"github.com/alvinashiatey/kiosk_client/internal/services"
	"github.com/gofiber/fiber/v2"
)

type SlideshowController interface {
	GetSlides(c *fiber.Ctx) error
}

type slideshowController struct {
	slideshowService services.SlideshowService
}

func NewSlideshowController(slideshowService services.SlideshowService) SlideshowController {
	return &slideshowController{slideshowService: slideshowService}
}

func (c slideshowController) GetSlides(ctx *fiber.Ctx) error {
	return ctx.JSON(c.slideshowService.GetSlides())
}
