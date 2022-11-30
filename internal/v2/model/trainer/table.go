package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	userOptional "github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"
)

type Table struct {
	optional.UserIDField
	optional.NameField
	optional.NicknameField
	optional.SkillField
	optional.AvatarField
	optional.TrainerStatusField
	optional.TrainerLevelField
	optional.EmailField
	optional.PhoneField
	optional.AddressField
	optional.IntroField
	optional.ExperienceField
	optional.MottoField
	optional.FacebookURLField
	optional.InstagramURLField
	optional.YoutubeURLField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "trainers"
}

type UserTable struct {
	userOptional.IDField
	userOptional.AccountTypeField
	userOptional.AccountField
	userOptional.PasswordField
	userOptional.DeviceTokenField
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
	userOptional.IsDeletedField
	userOptional.CreateAtField
	userOptional.UpdateAtField
}

func (UserTable) TableName() string {
	return "users"
}
