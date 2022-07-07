package review

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/course"
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
}

type ListInput struct {
	course.NameOptional
	user.NicknameOptional
	ScoreOptional
	PreloadInput
	OrderByInput
	PagingInput
}

// APIGetCMSReviewsInput /v2/cms/reviews [GET]
type APIGetCMSReviewsInput struct {
	Form APIGetCMSReviewsForm
}
type APIGetCMSReviewsForm struct {
	course.NameOptional
	user.NicknameOptional
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
