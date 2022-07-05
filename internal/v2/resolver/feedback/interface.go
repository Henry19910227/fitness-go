package feedback

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/feedback"
	"gorm.io/gorm"
)

type Resolver interface {
	APICreateFeedback(tx *gorm.DB, input *model.APICreateFeedbackInput) (output base.Output)
}
