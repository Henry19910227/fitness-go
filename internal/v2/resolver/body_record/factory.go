package body_record

import (
	bodyService "github.com/Henry19910227/fitness-go/internal/v2/service/body_record"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	bodySvc := bodyService.NewService(db)
	return New(bodySvc)
}
