package server

import (
	"github.com/alvinashiatey/kiosk_client/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func NewServer(app *fiber.App) *Server {
	return &Server{app: app}
}

func (s *Server) Start(devMode bool) {
	if devMode {
		utils.StartServer(s.app)
	} else {
		utils.StartServerWithGracefulShutdown(s.app)
	}
}
