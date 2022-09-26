package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
	planOptional "github.com/Henry19910227/fitness-go/internal/v2/field/plan/optional"
	workoutOptional "github.com/Henry19910227/fitness-go/internal/v2/field/workout/optional"
	workoutSetOptional "github.com/Henry19910227/fitness-go/internal/v2/field/workout_set/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
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
