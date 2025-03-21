package review

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review"
	"gorm.io/gorm"
)

type Resolver interface {
	APIGetCMSReviews(input *model.APIGetCMSReviewsInput) (output model.APIGetCMSReviewsOutput)
	APIUpdateCMSReview(input *model.APIUpdateCMSReviewInput) (output model.APIUpdateCMSReviewOutput)
	APIDeleteCMSReview(input *model.APIDeleteCMSReviewInput) (output model.APIDeleteCMSReviewOutput)
	APIGetStoreCourseReviews(input *model.APIGetStoreCourseReviewsInput) (output model.APIGetStoreCourseReviewsOutput)
	APIGetStoreCourseReview(input *model.APIGetStoreCourseReviewInput) (output model.APIGetStoreCourseReviewOutput)
	APICreateStoreCourseReview(tx *gorm.DB, input *model.APICreateStoreCourseReviewInput) (output model.APICreateStoreCourseReviewOutput)
	APIDeleteStoreCourseReview(tx *gorm.DB, input *model.APIDeleteStoreCourseReviewInput) (output model.APIDeleteStoreCourseReviewOutput)
}
