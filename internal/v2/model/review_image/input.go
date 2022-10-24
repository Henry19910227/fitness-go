package review_image

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/review_image/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/review_image/required"
)

type FindInput struct {
	optional.IDField
}

type DeleteInput struct {
	optional.IDField
}

// APIDeleteCMSReviewImageInput /v2/cms/review_image/{review_image_id} [DELETE]
type APIDeleteCMSReviewImageInput struct {
	Uri APIDeleteCMSReviewImageUri
}
type APIDeleteCMSReviewImageUri struct {
	required.IDField
}
