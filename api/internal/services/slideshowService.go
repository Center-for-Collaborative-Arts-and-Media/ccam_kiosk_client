package services

import (
	"github.com/alvinashiatey/kiosk_client/internal/repository"
)

type slideshowService struct {
	slidesRepo repository.SlidesRepository
}

type SlideshowService interface {
	GetSlides() interface{}
}

func NewSlideshowService(slidesRepo repository.SlidesRepository) SlideshowService {
	return &slideshowService{slidesRepo: slidesRepo}
}

func (s slideshowService) GetSlides() interface{} {
	return s.slidesRepo.FindAll()
}
