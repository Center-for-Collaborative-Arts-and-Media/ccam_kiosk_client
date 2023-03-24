package repository

import (
	"github.com/alvinashiatey/kiosk_client/internal/entities/model"
	"github.com/alvinashiatey/kiosk_client/internal/entities/web"
	"gorm.io/gorm"
)

type SlidesRepository interface {
	FindAll() web.RepositoryResult
}

type slidesRepository struct {
	data *gorm.DB
}

func NewSlidesRepository(data *gorm.DB) SlidesRepository {
	return &slidesRepository{data: data}
}

func (r slidesRepository) FindAll() web.RepositoryResult {
	var slides []model.Slideshow
	r.data.Find(&slides)
	return web.RepositoryResult{
		Result: slides,
		Error:  nil,
	}
}
