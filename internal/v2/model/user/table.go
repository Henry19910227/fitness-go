package user

import "github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"

type Table struct {
	optional.IDField
	optional.AccountTypeField
	optional.AccountField
	optional.PasswordField
	optional.DeviceTokenField
	optional.UserStatusField
	optional.UserTypeField
	optional.EmailField
	optional.NicknameField
	optional.AvatarField
	optional.SexField
	optional.BirthdayField
	optional.HeightField
	optional.WeightField
	optional.ExperienceField
	optional.TargetField
	optional.IsDeletedField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "users"
}
