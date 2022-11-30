package api_update_cms_trainer

import (
	certOptional "github.com/Henry19910227/fitness-go/internal/v2/field/certificate/required"
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	trainerAlbumOptional "github.com/Henry19910227/fitness-go/internal/v2/field/trainer_album/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/cms/trainer/{user_id} [PATCH]
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
	TrainerAlbumPhotos []*struct {
		trainerAlbumOptional.IDField
		trainerAlbumOptional.PhotoField
		trainerAlbumOptional.CreateAtField
	} `json:"trainer_album_photos,omitempty"`
}
