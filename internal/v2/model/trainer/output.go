package trainer

type Table struct {
	UserIDField
	NameField
	NicknameField
	SkillField
	AvatarField
	TrainerStatusField
	TrainerLevelField
	EmailField
	PhoneField
	AddressField
	IntroField
	ExperienceField
	MottoField
	FacebookURLField
	InstagramURLField
	YoutubeURLField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "trainers"
}
