package meal

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	dietModel "github.com/Henry19910227/fitness-go/internal/v2/model/diet"
	mealModel "github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
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
	findInput.ID = input.DietID
	dietOutput, err := r.dietService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if *input.UserID != *dietOutput.UserID {
		output.SetStatus(code.PermissionDenied)
		return output
	}
	//刪除meal
	delInput := mealModel.DeleteInput{}
	delInput.DietID = input.DietID
	if err := r.mealService.Tx(tx).Delete(&delInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser input
	items := make([]*mealModel.Table, 0)
	if err := util.Parser(input.Meals, &items); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	for _, item := range items {
		item.DietID = input.DietID
	}
	//新增meal
	if err := r.mealService.Tx(tx).Create(items); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIGetMeals(input *mealModel.APIGetMealsInput) (output mealModel.APIGetMealsOutput) {
	// parser input
	param := mealModel.ListInput{}
	param.UserID = input.UserID
	if err := util.Parser(input, &param); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// preload
	param.Preloads = []*preloadModel.Preload{
		{Field: "Food"},
		{Field: "Food.FoodCategory"},
	}
	// 調用 service
	result, _, err := r.mealService.List(&param)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := mealModel.APIGetMealsData{}
	if err := util.Parser(result, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = data
	return output
}
