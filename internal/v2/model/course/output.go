package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	productLabel "github.com/Henry19910227/fitness-go/internal/v2/model/product_label"
	"github.com/Henry19910227/fitness-go/internal/v2/model/review_statistic"
	saleItem "github.com/Henry19910227/fitness-go/internal/v2/model/sale_item"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_course_statistic"
)

type Output struct {
	Table
	Trainer             *trainer.Output               `json:"trainer,omitempty" gorm:"foreignKey:user_id;references:user_id"`            // 教練
	SaleItem            *saleItem.Output              `json:"sale_item,omitempty" gorm:"foreignKey:id;references:sale_id"`               // 銷售項目
	ReviewStatistic     *review_statistic.Output      `json:"review_statistic,omitempty" gorm:"foreignKey:course_id;references:id"`      // 評分統計
	UserCourseStatistic *user_course_statistic.Output `json:"user_course_statistic,omitempty" gorm:"foreignKey:course_id;references:id"` // 評分統計
}

func (Output) TableName() string {
	return "courses"
}

// APIGetFavoriteCoursesOutput /v2/favorite/courses [GET] 獲取收藏課表列表
type APIGetFavoriteCoursesOutput struct {
	base.Output
	Data   APIGetFavoriteCoursesData `json:"data"`
	Paging *paging.Output            `json:"paging,omitempty"`
}
type APIGetFavoriteCoursesData []*struct {
	IDField
	SaleTypeField
	CategoryField
	ScheduleTypeField
	NameField
	CoverField
	LevelField
	TrainTargetField
	PlanCountField
	WorkoutCountField
	CreateAtField
	UpdateAtField
	Trainer *struct {
		trainer.UserIDField
		trainer.NicknameField
	} `json:"trainer,omitempty"`
	ReviewStatistic struct {
		review_statistic.ScoreTotalRequired
		review_statistic.AmountRequired
	} `json:"review_statistic"`
}

// APIGetCMSCoursesOutput /cms/courses [GET] 獲取課表列表 API
type APIGetCMSCoursesOutput struct {
	base.Output
	Data   APIGetCMSCoursesData `json:"data"`
	Paging *paging.Output       `json:"paging,omitempty"`
}
type APIGetCMSCoursesData []*struct {
	IDField
	NameField
	CourseStatusField
	ScheduleTypeField
	SaleTypeField
	CreateAtField
	Trainer *struct {
		trainer.UserIDField
		trainer.NicknameField
	} `json:"trainer,omitempty"`
	SaleItem *struct {
		saleItem.IDField
		saleItem.NameField
		ProductLabel *struct {
			productLabel.IDField
			productLabel.ProductIDField
			productLabel.TwdField
		} `json:"product_label,omitempty"`
	} `json:"sale_item,omitempty"`
}

// APIGetCMSCourseOutput /cms/course [GET] 獲取課表詳細 API
type APIGetCMSCourseOutput struct {
	base.Output
	Data *APIGetCMSCourseData `json:"data,omitempty"`
}
type APIGetCMSCourseData struct {
	IDField
	UserIDField
	SaleTypeField
	CourseStatusField
	CategoryField
	ScheduleTypeField
	NameField
	CoverField
	IntroField
	FoodField
	LevelField
	SuitField
	EquipmentField
	PlaceField
	TrainTargetField
	BodyTargetField
	NoticeField
	CreateAtField
	UpdateAtField
	SaleItem *struct {
		saleItem.IDField
		saleItem.NameField
		ProductLabel *struct {
			productLabel.IDField
			productLabel.ProductIDField
			productLabel.TwdField
		} `json:"product_label,omitempty"`
	} `json:"sale_item,omitempty"`
}

// APIUpdateCMSCourseCoverOutput /v2/cms/course/{course_id}/cover [GET] 獲取課表詳細 API
type APIUpdateCMSCourseCoverOutput struct {
	base.Output
	Data *string `json:"data,omitempty" example:"123.jpg"`
}

// APICreateUserCourseOutput /v2/user/course [POST]
type APICreateUserCourseOutput struct {
	base.Output
	Data *APICreateUserCourseData `json:"data,omitempty"`
}
type APICreateUserCourseData struct {
	IDField
}

// APIGetUserCoursesOutput /v2/user/courses [GET]
type APIGetUserCoursesOutput struct {
	base.Output
	Data   APIGetUserCoursesData `json:"data"`
	Paging *paging.Output        `json:"paging,omitempty"`
}
type APIGetUserCoursesData []*struct {
	IDField
	SaleTypeField
	CourseStatusField
	CategoryField
	ScheduleTypeField
	NameField
	CoverField
	PlanCountField
	WorkoutCountField
	CreateAtField
	UpdateAtField
	Trainer *struct {
		trainer.UserIDField
		trainer.AvatarField
		trainer.NicknameField
		trainer.SkillField
	} `json:"trainer,omitempty"`
	SaleItem *struct {
		saleItem.IDField
		saleItem.NameField
		ProductLabel *struct {
			productLabel.IDField
			productLabel.ProductIDField
			productLabel.TwdField
		} `json:"product_label,omitempty"`
	} `json:"sale_item,omitempty"`
	ReviewStatistic struct {
		review_statistic.ScoreTotalRequired
		review_statistic.AmountRequired
	} `json:"review_statistic"`
}

// APIDeleteUserCourseOutput /v2/user/course/{course_id} [DELETE]
type APIDeleteUserCourseOutput struct {
	base.Output
}

// APIUpdateUserCourseOutput /v2/user/course/{course_id} [UPDATE]
type APIUpdateUserCourseOutput struct {
	base.Output
}

// APIGetUserCourseOutput /v2/user/course/{course_id} [GET]
type APIGetUserCourseOutput struct {
	base.Output
	Data *APIGetUserCourseData `json:"data,omitempty"`
}
type APIGetUserCourseData struct {
	IDField
	SaleTypeField
	CourseStatusField
	CategoryField
	ScheduleTypeField
	NameField
	CoverField
	PlanCountField
	WorkoutCountField
	CreateAtField
	UpdateAtField
	Trainer *struct {
		trainer.UserIDField
		trainer.AvatarField
		trainer.NicknameField
		trainer.SkillField
	} `json:"trainer,omitempty"`
	SaleItem *struct {
		saleItem.IDField
		saleItem.NameField
		ProductLabel *struct {
			productLabel.IDField
			productLabel.ProductIDField
			productLabel.TwdField
		} `json:"product_label,omitempty"`
	} `json:"sale_item,omitempty"`
	UserCourseStatistic *struct {
		user_course_statistic.FinishWorkoutCountField
		user_course_statistic.DurationField
	} `json:"user_course_statistic,omitempty"`
}

// APIGetTrainerCoursesOutput /v2/trainer/courses [GET]
type APIGetTrainerCoursesOutput struct {
	base.Output
	Data   *APIGetTrainerCoursesData `json:"data,omitempty"`
	Paging *paging.Output            `json:"paging,omitempty"`
}
type APIGetTrainerCoursesData []*struct {
	IDField
	SaleTypeField
	CourseStatusField
	CategoryField
	ScheduleTypeField
	NameField
	CoverField
	LevelField
	PlanCountField
	WorkoutCountField
	CreateAtField
	UpdateAtField
}

// APICreateTrainerCourseOutput /v2/trainer/course [POST]
type APICreateTrainerCourseOutput struct {
	base.Output
	Data *APICreateTrainerCourseData `json:"data,omitempty"`
}
type APICreateTrainerCourseData struct {
	IDField
}
