package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
)

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.SaleTypeField
	optional.SaleIDField
	optional.CourseStatusField
	optional.CategoryField
	optional.ScheduleTypeField
	optional.NameField
	optional.CoverField
	optional.IntroField
	optional.FoodField
	optional.LevelField
	optional.SuitField
	optional.EquipmentField
	optional.PlaceField
	optional.TrainTargetField
	optional.BodyTargetField
	optional.NoticeField
	optional.PlanCountField
	optional.WorkoutCountField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "courses"
}
