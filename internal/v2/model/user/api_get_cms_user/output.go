package api_get_cms_user

import (
	userOptional "github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/cms/course/{course_id}/users [GET]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	userOptional.IDField
	userOptional.AccountTypeField
	userOptional.UserStatusField
	userOptional.UserTypeField
	userOptional.EmailField
	userOptional.NicknameField
	userOptional.AvatarField
	userOptional.SexField
	userOptional.BirthdayField
	userOptional.HeightField
	userOptional.WeightField
	userOptional.ExperienceField
	userOptional.TargetField
	userOptional.CreateAtField
	userOptional.UpdateAtField
}
