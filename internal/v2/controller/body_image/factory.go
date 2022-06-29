package body_image

import (
	bodyImage "github.com/Henry19910227/fitness-go/internal/v2/resolver/body_image"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := bodyImage.NewResolver(db)
	return New(resolver)
}
