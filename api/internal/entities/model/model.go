package model

import "github.com/google/uuid"

var Models = []interface{}{&Slideshow{}, &Media{}, &Video{}, &Image{}}

type Media struct {
	Base
	MediaID uuid.UUID `json:"media_id"`
	Videos  []Video   `json:"videos" gorm:"many2many:media_videos;joinForeignKey:VideoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Images  []Image   `json:"images" gorm:"many2many:media_images;joinForeignKey:ImageID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Embeds  []Embed   `json:"embeds" gorm:"many2many:media_embeds;joinForeignKey:EmbedID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Embed struct {
	Base
	EmbedID   uuid.UUID `json:"embed_id"`
	URL       string    `json:"url"`
	Caption   string    `json:"caption"`
	Order     uint16    `json:"order"`
	EmbedHtml string    `json:"embed_html"`
}

type Video struct {
	Base
	Type    string    `json:"type"`
	VideoID uuid.UUID `json:"video_id"`
	URL     string    `json:"url"`
	Caption string    `json:"caption"`
	Order   uint16    `json:"order"`
}

type Image struct {
	Base
	Type     string    `json:"type"`
	ImageID  uuid.UUID `json:"image_id"`
	URL      string    `json:"url"`
	Duration int       `json:"duration"`
	Caption  string    `json:"caption"`
	Order    uint16    `json:"order"`
}

type Slideshow struct {
	Base
	SlideID     uuid.UUID `json:"slide_id"`
	Title       string    `json:"title"  gorm:"varchar(255)"`
	TextOverlay string    `json:"text_overlay"`
	Media       Media     `json:"media" gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Archive     bool      `json:"archive" gorm:"default:false"`
}
