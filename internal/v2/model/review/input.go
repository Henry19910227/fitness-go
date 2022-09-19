package review

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	IDOptional
	PreloadInput
}

type DeleteInput struct {
	IDOptional
}

type ListInput struct {
	courseOptional.NameField
	user.NicknameOptional
	ScoreOptional
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
	user.NicknameOptional
	ScoreOptional
	OrderByInput
	PagingInput
}

// APIUpdateCMSReviewInput /v2/cms/review/{review_id} [PATCH]
type APIUpdateCMSReviewInput struct {
	Uri  APIUpdateCMSReviewUri
	Body APIUpdateCMSReviewBody
}
type APIUpdateCMSReviewUri struct {
	IDRequired
}
type APIUpdateCMSReviewBody struct {
	ScoreOptional
	BodyOptional
}

// APIDeleteCMSReviewInput /v2/cms/review/{review_id} [DELETE]
type APIDeleteCMSReviewInput struct {
	Uri APIDeleteCMSReviewUri
}
type APIDeleteCMSReviewUri struct {
	IDRequired
}
