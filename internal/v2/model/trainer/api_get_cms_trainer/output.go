package api_get_cms_trainer

import (
	bankAccountOptional "github.com/Henry19910227/fitness-go/internal/v2/field/bank_account/optional"
	cardOptional "github.com/Henry19910227/fitness-go/internal/v2/field/card/optional"
	certOptional "github.com/Henry19910227/fitness-go/internal/v2/field/certificate/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	trainerAlbumOptional "github.com/Henry19910227/fitness-go/internal/v2/field/trainer_album/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/cms/trainer/{user_id} [GET]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
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
	Certificates []*struct {
		certOptional.IDField
		certOptional.ImageField
		certOptional.NameField
	} `json:"certificates,omitempty"`
	Card struct {
		cardOptional.FrontImageField
		cardOptional.BackImageField
	} `json:"card,omitempty"`
	BankAccount struct {
		bankAccountOptional.AccountField
		bankAccountOptional.AccountImageField
		bankAccountOptional.AccountNameField
		bankAccountOptional.BankCodeField
		bankAccountOptional.BranchField
	} `json:"bank_account,omitempty"`
	TrainerAlbumPhotos []*struct {
		trainerAlbumOptional.IDField
		trainerAlbumOptional.PhotoField
		trainerAlbumOptional.CreateAtField
	} `json:"trainer_album_photos,omitempty"`
}
