package router

import (
	"github.com/alvinashiatey/kiosk_client/internal/controller"
	"github.com/alvinashiatey/kiosk_client/internal/repository"
	"github.com/alvinashiatey/kiosk_client/internal/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func slideshowRoutes(router fiber.Router, db *gorm.DB) {
	slideshowRepo := repository.NewSlidesRepository(db)
	slideshowService := services.NewSlideshowService(slideshowRepo)
	slideshowController := controller.NewSlideshowController(slideshowService)
	playlist := router.Group("/slideshow")
	playlist.Get("/", slideshowController.GetSlides)
}
