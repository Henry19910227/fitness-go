package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/certificate"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer_albums"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer_statistic"
)

type Output struct {
	Table
	TrainerAlbums *[]trainer_albums.Output `json:"trainer_album_photos,omitempty" gorm:"foreignKey:user_id;references:user_id"`
	Certificates  *[]certificate.Output    `json:"certificates,omitempty" gorm:"foreignKey:user_id;references:user_id"`
	TrainerStatistic  *trainer_statistic.Output    `json:"trainer_statistic,omitempty" gorm:"foreignKey:user_id;references:user_id"`
}

func (Output) TableName() string {
	return "trainers"
}

// APIGetTrainerProfileOutput /v2/trainer/profile [PATCH]
type APIGetTrainerProfileOutput struct {
	base.Output
	Data *APIGetTrainerProfileData `json:"data,omitempty"`
}
type APIGetTrainerProfileData struct {
	Table
	Certificates []*struct{
		certificate.IDField
		certificate.ImageField
		certificate.NameField
	} `json:"certificates,omitempty"`
	TrainerAlbumPhotos []*struct{
		trainer_albums.IDField
		trainer_albums.PhotoField
		trainer_albums.CreateAtField
	} `json:"trainer_album_photos,omitempty"`
	TrainerStatistic *struct{
		trainer_statistic.CourseCount
		trainer_statistic.ReviewScore
		trainer_statistic.StudentCount
	} `json:"trainer_statistic,omitempty"`
}

// APIGetFavoriteTrainersOutput /v2/favorite/trainers [GET] 獲取
type APIGetFavoriteTrainersOutput struct {
	base.Output
	Data   APIGetFavoriteTrainersData `json:"data"`
	Paging *paging.Output             `json:"paging,omitempty"`
}
type APIGetFavoriteTrainersData []*struct {
	UserIDField
	NicknameField
	SkillField
	AvatarField
}

// APIUpdateCMSTrainerAvatarOutput /v2/cms/trainer/{user_id}/avatar [PATCH] 獲取課表詳細 API
type APIUpdateCMSTrainerAvatarOutput struct {
	base.Output
	Data *string `json:"data,omitempty" example:"123.jpg"`
}
