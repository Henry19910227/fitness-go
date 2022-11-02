package trainer

import (
	bankAccountRequired "github.com/Henry19910227/fitness-go/internal/v2/field/bank_account/required"
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
	"github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"mime/multipart"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input
type WhereInput = where.Input
type JoinInput = join.Input
type CustomOrderByInput = orderBy.CustomInput

type FindInput struct {
	optional.UserIDField
	PreloadInput
}

type ListInput struct {
	optional.UserIDField
	JoinInput
	WhereInput
	PagingInput
	PreloadInput
	OrderByInput
	CustomOrderByInput
}

type FavoriteListInput struct {
	optional.UserIDField
	PagingInput
	PreloadInput
	OrderByInput
}

// APIGetTrainerProfileInput /v2/trainer/profile [PATCH]
type APIGetTrainerProfileInput struct {
	required.UserIDField
}

// APICreateTrainerInput /v2/trainer [POST]
type APICreateTrainerInput struct {
	required.UserIDField
	Form APICreateTrainerForm
	Avatar file.Input // 必填_教練形象照
	CartFontImage file.Input // 必填_身分證正面照片
	CartBackImage file.Input // 必填_身分證背面照片
	TrainerAlbumPhotos []*file.Input // 選填_教練相簿照片(可一次傳多張)
	CertificateImages []*file.Input // 選填_證照照片(可一次傳多張)
	AccountImage file.Input // 必填_帳戶照片
}
type APICreateTrainerForm struct {
	required.NameField
	required.NicknameField
	required.SkillField
	required.EmailField
	required.PhoneField
	required.AddressField
	required.IntroField
	required.ExperienceField
	optional.MottoField
	optional.FacebookURLField
	optional.InstagramURLField
	optional.YoutubeURLField
	bankAccountRequired.AccountNameField
	bankAccountRequired.AccountField
	bankAccountRequired.BranchField
	bankAccountRequired.BankCodeField
	CertificateNames []string  `json:"certificate_names" form:"certificate_names" binding:"omitempty,max=20" example:"A級教練證照,B級教練證照"` // 選填_證照名稱(需對應證照照片數量與順序)
}

// APIGetStoreTrainerInput /v2/store/trainer/{user_id} [GET]
type APIGetStoreTrainerInput struct {
	Uri APIGetStoreTrainerUri
}
type APIGetStoreTrainerUri struct {
	required.UserIDField
}

// APIGetStoreTrainersInput /v2/store/trainers [GET]
type APIGetStoreTrainersInput struct {
	required.UserIDField
	Query APIGetStoreTrainersQuery
}
type APIGetStoreTrainersQuery struct {
	OrderField *string `json:"order_field" form:"order_field" binding:"omitempty,oneof=latest popular" example:"latest"` // 排序類型(latest:最新/popular:熱門)-單選
	PagingInput
}

// APIUpdateCMSTrainerAvatarInput /v2/cms/trainer/avatar [PATCH]
type APIUpdateCMSTrainerAvatarInput struct {
	required.UserIDField
	CoverNamed string
	File       multipart.File
}

// APIGetFavoriteTrainersInput /v2/favorite/trainers [GET]
type APIGetFavoriteTrainersInput struct {
	required.UserIDField
	Form APIGetFavoriteTrainersForm
}
type APIGetFavoriteTrainersForm struct {
	PagingInput
}
