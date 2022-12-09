package meal

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	dietModel "github.com/Henry19910227/fitness-go/internal/v2/model/diet"
	mealModel "github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	"github.com/Henry19910227/fitness-go/internal/v2/model/meal/api_put_meals"
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

func (r *resolver) APIPutMeals(tx *gorm.DB, input *api_put_meals.Input) (output api_put_meals.Output) {
	defer tx.Rollback()
	// 查詢diet資訊
	findInput := dietModel.FindInput{}
	findInput.ID = input.Uri.DietID
	dietOutput, err := r.dietService.Tx(tx).Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt64(dietOutput.UserID, 0) != input.UserID {
		output.Set(code.BadRequest, "非此diet創建者，無法修改meals")
		return output
	}
	// 刪除舊的meal
	delInput := mealModel.DeleteInput{}
	delInput.DietID = input.Uri.DietID
	if err := r.mealService.Tx(tx).Delete(&delInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 新增meal
	mealTables := make([]*mealModel.Table, 0)
	for _, mealItem := range input.Body {
		mealTable := mealModel.Table{}
		mealTable.DietID = input.Uri.DietID
		mealTable.FoodID = mealItem.FoodID
		mealTable.Type = mealItem.Type
		mealTable.Amount = mealItem.Amount
		mealTables = append(mealTables, &mealTable)
	}
	if err := r.mealService.Tx(tx).Create(mealTables); err != nil {
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
