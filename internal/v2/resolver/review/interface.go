package review

import model "github.com/Henry19910227/fitness-go/internal/v2/model/review"

type Resolver interface {
	APIGetCMSReviews(input *model.APIGetCMSReviewsInput) (output model.APIGetCMSReviewsOutput)
}
