package course

import (
	actionOptional "github.com/Henry19910227/fitness-go/internal/v2/field/action/optional"
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	planOptional "github.com/Henry19910227/fitness-go/internal/v2/field/plan/optional"
	productLabelOptional "github.com/Henry19910227/fitness-go/internal/v2/field/product_label/optional"
	saleItemOptional "github.com/Henry19910227/fitness-go/internal/v2/field/sale_item/optional"
	trainerOptional "github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/user/optional"
	workoutOptional "github.com/Henry19910227/fitness-go/internal/v2/field/workout/optional"
	workoutSetOptional "github.com/Henry19910227/fitness-go/internal/v2/field/workout_set/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course_training_avg_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/model/favorite_course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/model/review_statistic"
	saleItem "github.com/Henry19910227/fitness-go/internal/v2/model/sale_item"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_course_asset"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_course_statistic"
)

type Output struct {
	Table
	Trainer                    *trainer.Output                       `json:"trainer,omitempty" gorm:"foreignKey:user_id;references:user_id"`                    // 教練
	SaleItem                   *saleItem.Output                      `json:"sale_item,omitempty" gorm:"foreignKey:id;references:sale_id"`                       // 銷售項目
	ReviewStatistic            *review_statistic.Output              `json:"review_statistic,omitempty" gorm:"foreignKey:course_id;references:id"`              // 評分統計
	UserCourseStatistic        *user_course_statistic.Output         `json:"user_course_statistic,omitempty" gorm:"foreignKey:course_id;references:id"`         // 用戶課表統計
	UserCourseAsset            *user_course_asset.Output             `json:"user_course_asset,omitempty" gorm:"foreignKey:course_id;references:id"`             // 課表購買紀錄
	CourseTrainingAvgStatistic *course_training_avg_statistic.Output `json:"course_training_avg_statistic,omitempty" gorm:"foreignKey:course_id;references:id"` // 課表完成度統計
	FavoriteCourse             *favorite_course.Output               `json:"favorite_course,omitempty" gorm:"foreignKey:course_id;references:id"`               // 課表完成度統計
	Plans                      []*plan.Output                        `json:"plans,omitempty" gorm:"foreignKey:course_id;references:id"`                         // 計畫
}

func (Output) TableName() string {
	return "courses"
}

func (o *Output) UserCourseAssetOnSafe() user_course_asset.Output {
	if o.UserCourseAsset != nil {
		return *o.UserCourseAsset
	}
	return user_course_asset.Output{}
}

func (o *Output) FavoriteCourseOnSafe() favorite_course.Output {
	if o.FavoriteCourse != nil {
		return *o.FavoriteCourse
	}
	return favorite_course.Output{}
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

// APIUpdateUserCourseOutput /v2/user/course/{course_id} [PATCH]
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
	AllowAccess *int `json:"allow_access" example:"0"` // 是否允許訪問此課表(0:否/1:是)
	Favorite    *int `json:"favorite" example:"1"`     //是否收藏(0:否/1:是)
	Trainer     *struct {
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

// APIGetUserCourseStructureOutput /v2/user/course/{course_id}/structure [GET]
type APIGetUserCourseStructureOutput struct {
	base.Output
	Data *APIGetUserCourseStructureData `json:"data,omitempty"`
}
type APIGetUserCourseStructureData struct {
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
	courseOptional.CreateAtField
	courseOptional.UpdateAtField
	AllowAccess *int `json:"allow_access" example:"0"` // 是否允許訪問此課表(0:否/1:是)
	Favorite    *int `json:"favorite" example:"1"`     //是否收藏(0:否/1:是)
	Plans       []*struct {
		planOptional.IDField
		planOptional.NameField
		planOptional.WorkoutCountField
		planOptional.CreateAtField
		planOptional.UpdateAtField
		Workouts []*struct {
			workoutOptional.IDField
			workoutOptional.NameField
			workoutOptional.EquipmentField
			workoutOptional.StartAudioField
			workoutOptional.EndAudioField
			workoutOptional.WorkoutSetCountField
			workoutOptional.CreateAtField
			workoutOptional.UpdateAtField
			WorkoutSets []*struct {
				workoutSetOptional.IDField
				workoutSetOptional.TypeField
				workoutSetOptional.AutoNextField
				workoutSetOptional.StartAudioField
				workoutSetOptional.ProgressAudioField
				workoutSetOptional.RemarkField
				workoutSetOptional.WeightField
				workoutSetOptional.RepsField
				workoutSetOptional.DistanceField
				workoutSetOptional.DurationField
				workoutSetOptional.InclineField
				workoutSetOptional.CreateAtField
				workoutSetOptional.UpdateAtField
				Action *struct {
					actionOptional.IDField
					actionOptional.NameField
					actionOptional.SourceField
					actionOptional.TypeField
					actionOptional.CategoryField
					actionOptional.BodyField
					actionOptional.EquipmentField
					actionOptional.IntroField
					actionOptional.CoverField
					actionOptional.VideoField
					actionOptional.CreateAtField
					actionOptional.UpdateAtField
				} `json:"action,omitempty"`
			} `json:"workout_sets,omitempty"`
		} `json:"workouts,omitempty"`
	} `json:"plans,omitempty"`
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

// APIUpdateTrainerCourseOutput /v2/trainer/course/{course_id} [PATCH]
type APIUpdateTrainerCourseOutput struct {
	base.Output
	Data *APIUpdateTrainerCourseData `json:"data,omitempty"`
}
type APIUpdateTrainerCourseData struct {
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

// APIDeleteTrainerCourseOutput /v2/trainer/course/{course_id} [DELETE]
type APIDeleteTrainerCourseOutput struct {
	base.Output
}

// APISubmitTrainerCourseOutput /v2/trainer/course/{course_id}/submit [POST]
type APISubmitTrainerCourseOutput struct {
	base.Output
}

// APIGetProductCourseOutput /v2/product/course/{course_id} [GET]
type APIGetProductCourseOutput struct {
	base.Output
	Data *APIGetProductCourseData `json:"data,omitempty"`
}
type APIGetProductCourseData struct {
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
	courseOptional.CreateAtField
	courseOptional.UpdateAtField
	AllowAccess *int `json:"allow_access" example:"0"` // 是否允許訪問此課表(0:否/1:是)
	Favorite    *int `json:"favorite" example:"1"`     //是否收藏(0:否/1:是)
	Trainer     *struct {
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
		review_statistic.FiveTotalField
		review_statistic.FourTotalField
		review_statistic.ThreeTotalField
		review_statistic.TwoTotalField
		review_statistic.OneTotalField
	} `json:"review_statistic"`
}
