package rda

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/diet"
	"github.com/Henry19910227/fitness-go/internal/v2/service/rda"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	rdaService := rda.NewService(db)
	dietService := diet.NewService(db)
	return New(rdaService, dietService)
}
