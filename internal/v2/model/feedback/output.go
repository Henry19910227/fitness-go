package feedback

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/feedback_image"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user"
)

type Output struct {
	Table
	Images []*feedback_image.Output `json:"images,omitempty" gorm:"foreignKey:feedback_id;references:id"` // 反饋圖片
	User   *user.Output             `json:"user,omitempty" gorm:"foreignKey:id;references:user_id"`       // 用戶
}

func (Output) TableName() string {
	return "feedbacks"
}

// APICreateFeedbackOutput /v2/feedback [POST]
type APICreateFeedbackOutput struct {
	base.Output
}

// APIGetCMSFeedbacksOutput /v2/cms/feedbacks [GET]
type APIGetCMSFeedbacksOutput struct {
	base.Output
	Data   APIGetCMSFeedbacksData `json:"data"`
	Paging *paging.Output         `json:"paging,omitempty"`
}
type APIGetCMSFeedbacksData []*struct {
	IDField
	VersionField
	PlatformField
	OSVersionField
	PhoneModelField
	BodyField
	CreateAtField
	UpdateAtField
	Images []*struct {
		feedback_image.IDField
		feedback_image.ImageField
		feedback_image.CreateAtField
	} `json:"images,omitempty"`
	User *struct {
		user.IDField
		user.NicknameField
	} `json:"user,omitempty"`
}
