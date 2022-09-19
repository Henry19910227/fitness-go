package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/entity/course"
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
	course.IDOptional
	PlanID       *int64 `json:"plan_id,omitempty"`        // 計畫 id
	WorkoutID    *int64 `json:"workout_id,omitempty"`     // 訓練 id
	WorkoutSetID *int64 `json:"workout_set_id,omitempty"` // 訓練組 id
	PreloadInput
}

type DeleteInput struct {
	course.IDRequired
}

type ListInput struct {
	course.IDOptional
	course.UserIDOptional
	course.NameOptional
	course.CourseStatusOptional
	course.SaleTypeOptional
	SaleTypes           []int // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表/4:個人課表)
	IgnoredCourseStatus []int // 課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)
	IDs                 []int64
	PagingInput
	PreloadInput
	OrderByInput
}

type FavoriteListInput struct {
	course.UserIDOptional
	PagingInput
	PreloadInput
	OrderByInput
}

type ProgressListInput struct {
	course.UserIDRequired
	PagingInput
	PreloadInput
	OrderByInput
}

type ChargeListInput struct {
	course.UserIDRequired
	PagingInput
	PreloadInput
	OrderByInput
}

// APIGetFavoriteCoursesInput /v2/favorite/courses [GET]
type APIGetFavoriteCoursesInput struct {
	course.UserIDRequired
	Form APIGetFavoriteCoursesForm
}
type APIGetFavoriteCoursesForm struct {
	PagingInput
}

// APIGetCMSCoursesInput /v2/cms/courses [GET]
type APIGetCMSCoursesInput struct {
	course.IDOptional
	course.NameOptional
	course.CourseStatusOptional
	course.SaleTypeOptional
	PagingInput
	OrderByInput
}

// APIGetCMSCourseInput /v2/cms/course/{course_id} [GET]
type APIGetCMSCourseInput struct {
	course.IDRequired
}

// APIUpdateCMSCoursesStatusInput /v2/cms/courses/course_status [PATCH]
type APIUpdateCMSCoursesStatusInput struct {
	IDs []int64 `json:"ids" binding:"required"` // 課表 id
	course.CourseStatusRequired
}

// APIUpdateCMSCourseCoverInput /v2/cms/course/{course_id}/cover [PATCH]
type APIUpdateCMSCourseCoverInput struct {
	course.IDRequired
	CoverNamed string
	File       multipart.File
}

// APIGetUserCoursesInput /v2/user/courses [GET]
type APIGetUserCoursesInput struct {
	course.UserIDRequired
	Query APIGetUserCoursesQuery
}
type APIGetUserCoursesQuery struct {
	Type int `form:"type" binding:"required,oneof=1 2 3" example:"1"` // 搜尋類別(1:進行中課表/2:付費課表/3:個人課表)
	PagingInput
}

// APICreateUserCourseInput /v2/user/course [POST]
type APICreateUserCourseInput struct {
	course.UserIDRequired
	Body APICreateUserCourseBody
}
type APICreateUserCourseBody struct {
	course.NameRequired
	course.ScheduleTypeRequired
}

// APIDeleteUserCourseInput /v2/user/course/{course_id} [DELETE]
type APIDeleteUserCourseInput struct {
	course.UserIDRequired
	Uri APIDeleteUserCourseUri
}
type APIDeleteUserCourseUri struct {
	course.IDRequired
}

// APIUpdateUserCourseInput /v2/user/course/{course_id} [UPDATE]
type APIUpdateUserCourseInput struct {
	course.UserIDRequired
	Uri  APIUpdateUserCourseUri
	Body APIUpdateUserCourseBody
}
type APIUpdateUserCourseUri struct {
	course.IDRequired
}
type APIUpdateUserCourseBody struct {
	course.NameOptional
}

// APIGetUserCourseInput /v2/user/course/{course_id} [GET]
type APIGetUserCourseInput struct {
	course.UserIDRequired
	Uri APIGetUserCourseUri
}
type APIGetUserCourseUri struct {
	course.IDRequired
}

// APIGetTrainerCoursesInput /v2/trainer/courses [GET]
type APIGetTrainerCoursesInput struct {
	course.UserIDRequired
	Query APIGetTrainerCoursesQuery
}
type APIGetTrainerCoursesQuery struct {
	course.CourseStatusOptional
	PagingInput
}

// APICreateTrainerCourseInput /v2/trainer/course [POST]
type APICreateTrainerCourseInput struct {
	course.UserIDRequired
	Body APICreateTrainerCourseBody
}
type APICreateTrainerCourseBody struct {
	course.NameRequired
	course.CategoryRequired
	course.LevelRequired
	course.ScheduleTypeRequired
}
