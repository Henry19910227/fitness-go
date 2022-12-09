package meal

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	"github.com/Henry19910227/fitness-go/internal/v2/model/meal/api_put_meals"
	"gorm.io/gorm"
)

type Resolver interface {
	APIPutMeals(tx *gorm.DB, input *api_put_meals.Input) (output api_put_meals.Output)
	APIGetMeals(input *model.APIGetMealsInput) (output model.APIGetMealsOutput)
}
