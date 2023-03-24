package database

import (
	"github.com/alvinashiatey/kiosk_client/internal/entities/model"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(model.Models...); err != nil {
		log.Error().Err(err).Msg("failed to connect to database")
	}
}
