package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"mime/multipart"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type GenerateInput struct {
	DataAmount int
	UserID     []*base.GenerateSetting
}

type FindInput struct {
	optional.IDField
	PlanID       *int64 `json:"plan_id,omitempty"`        // 計畫 id
	WorkoutID    *int64 `json:"workout_id,omitempty"`     // 訓練 id
	WorkoutSetID *int64 `json:"workout_set_id,omitempty"` // 訓練組 id
	PreloadInput
}

type DeleteInput struct {
	required.IDField
}

type ListInput struct {
	optional.IDField
	optional.UserIDField
	optional.NameField
	optional.CourseStatusField
	optional.SaleTypeField
	SaleTypes           []int // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表/4:個人課表)
	IgnoredCourseStatus []int // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	IDs                 []int64
	PagingInput
	PreloadInput
	OrderByInput
}

type FavoriteListInput struct {
	optional.UserIDField
	PagingInput
	PreloadInput
	OrderByInput
}

type ProgressListInput struct {
	required.UserIDField
	PagingInput
	PreloadInput
	OrderByInput
}

type ChargeListInput struct {
	required.UserIDField
	PagingInput
	PreloadInput
	OrderByInput
}

// APIGetFavoriteCoursesInput /v2/favorite/courses [GET]
type APIGetFavoriteCoursesInput struct {
	required.UserIDField
	Form APIGetFavoriteCoursesForm
}
type APIGetFavoriteCoursesForm struct {
	PagingInput
}

// APIGetCMSCoursesInput /v2/cms/courses [GET]
type APIGetCMSCoursesInput struct {
	optional.IDField
	optional.NameField
	optional.CourseStatusField
	optional.SaleTypeField
	PagingInput
	OrderByInput
}

// APIGetCMSCourseInput /v2/cms/course/{course_id} [GET]
type APIGetCMSCourseInput struct {
	required.IDField
}

// APIUpdateCMSCoursesStatusInput /v2/cms/courses/course_status [PATCH]
type APIUpdateCMSCoursesStatusInput struct {
	IDs []int64 `json:"ids" binding:"required"` // 課表 id
	required.CourseStatusField
}

// APIUpdateCMSCourseCoverInput /v2/cms/course/{course_id}/cover [PATCH]
type APIUpdateCMSCourseCoverInput struct {
	required.IDField
	CoverNamed string
	File       multipart.File
}

// APIGetUserCoursesInput /v2/user/courses [GET]
type APIGetUserCoursesInput struct {
	required.UserIDField
	Query APIGetUserCoursesQuery
}
type APIGetUserCoursesQuery struct {
	Type int `form:"type" binding:"required,oneof=1 2 3" example:"1"` // 搜尋類別(1:進行中課表/2:付費課表/3:個人課表)
	PagingInput
}

// APICreateUserCourseInput /v2/user/course [POST]
type APICreateUserCourseInput struct {
	required.UserIDField
	Body APICreateUserCourseBody
}
type APICreateUserCourseBody struct {
	required.NameField
	required.ScheduleTypeField
}

// APIDeleteUserCourseInput /v2/user/course/{course_id} [DELETE]
type APIDeleteUserCourseInput struct {
	required.UserIDField
	Uri APIDeleteUserCourseUri
}
type APIDeleteUserCourseUri struct {
	required.IDField
}

// APIUpdateUserCourseInput /v2/user/course/{course_id} [UPDATE]
type APIUpdateUserCourseInput struct {
	required.UserIDField
	Uri  APIUpdateUserCourseUri
	Body APIUpdateUserCourseBody
}
type APIUpdateUserCourseUri struct {
	required.IDField
}
type APIUpdateUserCourseBody struct {
	optional.NameField
}

// APIGetUserCourseInput /v2/user/course/{course_id} [GET]
type APIGetUserCourseInput struct {
	required.UserIDField
	Uri APIGetUserCourseUri
}
type APIGetUserCourseUri struct {
	required.IDField
}

// APIGetTrainerCoursesInput /v2/trainer/courses [GET]
type APIGetTrainerCoursesInput struct {
	required.UserIDField
	Query APIGetTrainerCoursesQuery
}
type APIGetTrainerCoursesQuery struct {
	optional.CourseStatusField
	PagingInput
}

// APICreateTrainerCourseInput /v2/trainer/course [POST]
type APICreateTrainerCourseInput struct {
	required.UserIDField
	Body APICreateTrainerCourseBody
}
type APICreateTrainerCourseBody struct {
	required.NameField
	required.CategoryField
	required.LevelField
	required.ScheduleTypeField
}
