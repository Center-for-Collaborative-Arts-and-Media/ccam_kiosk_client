package main

import (
	"flag"
	"github.com/alvinashiatey/kiosk_client/database"
	"github.com/alvinashiatey/kiosk_client/internal/config"
	"github.com/alvinashiatey/kiosk_client/internal/router"
	"github.com/alvinashiatey/kiosk_client/server"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	DevMode = flag.Bool("dev", false, "Development mode")
)

func init() {
	flag.Parse()
	config.InitConfig(*DevMode)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if *DevMode {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	database.InitDatabase()
	database.AutoMigrate(database.DB)
}

func main() {
	app := fiber.New()
	router.RegisterRoutes(app, database.DB)
	newServer := server.NewServer(app)
	newServer.Start(*DevMode)
}
