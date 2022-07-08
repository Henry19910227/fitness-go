package review_image

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

type Output struct {
	Table
}

func (Output) TableName() string {
	return "review_images"
}

// APIDeleteCMSReviewImageOutput /v2/cms/review_image/{review_image_id} [DELETE]
type APIDeleteCMSReviewImageOutput struct {
	base.Output
}
