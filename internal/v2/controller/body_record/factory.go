package body_record

import (
	body "github.com/Henry19910227/fitness-go/internal/v2/resolver/body_record"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := body.NewResolver(db)
	return New(resolver)
}
