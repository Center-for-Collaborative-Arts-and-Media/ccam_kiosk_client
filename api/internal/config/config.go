package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

type AppConfig struct {
	App struct {
		Name       string
		Port       string
		JWTSecret  string
		JWTRefresh string
		DevMode    bool
	}

	DB struct {
		Name string
	}
}

var Config AppConfig

func InitConfig(DevMode bool) *AppConfig {
	if DevMode {
		Config.App.DevMode = true
		if err := godotenv.Load(".env"); err != nil {
			log.Error().Err(err).Msg("Error loading .env file")
		}
	}

	Config.App.Name = os.Getenv("APP_NAME")
	Config.App.Port = os.Getenv("PORT")
	Config.App.JWTSecret = os.Getenv("JWT_SECRET_KEY")
	Config.App.JWTRefresh = os.Getenv("JWT_REFRESH_KEY")
	Config.DB.Name = os.Getenv("DB_NAME")
	return &Config
}
