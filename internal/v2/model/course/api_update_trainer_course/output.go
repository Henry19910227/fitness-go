package api_update_trainer_course

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	productLabelOptional "github.com/Henry19910227/fitness-go/internal/v2/field/product_label/optional"
	saleItemOptional "github.com/Henry19910227/fitness-go/internal/v2/field/sale_item/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/trainer/course/{course_id} [PATCH]
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
