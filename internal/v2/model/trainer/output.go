package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	userOptional "github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/certificate"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer_albums"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer_statistic"
)

type Output struct {
	Table
	User             *UserOutput               `json:"user,omitempty" gorm:"foreignKey:id;references:user_id"`
	TrainerAlbums    *[]trainer_albums.Output  `json:"trainer_album_photos,omitempty" gorm:"foreignKey:user_id;references:user_id"`
	Certificates     *[]certificate.Output     `json:"certificates,omitempty" gorm:"foreignKey:user_id;references:user_id"`
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
		certificate.IDField
		certificate.ImageField
		certificate.NameField
	} `json:"certificates,omitempty"`
	TrainerAlbumPhotos []*struct {
		trainer_albums.IDField
		trainer_albums.PhotoField
		trainer_albums.CreateAtField
	} `json:"trainer_album_photos,omitempty"`
	TrainerStatistic *struct {
		trainer_statistic.CourseCount
		trainer_statistic.ReviewScore
		trainer_statistic.StudentCount
	} `json:"trainer_statistic,omitempty"`
}

// APIGetTrainerOutput /v2/trainer/{user_id} [GET]
type APIGetTrainerOutput struct {
	base.Output
	Data *APIGetTrainerData `json:"data,omitempty"`
}
type APIGetTrainerData struct {
	Table
	userOptional.IsDeletedField
	Certificates []*struct {
		certificate.IDField
		certificate.ImageField
		certificate.NameField
	} `json:"certificates,omitempty"`
	TrainerAlbumPhotos []*struct {
		trainer_albums.IDField
		trainer_albums.PhotoField
		trainer_albums.CreateAtField
	} `json:"trainer_album_photos,omitempty"`
	TrainerStatistic *struct {
		trainer_statistic.CourseCount
		trainer_statistic.ReviewScore
		trainer_statistic.StudentCount
	} `json:"trainer_statistic,omitempty"`
}

// APIGetTrainersOutput /v2/trainers [GET]
type APIGetTrainersOutput struct {
	base.Output
	Data *APIGetTrainersData `json:"data,omitempty"`
	Paging *paging.Output `json:"paging,omitempty"`
}
type APIGetTrainersData []*struct {
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
