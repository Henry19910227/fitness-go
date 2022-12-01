package api_update_user_course

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/user/course/{course_id} [PATCH]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}

type Data struct {
	courseOptional.IDField
	courseOptional.SaleTypeField
	courseOptional.SaleIDField
	courseOptional.CourseStatusField
	courseOptional.CategoryField
	courseOptional.ScheduleTypeField
	courseOptional.NameField
	courseOptional.CoverField
	courseOptional.IntroField
	courseOptional.FoodField
	courseOptional.LevelField
	courseOptional.SuitField
	courseOptional.EquipmentField
	courseOptional.PlaceField
	courseOptional.TrainTargetField
	courseOptional.BodyTargetField
	courseOptional.NoticeField
	courseOptional.PlanCountField
	courseOptional.WorkoutCountField
	courseOptional.CreateAtField
	courseOptional.UpdateAtField
}
