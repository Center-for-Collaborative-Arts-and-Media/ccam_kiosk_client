package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

func fiberConnURL() string {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	return fmt.Sprintf("0.0.0.0:%s", PORT)
}

func StartServerWithGracefulShutdown(a *fiber.App) {
	idleConnectionsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Error().Err(err).Msg("Server is not shutting down")
		}

		close(idleConnectionsClosed)
	}()

	fiberConnURL := fiberConnURL()
	if err := a.Listen(fiberConnURL); err != nil {
		log.Error().Err(err).Msg("Server is not running")
	}
	<-idleConnectionsClosed
}

func StartServer(a *fiber.App) {
	fiberConnURL := fiberConnURL()
	if err := a.Listen(fiberConnURL); err != nil {
		log.Error().Err(err).Msg("Server is not running")
	}
}

func GenerateUUID() uuid.UUID {
	id, _ := uuid.NewUUID()
	return id
}
