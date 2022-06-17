package meal

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	"github.com/Henry19910227/fitness-go/internal/v2/service/meal"
	"gorm.io/gorm"
)

type resolver struct {
	mealService meal.Service
}

func New(mealService meal.Service) Resolver {
	return &resolver{mealService: mealService}
}

func (r *resolver) APIPutMeals(tx *gorm.DB, input *model.APIPutMealsInput) error {
	defer tx.Rollback()
	//刪除meal
	delInput := model.DeleteInput{}
	delInput.DietID = input.DietID
	if err := r.mealService.Tx(tx).Delete(&delInput); err != nil {
		return err
	}
	// parser input
	items := make([]*model.Table, 0)
	if err := util.Parser(input.Meals, &items); err != nil {
		return err
	}
	for _, item := range items {
		item.DietID = input.DietID
	}
	//新增meal
	if err := r.mealService.Tx(tx).Create(items); err != nil {
		return err
	}
	tx.Commit()
	return nil
}
