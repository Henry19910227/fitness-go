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

// APIGetFavoriteCoursesInput /v2/favorite/courses [GET]
type APIGetFavoriteCoursesInput struct {
	required.UserIDField
	Form APIGetFavoriteCoursesForm
}
type APIGetFavoriteCoursesForm struct {
	PagingInput
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
