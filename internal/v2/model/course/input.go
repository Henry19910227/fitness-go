package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
	planOptional "github.com/Henry19910227/fitness-go/internal/v2/field/plan/optional"
	workoutOptional "github.com/Henry19910227/fitness-go/internal/v2/field/workout/optional"
	workoutSetOptional "github.com/Henry19910227/fitness-go/internal/v2/field/workout_set/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
	"github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"mime/multipart"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type WhereInput = where.Input
type JoinInput = join.Input
type OrderByInput = orderBy.Input
type CustomOrderByInput = orderBy.CustomInput

type GenerateInput struct {
	DataAmount int
	UserID     []*base.GenerateSetting
}

type FindInput struct {
	optional.IDField
	planOptional.PlanIDField
	workoutOptional.WorkoutIDField
	workoutSetOptional.WorkoutSetIDField
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
	optional.ScheduleTypeField
	JoinInput
	WhereInput
	PagingInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
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

// APIUpdateUserCourseInput /v2/user/course/{course_id} [PATCH]
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

// APIGetUserCourseStructureInput /v2/user/course/{course_id}/structure [GET]
type APIGetUserCourseStructureInput struct {
	required.UserIDField
	Uri APIGetUserCourseStructureUri
}
type APIGetUserCourseStructureUri struct {
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

// APIGetTrainerCourseInput /v2/trainer/course/{course_id} [GET]
type APIGetTrainerCourseInput struct {
	required.UserIDField
	Uri APIGetTrainerCourseUri
}
type APIGetTrainerCourseUri struct {
	required.IDField
}

// APIUpdateTrainerCourseInput /v2/trainer/course/{course_id} [PATCH]
type APIUpdateTrainerCourseInput struct {
	required.UserIDField
	Cover *file.Input
	Uri   APIUpdateTrainerCourseUri
	Form  APIUpdateTrainerCourseForm
}
type APIUpdateTrainerCourseUri struct {
	required.IDField
}
type APIUpdateTrainerCourseForm struct {
	optional.SaleTypeField
	optional.SaleIDField
	optional.CategoryField
	optional.NameField
	optional.IntroField
	optional.FoodField
	optional.LevelField
	optional.SuitField
	optional.EquipmentField
	optional.PlaceField
	optional.TrainTargetField
	optional.BodyTargetField
	optional.NoticeField
}

// APIDeleteTrainerCourseInput /v2/trainer/course/{course_id} [DELETE]
type APIDeleteTrainerCourseInput struct {
	required.UserIDField
	Uri APIDeleteTrainerCourseUri
}
type APIDeleteTrainerCourseUri struct {
	required.IDField
}

// APISubmitTrainerCourseInput /v2/trainer/course/{course_id}/submit [POST]
type APISubmitTrainerCourseInput struct {
	required.UserIDField
	Uri APISubmitTrainerCourseUri
}
type APISubmitTrainerCourseUri struct {
	required.IDField
}

// APIGetProductCourseInput /v2/product/course/{course_id} [GET]
type APIGetProductCourseInput struct {
	required.UserIDField
	Uri APIGetProductCourseUri
}
type APIGetProductCourseUri struct {
	required.IDField
}

// APIGetProductCourseStructureInput /v2/product/course/{course_id}/structure [GET]
type APIGetProductCourseStructureInput struct {
	required.UserIDField
	Uri APIGetProductCourseStructureUri
}
type APIGetProductCourseStructureUri struct {
	required.IDField
}
