package rda

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/rda"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := rda.NewResolver(db)
	return New(resolver)
}
