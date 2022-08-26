package course

import (
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
	IDOptional
	PlanID    *int64 `json:"plan_id,omitempty"`    // 計畫 id
	WorkoutID *int64 `json:"workout_id,omitempty"` // 訓練 id
	PreloadInput
}

type ListInput struct {
	IDOptional
	NameOptional
	CourseStatusOptional
	SaleTypeOptional
	IgnoredCourseStatus []int // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	IDs                 []int64
	PagingInput
	PreloadInput
	OrderByInput
}

type FavoriteListInput struct {
	UserIDOptional
	PagingInput
	PreloadInput
	OrderByInput
}

// APIGetFavoriteCoursesInput /v2/favorite/courses [GET]
type APIGetFavoriteCoursesInput struct {
	UserIDRequired
	Form APIGetFavoriteCoursesForm
}
type APIGetFavoriteCoursesForm struct {
	PagingInput
}

// APIGetCMSCoursesInput /v2/cms/courses [GET]
type APIGetCMSCoursesInput struct {
	IDOptional
	NameOptional
	CourseStatusOptional
	SaleTypeOptional
	PagingInput
	OrderByInput
}

// APIGetCMSCourseInput /v2/cms/course/{course_id} [GET]
type APIGetCMSCourseInput struct {
	IDRequired
}

// APIUpdateCMSCoursesStatusInput /v2/cms/courses/course_status [PATCH]
type APIUpdateCMSCoursesStatusInput struct {
	IDs []int64 `json:"ids" binding:"required"` // 課表 id
	CourseStatusRequired
}

// APIUpdateCMSCourseCoverInput /v2/cms/course/{course_id}/cover [PATCH]
type APIUpdateCMSCourseCoverInput struct {
	IDRequired
	CoverNamed string
	File       multipart.File
}

// APICreateUserCourseInput /v2/user/course [POST]
type APICreateUserCourseInput struct {
	UserIDRequired
	Body APICreateUserCourseBody
}
type APICreateUserCourseBody struct {
	NameRequired
	ScheduleTypeRequired
}
