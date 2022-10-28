package review

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	reviewOptional "github.com/Henry19910227/fitness-go/internal/v2/field/review/optional"
	reviewRequired "github.com/Henry19910227/fitness-go/internal/v2/field/review/required"
	userOptional "github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
	"github.com/Henry19910227/fitness-go/internal/v2/model/group"
	"github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/where"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type WhereInput = where.Input
type JoinInput = join.Input
type GroupInput = group.Input
type OrderByInput = orderBy.Input
type CustomOrderByInput = orderBy.CustomInput

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
	reviewOptional.CourseIDField
	reviewOptional.ScoreField
	reviewOptional.UserIDField
	JoinInput
	GroupInput
	WhereInput
	PagingInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
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

// APIGetStoreCourseReviewsInput /v2/store/course/{course_id}/reviews [GET]
type APIGetStoreCourseReviewsInput struct {
	userRequired.UserIDField
	Uri   APIGetStoreCourseReviewsUri
	Query APIGetStoreCourseReviewsQuery
}
type APIGetStoreCourseReviewsUri struct {
	reviewRequired.CourseIDField
}
type APIGetStoreCourseReviewsQuery struct {
	FilterType *int `json:"filter_type" form:"filter_type" binding:"omitempty,oneof=1 2" example:"1"` //篩選類型(1:全部/2:有照片)
	PagingInput
}

// APICreateStoreCourseReviewInput /v2/store/course/{course_id}/review [POST]
type APICreateStoreCourseReviewInput struct {
	userRequired.UserIDField
	Files []*file.Input
	Uri   APICreateStoreCourseReviewUri
	Form  APICreateStoreCourseReviewForm
}
type APICreateStoreCourseReviewUri struct {
	reviewRequired.CourseIDField
}
type APICreateStoreCourseReviewForm struct {
	reviewRequired.ScoreField
	reviewOptional.BodyField
}

// APIDeleteStoreCourseReviewInput /v2/store/course/{course_id}/review [DELETE]
type APIDeleteStoreCourseReviewInput struct {
	userRequired.UserIDField
	Uri APIDeleteStoreCourseReviewUri
}
type APIDeleteStoreCourseReviewUri struct {
	reviewRequired.IDField
}
