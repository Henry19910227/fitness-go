package review

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	reviewOptional "github.com/Henry19910227/fitness-go/internal/v2/field/review/optional"
	reviewImageOptional "github.com/Henry19910227/fitness-go/internal/v2/field/review_image/optional"
	userOptional "github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"
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
		userOptional.IDField
		userOptional.NicknameField
		userOptional.AvatarField
	} `json:"user,omitempty"`
	Course *struct {
		courseOptional.IDField
		courseOptional.NameField
	} `json:"course,omitempty"`
	Images []*struct {
		reviewImageOptional.IDField
		reviewImageOptional.ImageField
		reviewImageOptional.CreateAtField
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

// APIGetStoreCourseReviewsOutput /v2/store/course/{course_id}/reviews [GET]
type APIGetStoreCourseReviewsOutput struct {
	base.Output
	Data   *APIGetStoreCourseReviewsData `json:"data,omitempty"`
	Paging *paging.Output                `json:"paging,omitempty"`
}
type APIGetStoreCourseReviewsData []*struct {
	reviewOptional.IDField
	reviewOptional.ScoreField
	reviewOptional.BodyField
	reviewOptional.CreateAtField
	User *struct {
		userOptional.IDField
		userOptional.NicknameField
		userOptional.AvatarField
	} `json:"user,omitempty"`
	Course *struct {
		courseOptional.IDField
		courseOptional.NameField
	} `json:"course,omitempty"`
	Images []*struct {
		reviewImageOptional.IDField
		reviewImageOptional.ImageField
		reviewImageOptional.CreateAtField
	} `json:"images,omitempty"`
}

// APIGetStoreCourseReviewOutput /v2/store/course/review/{review_id} [GET]
type APIGetStoreCourseReviewOutput struct {
	base.Output
	Data *APIGetStoreCourseReviewData `json:"data,omitempty"`
}
type APIGetStoreCourseReviewData struct {
	reviewOptional.IDField
	reviewOptional.ScoreField
	reviewOptional.BodyField
	reviewOptional.CreateAtField
	User *struct {
		userOptional.IDField
		userOptional.NicknameField
		userOptional.AvatarField
	} `json:"user,omitempty"`
	Course *struct {
		courseOptional.IDField
		courseOptional.NameField
	} `json:"course,omitempty"`
	Images []*struct {
		reviewImageOptional.IDField
		reviewImageOptional.ImageField
		reviewImageOptional.CreateAtField
	} `json:"images,omitempty"`
}

// APICreateStoreCourseReviewOutput /v2/store/course/{course_id}/review [POST]
type APICreateStoreCourseReviewOutput struct {
	base.Output
}

// APIDeleteStoreCourseReviewOutput /v2/store/course/review/{review_id} [DELETE]
type APIDeleteStoreCourseReviewOutput struct {
	base.Output
}
