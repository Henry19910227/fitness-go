package review

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	reviewOptional "github.com/Henry19910227/fitness-go/internal/v2/field/review/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/review_image"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user"
)

type Output struct {
	Table
	User   *user.Output           `json:"user,omitempty" gorm:"foreignKey:id;references:user_id"`
	Course *course.Output         `json:"course,omitempty" gorm:"foreignKey:id;references:course_id"`
	Images []*review_image.Output `json:"images,omitempty" gorm:"foreignKey:review_id;references:id"`
}

func (Output) TableName() string {
	return "reviews"
}

// APIGetCMSReviewsOutput /v2/cms/reviews [GET]
type APIGetCMSReviewsOutput struct {
	base.Output
	Data   APIGetCMSReviewsData `json:"data"`
	Paging *paging.Output       `json:"paging,omitempty"`
}
type APIGetCMSReviewsData []*struct {
	reviewOptional.IDField
	reviewOptional.ScoreField
	reviewOptional.BodyField
	reviewOptional.CreateAtField
	User *struct {
		user.IDField
		user.NicknameField
	} `json:"user,omitempty"`
	Course *struct {
		courseOptional.IDField
		courseOptional.NameField
	} `json:"course,omitempty"`
	Images []*struct {
		review_image.IDField
		review_image.ImageField
		review_image.CreateAtField
	} `json:"images,omitempty"`
}

// APIUpdateCMSReviewOutput /v2/cms/review/{review_id} [PATCH]
type APIUpdateCMSReviewOutput struct {
	base.Output
}

// APIDeleteCMSReviewOutput /v2/cms/review/{review_id} [DELETE]
type APIDeleteCMSReviewOutput struct {
	base.Output
}
