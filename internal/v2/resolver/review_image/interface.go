package review_image

import model "github.com/Henry19910227/fitness-go/internal/v2/model/review_image"

type Resolver interface {
	APIDeleteCMSReviewImage(input *model.APIDeleteCMSReviewImageInput) (output model.APIDeleteCMSReviewImageOutput)
}
