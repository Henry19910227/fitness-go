package workout

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/workout"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := workout.NewResolver(db)
	return New(resolver)
}
