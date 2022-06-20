package meal

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	"gorm.io/gorm"
)

type Resolver interface {
	APIPutMeals(tx *gorm.DB, input *model.APIPutMealsInput) (output model.APIPutMealsOutput)
	APIGetMeals(input *model.APIGetMealsInput) (output model.APIGetMealsOutput)
}
