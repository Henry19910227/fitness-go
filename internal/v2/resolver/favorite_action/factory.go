package favorite_action

import (

	"github.com/Henry19910227/fitness-go/internal/v2/service/favorite_action"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	favoriteActionService := favorite_action.NewService(db)
	return New(favoriteActionService)
}
