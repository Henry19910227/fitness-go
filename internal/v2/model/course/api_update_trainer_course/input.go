package api_update_trainer_course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
)

// Input /v2/trainer/course/{course_id} [PATCH]
type Input struct {
	required.UserIDField
	Cover *file.Input
	Uri   Uri
	Form  Form
}
type Uri struct {
	required.CourseIDField
}
type Form struct {
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
