package api_create_trainer_action

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/action/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
)

// Input /v2/trainer/course/{course_id}/action [POST] 新增教練動作 API
type Input struct {
	userRequired.UserIDField
	Cover *file.Input
	Video *file.Input
	Uri   Uri
	Form  Form
}
type Uri struct {
	required.CourseIDField
}
type Form struct {
	required.NameField
	required.TypeField
	required.CategoryField
	required.BodyField
	required.EquipmentField
	required.IntroField
}
