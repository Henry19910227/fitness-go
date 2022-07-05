package feedback

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/feedback_image"
)

type Output struct {
	Table
	Images []*feedback_image.Output
}

func (Output) TableName() string {
	return "feedbacks"
}

// APICreateFeedbackOutput /v2/feedback [POST]
type APICreateFeedbackOutput struct {
	base.Output
}
