package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
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
