package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
	planOptional "github.com/Henry19910227/fitness-go/internal/v2/field/plan/optional"
	reviewOptional "github.com/Henry19910227/fitness-go/internal/v2/field/review/optional"
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

// APIGetProductCoursesInput /v2/product/courses [GET]
type APIGetProductCoursesInput struct {
	required.UserIDField
	Query APIGetProductCoursesQuery
}
type APIGetProductCoursesQuery struct {
	optional.NameField
	reviewOptional.ScoreField
	Level        *string `json:"level,omitempty" form:"level" binding:"omitempty,course_level" example:"3"`                          // 強度(1:初級/2:中級/3:中高級/4:高級)-複選
	Category     *string `json:"category,omitempty" form:"category" binding:"omitempty,course_category" example:"3"`                 // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)-複選
	Suit         *string `json:"suit,omitempty" form:"suit" binding:"omitempty,course_suit" example:"7"`                             // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)-複選
	Equipment    *string `json:"equipment,omitempty" form:"equipment" binding:"omitempty,course_equipment" example:"5"`              // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)-複選
	Place        *string `json:"place,omitempty" form:"place" binding:"omitempty,course_place" example:"1,2,3"`                      // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)-複選
	TrainTarget  *string `json:"train_target,omitempty" form:"train_target" binding:"omitempty,course_trainer_target" example:"1"`   // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)-複選
	BodyTarget   *string `json:"body_target,omitempty" form:"body_target" binding:"omitempty,course_body_target" example:"2"`        // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)-複選
	SaleType     *string `json:"sale_type,omitempty" form:"sale_type" binding:"omitempty,course_sale_type" example:"1"`              // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)-複選
	TrainerSex   *string `json:"trainer_sex,omitempty" form:"trainer_sex" binding:"omitempty,course_trainer_sex" example:"m"`        // 教練性別(m:男性/f:女性)-複選
	TrainerSkill *string `json:"trainer_skill,omitempty" form:"trainer_skill" binding:"omitempty,course_trainer_skill" example:"5"`  // 專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)
	OrderField 	 *string `json:"order_field" form:"order_field" binding:"omitempty,oneof=latest popular" example:"latest"` 			 // 排序類型(latest:最新/popular:熱門)-單選
	PagingInput
}
