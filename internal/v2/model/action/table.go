package action

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/action/optional"
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
)

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.CourseIDField
	optional.NameField
	optional.SourceField
	optional.TypeField
	optional.CategoryField
	optional.BodyField
	optional.EquipmentField
	optional.IntroField
	optional.CoverField
	optional.VideoField
	optional.StatusField
	optional.IsDeletedField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "actions"
}

type CourseTable struct {
	courseOptional.IDField
	courseOptional.UserIDField
}

func (CourseTable) TableName() string {
	return "courses"
}