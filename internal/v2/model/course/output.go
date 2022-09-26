package course

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	productLabelOptional "github.com/Henry19910227/fitness-go/internal/v2/field/product_label/optional"
	saleItemOptional "github.com/Henry19910227/fitness-go/internal/v2/field/sale_item/optional"
	trainerOptional "github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course_training_avg_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/review_statistic"
	saleItem "github.com/Henry19910227/fitness-go/internal/v2/model/sale_item"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_course_statistic"
)

type Output struct {
	Table
	Trainer                    *trainer.Output                       `json:"trainer,omitempty" gorm:"foreignKey:user_id;references:user_id"`                    // 教練
	SaleItem                   *saleItem.Output                      `json:"sale_item,omitempty" gorm:"foreignKey:id;references:sale_id"`                       // 銷售項目
	ReviewStatistic            *review_statistic.Output              `json:"review_statistic,omitempty" gorm:"foreignKey:course_id;references:id"`              // 評分統計
	UserCourseStatistic        *user_course_statistic.Output         `json:"user_course_statistic,omitempty" gorm:"foreignKey:course_id;references:id"`         // 用戶課表統計
	CourseTrainingAvgStatistic *course_training_avg_statistic.Output `json:"course_training_avg_statistic,omitempty" gorm:"foreignKey:course_id;references:id"` // 課表完成度統計
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
	courseOptional.IDField
	courseOptional.SaleTypeField
	courseOptional.CategoryField
	courseOptional.ScheduleTypeField
	courseOptional.NameField
	courseOptional.CoverField
	courseOptional.LevelField
	courseOptional.TrainTargetField
	courseOptional.PlanCountField
	courseOptional.WorkoutCountField
	courseOptional.CreateAtField
	courseOptional.UpdateAtField
	Trainer *struct {
		trainerOptional.UserIDField
		trainerOptional.NicknameField
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
	courseOptional.IDField
	courseOptional.NameField
	courseOptional.CourseStatusField
	courseOptional.ScheduleTypeField
	courseOptional.SaleTypeField
	courseOptional.CreateAtField
	Trainer *struct {
		trainerOptional.UserIDField
		trainerOptional.NicknameField
	} `json:"trainer,omitempty"`
	SaleItem *struct {
		saleItemOptional.IDField
		saleItemOptional.NameField
		ProductLabel *struct {
			productLabelOptional.IDField
			productLabelOptional.ProductIDField
			productLabelOptional.TwdField
		} `json:"product_label,omitempty"`
	} `json:"sale_item,omitempty"`
}

// APIGetCMSCourseOutput /cms/course [GET] 獲取課表詳細 API
type APIGetCMSCourseOutput struct {
	base.Output
	Data *APIGetCMSCourseData `json:"data,omitempty"`
}
type APIGetCMSCourseData struct {
	courseOptional.IDField
	courseOptional.UserIDField
	courseOptional.SaleTypeField
	courseOptional.CourseStatusField
	courseOptional.CategoryField
	courseOptional.ScheduleTypeField
	courseOptional.NameField
	courseOptional.CoverField
	courseOptional.IntroField
	courseOptional.FoodField
	courseOptional.LevelField
	courseOptional.SuitField
	courseOptional.EquipmentField
	courseOptional.PlaceField
	courseOptional.TrainTargetField
	courseOptional.BodyTargetField
	courseOptional.NoticeField
	courseOptional.CreateAtField
	courseOptional.UpdateAtField
	SaleItem *struct {
		saleItemOptional.IDField
		saleItemOptional.NameField
		ProductLabel *struct {
			productLabelOptional.IDField
			productLabelOptional.ProductIDField
			productLabelOptional.TwdField
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
	courseOptional.IDField
}

// APIGetUserCoursesOutput /v2/user/courses [GET]
type APIGetUserCoursesOutput struct {
	base.Output
	Data   APIGetUserCoursesData `json:"data"`
	Paging *paging.Output        `json:"paging,omitempty"`
}
type APIGetUserCoursesData []*struct {
	courseOptional.IDField
	courseOptional.SaleTypeField
	courseOptional.CourseStatusField
	courseOptional.CategoryField
	courseOptional.ScheduleTypeField
	courseOptional.NameField
	courseOptional.CoverField
	courseOptional.PlanCountField
	courseOptional.WorkoutCountField
	courseOptional.CreateAtField
	courseOptional.UpdateAtField
	Trainer *struct {
		trainerOptional.UserIDField
		trainerOptional.AvatarField
		trainerOptional.NicknameField
		trainerOptional.SkillField
	} `json:"trainer,omitempty"`
	SaleItem *struct {
		saleItemOptional.IDField
		saleItemOptional.NameField
		ProductLabel *struct {
			productLabelOptional.IDField
			productLabelOptional.ProductIDField
			productLabelOptional.TwdField
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
	courseOptional.IDField
	courseOptional.SaleTypeField
	courseOptional.CourseStatusField
	courseOptional.CategoryField
	courseOptional.ScheduleTypeField
	courseOptional.NameField
	courseOptional.CoverField
	courseOptional.PlanCountField
	courseOptional.WorkoutCountField
	courseOptional.CreateAtField
	courseOptional.UpdateAtField
	Trainer *struct {
		trainerOptional.UserIDField
		trainerOptional.AvatarField
		trainerOptional.NicknameField
		trainerOptional.SkillField
	} `json:"trainer,omitempty"`
	SaleItem *struct {
		saleItemOptional.IDField
		saleItemOptional.NameField
		ProductLabel *struct {
			productLabelOptional.IDField
			productLabelOptional.ProductIDField
			productLabelOptional.TwdField
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
	courseOptional.IDField
	courseOptional.SaleTypeField
	courseOptional.CourseStatusField
	courseOptional.CategoryField
	courseOptional.ScheduleTypeField
	courseOptional.NameField
	courseOptional.CoverField
	courseOptional.LevelField
	courseOptional.PlanCountField
	courseOptional.WorkoutCountField
	courseOptional.CreateAtField
	courseOptional.UpdateAtField
}

// APICreateTrainerCourseOutput /v2/trainer/course [POST]
type APICreateTrainerCourseOutput struct {
	base.Output
	Data *APICreateTrainerCourseData `json:"data,omitempty"`
}
type APICreateTrainerCourseData struct {
	courseOptional.IDField
}

// APIGetTrainerCourseOutput /v2/trainer/course/{course_id} [GET]
type APIGetTrainerCourseOutput struct {
	base.Output
	Data *APIGetTrainerCourseData `json:"data,omitempty"`
}
type APIGetTrainerCourseData struct {
	courseOptional.IDField
	courseOptional.SaleTypeField
	courseOptional.SaleIDField
	courseOptional.CourseStatusField
	courseOptional.CategoryField
	courseOptional.ScheduleTypeField
	courseOptional.NameField
	courseOptional.CoverField
	courseOptional.IntroField
	courseOptional.FoodField
	courseOptional.LevelField
	courseOptional.SuitField
	courseOptional.EquipmentField
	courseOptional.PlaceField
	courseOptional.TrainTargetField
	courseOptional.BodyTargetField
	courseOptional.NoticeField
	courseOptional.PlanCountField
	courseOptional.WorkoutCountField
	optional.CreateAtField
	optional.UpdateAtField
	Trainer *struct {
		trainerOptional.UserIDField
		trainerOptional.AvatarField
		trainerOptional.NicknameField
		trainerOptional.SkillField
	} `json:"trainer,omitempty"`
	SaleItem *struct {
		saleItemOptional.IDField
		saleItemOptional.NameField
		ProductLabel *struct {
			productLabelOptional.IDField
			productLabelOptional.ProductIDField
			productLabelOptional.TwdField
		} `json:"product_label,omitempty"`
	} `json:"sale_item,omitempty"`
}
