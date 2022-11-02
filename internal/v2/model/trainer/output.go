package trainer

import (
	certOptional "github.com/Henry19910227/fitness-go/internal/v2/field/certificate/required"
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	trainerAlbumOptional "github.com/Henry19910227/fitness-go/internal/v2/field/trainer_album/optional"
	trainerStatisticRequired "github.com/Henry19910227/fitness-go/internal/v2/field/trainer_statistic/required"
	userOptional "github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/certificate"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer_album"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer_statistic"
)

type Output struct {
	Table
	User             *UserOutput             `json:"user,omitempty" gorm:"foreignKey:id;references:user_id"`
	TrainerAlbums    *[]trainer_album.Output `json:"trainer_album_photos,omitempty" gorm:"foreignKey:user_id;references:user_id"`
	Certificates     *[]certificate.Output   `json:"certificates,omitempty" gorm:"foreignKey:user_id;references:user_id"`
	TrainerStatistic *trainer_statistic.Output `json:"trainer_statistic,omitempty" gorm:"foreignKey:user_id;references:user_id"`
}

func (Output) TableName() string {
	return "trainers"
}

type UserOutput struct {
	UserTable
}

func (UserOutput) TableName() string {
	return "users"
}

func (o *Output) UserOnSafe() UserOutput {
	if o.User != nil {
		return *o.User
	}
	return UserOutput{}
}

func (o *Output) TrainerStatisticOnSafe() trainer_statistic.Output {
	if o.TrainerStatistic != nil {
		return *o.TrainerStatistic
	}
	return trainer_statistic.Output{}
}

// APIGetTrainerProfileOutput /v2/trainer/profile [PATCH]
type APIGetTrainerProfileOutput struct {
	base.Output
	Data *APIGetTrainerProfileData `json:"data,omitempty"`
}
type APIGetTrainerProfileData struct {
	Table
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
	TrainerStatistic *struct {
		trainerStatisticRequired.CourseCountField
		trainerStatisticRequired.ReviewScoreField
		trainerStatisticRequired.StudentCountField
	} `json:"trainer_statistic,omitempty"`
}

// APICreateTrainerOutput /v2/trainer [POST]
type APICreateTrainerOutput struct {
	base.Output
	Data *APICreateTrainerData `json:"data,omitempty"`
}
type APICreateTrainerData struct {
	Table
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
	TrainerStatistic *struct {
		trainerStatisticRequired.CourseCountField
		trainerStatisticRequired.ReviewScoreField
		trainerStatisticRequired.StudentCountField
	} `json:"trainer_statistic,omitempty"`
}

// APIGetStoreTrainerOutput /v2/store/trainer/{user_id} [GET]
type APIGetStoreTrainerOutput struct {
	base.Output
	Data *APIGetStoreTrainerData `json:"data,omitempty"`
}
type APIGetStoreTrainerData struct {
	Table
	userOptional.IsDeletedField
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
	TrainerStatistic *struct {
		trainerStatisticRequired.CourseCountField
		trainerStatisticRequired.ReviewScoreField
		trainerStatisticRequired.StudentCountField
	} `json:"trainer_statistic,omitempty"`
}

// APIGetStoreTrainersOutput /v2/store/trainers [GET]
type APIGetStoreTrainersOutput struct {
	base.Output
	Data *APIGetStoreTrainersData `json:"data,omitempty"`
	Paging *paging.Output         `json:"paging,omitempty"`
}
type APIGetStoreTrainersData []*struct {
	optional.UserIDField
	optional.NicknameField
	optional.SkillField
	optional.AvatarField
	optional.CreateAtField
	optional.UpdateAtField
}

// APIGetFavoriteTrainersOutput /v2/favorite/trainers [GET] 獲取
type APIGetFavoriteTrainersOutput struct {
	base.Output
	Data   APIGetFavoriteTrainersData `json:"data"`
	Paging *paging.Output             `json:"paging,omitempty"`
}
type APIGetFavoriteTrainersData []*struct {
	optional.UserIDField
	optional.NicknameField
	optional.SkillField
	optional.AvatarField
}

// APIUpdateCMSTrainerAvatarOutput /v2/cms/trainer/{user_id}/avatar [PATCH] 獲取課表詳細 API
type APIUpdateCMSTrainerAvatarOutput struct {
	base.Output
	Data *string `json:"data,omitempty" example:"123.jpg"`
}
