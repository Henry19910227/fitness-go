package review

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	reviewOptional "github.com/Henry19910227/fitness-go/internal/v2/field/review/optional"
	reviewRequired "github.com/Henry19910227/fitness-go/internal/v2/field/review/required"
	userOptional "github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	reviewOptional.IDField
	PreloadInput
}

type DeleteInput struct {
	reviewOptional.IDField
}

type ListInput struct {
	courseOptional.NameField
	userOptional.NicknameField
	reviewOptional.ScoreField
	PreloadInput
	OrderByInput
	PagingInput
}

// APIGetCMSReviewsInput /v2/cms/reviews [GET]
type APIGetCMSReviewsInput struct {
	Query APIGetCMSReviewsQuery
}
type APIGetCMSReviewsQuery struct {
	courseOptional.NameField
	userOptional.NicknameField
	reviewOptional.ScoreField
	OrderByInput
	PagingInput
}

// APIUpdateCMSReviewInput /v2/cms/review/{review_id} [PATCH]
type APIUpdateCMSReviewInput struct {
	Uri  APIUpdateCMSReviewUri
	Body APIUpdateCMSReviewBody
}
type APIUpdateCMSReviewUri struct {
	reviewRequired.IDField
}
type APIUpdateCMSReviewBody struct {
	reviewOptional.ScoreField
	reviewOptional.BodyField
}

// APIDeleteCMSReviewInput /v2/cms/review/{review_id} [DELETE]
type APIDeleteCMSReviewInput struct {
	Uri APIDeleteCMSReviewUri
}
type APIDeleteCMSReviewUri struct {
	reviewRequired.IDField
}
