package rda

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/rda/api_update_rda"
	"gorm.io/gorm"
)

type Resolver interface {
	APIUpdateRDA(tx *gorm.DB, input *api_update_rda.Input) (output api_update_rda.Output)
}
