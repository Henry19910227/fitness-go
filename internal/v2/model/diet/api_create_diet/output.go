package api_create_diet

import (
	dietOptional "github.com/Henry19910227/fitness-go/internal/v2/field/diet/optional"
	rdaOptional "github.com/Henry19910227/fitness-go/internal/v2/field/rda/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	dietOptional.IDField
	dietOptional.ScheduleAtField
	dietOptional.CreateAtField
	dietOptional.UpdateAtField
	//Meals []*struct{
	//	mealOptional.IDField
	//	mealOptional.TypeField
	//	mealOptional.AmountField
	//	Food *struct{
	//		foodOptional.IDField
	//		foodOptional.UserIDField
	//		foodOptional.SourceField
	//		foodOptional.NameField
	//		foodOptional.CalorieField
	//		foodOptional.AmountDescField
	//		FoodCategory *struct{
	//			foodCategoryOptional.IDField
	//			foodCategoryOptional.TagField
	//			foodCategoryOptional.TitleField
	//		} `json:"food_category,omitempty"`
	//	} `json:"food,omitempty"`
	//} `json:"meals,omitempty"`
	RDA *struct {
		rdaOptional.TDEEField
		rdaOptional.CalorieField
		rdaOptional.ProteinField
		rdaOptional.FatField
		rdaOptional.CarbsField
		rdaOptional.GrainField
		rdaOptional.VegetableField
		rdaOptional.FruitField
		rdaOptional.MeatField
		rdaOptional.DairyField
		rdaOptional.NutField
	} `json:"rda,omitempty"`
}
