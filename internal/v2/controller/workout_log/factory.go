package workout_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/workout_log"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := workout_log.NewResolver(db)
	return New(resolver)
}
