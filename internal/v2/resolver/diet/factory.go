package diet

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/diet"
	"github.com/Henry19910227/fitness-go/internal/v2/service/rda"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	dietService := diet.NewService(db)
	rdaService := rda.NewService(db)
	return New(dietService, rdaService)
}
