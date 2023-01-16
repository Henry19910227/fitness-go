package api_get_store_course

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	productLabelOptional "github.com/Henry19910227/fitness-go/internal/v2/field/product_label/optional"
	reviewRequired "github.com/Henry19910227/fitness-go/internal/v2/field/review_statistic/required"
	saleItemOptional "github.com/Henry19910227/fitness-go/internal/v2/field/sale_item/optional"
	trainerOptional "github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/store/course/{course_id} [GET]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
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
		trainerOptional.TrainerLevelField
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
		reviewRequired.ScoreTotalField
		reviewRequired.AmountField
		reviewRequired.FiveTotalField
		reviewRequired.FourTotalField
		reviewRequired.ThreeTotalField
		reviewRequired.TwoTotalField
		reviewRequired.OneTotalField
	} `json:"review_statistic"`
}
