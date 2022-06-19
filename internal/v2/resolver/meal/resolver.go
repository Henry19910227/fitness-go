package meal

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	dietModel "github.com/Henry19910227/fitness-go/internal/v2/model/diet"
	mealModel "github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	"github.com/Henry19910227/fitness-go/internal/v2/service/diet"
	"github.com/Henry19910227/fitness-go/internal/v2/service/meal"
	"gorm.io/gorm"
)

type resolver struct {
	mealService meal.Service
	dietService diet.Service
}

func New(mealService meal.Service, dietService diet.Service) Resolver {
	return &resolver{mealService: mealService, dietService: dietService}
}

func (r *resolver) APIPutMeals(tx *gorm.DB, input *mealModel.APIPutMealsInput) (output mealModel.APIPutMealsOutput) {
	defer tx.Rollback()
	findInput := dietModel.FindInput{}
	dietOutput, err := r.dietService.Find(&findInput)
	if err != nil {
		output.SetStatus(code.BadRequest)
		return output
	}
	if input.UserID != util.OnNilJustReturnInt64(dietOutput.UserID, 0) {
		output.SetStatus(code.PermissionDenied)
		return output
	}
	//刪除meal
	delInput := mealModel.DeleteInput{}
	delInput.DietID = input.DietID
	if err := r.mealService.Tx(tx).Delete(&delInput); err != nil {
		output.SetStatus(code.BadRequest)
		return output
	}
	// parser input
	items := make([]*mealModel.Table, 0)
	if err := util.Parser(input.Meals, &items); err != nil {
		output.SetStatus(code.BadRequest)
		return output
	}
	for _, item := range items {
		item.DietID = input.DietID
	}
	//新增meal
	if err := r.mealService.Tx(tx).Create(items); err != nil {
		output.SetStatus(code.BadRequest)
		return output
	}
	tx.Commit()
	output.SetStatus(code.Success)
	return output
}
