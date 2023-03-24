package database

import (
	"fmt"
	"github.com/alvinashiatey/kiosk_client/internal/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDatabase() *gorm.DB {
	dbPath := fmt.Sprintf("data/%s.db", config.Config.DB.Name)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to database")
		return nil
	}
	DB = db
	log.Info().Msg("Connection opened to Database successfully")
	return DB
}
