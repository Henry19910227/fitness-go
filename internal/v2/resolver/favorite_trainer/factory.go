package favorite_trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/favorite_trainer"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	favoriteTrainerService := favorite_trainer.NewService(db)
	return New(favoriteTrainerService)
}
