package rda

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/bmr"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/food_calorie"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/tdee"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	dietModel "github.com/Henry19910227/fitness-go/internal/v2/model/diet"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/rda"
	"github.com/Henry19910227/fitness-go/internal/v2/model/rda/api_calculate_rda"
	"github.com/Henry19910227/fitness-go/internal/v2/model/rda/api_update_rda"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"github.com/Henry19910227/fitness-go/internal/v2/service/diet"
	"github.com/Henry19910227/fitness-go/internal/v2/service/rda"
	"gorm.io/gorm"
	"math"
	"time"
)

type resolver struct {
	rdaService  rda.Service
	dietService diet.Service
	tdeeTool    tdee.Tool
	bmrTool     bmr.Tool
	calorieTool food_calorie.Tool
}

func New(rdaService rda.Service, dietService diet.Service, tdeeTool tdee.Tool, bmrTool bmr.Tool, calorieTool food_calorie.Tool) Resolver {
	return &resolver{rdaService: rdaService, dietService: dietService, tdeeTool: tdeeTool, bmrTool: bmrTool, calorieTool: calorieTool}
}

func (r *resolver) APICalculateRDA(input *api_calculate_rda.Input) (output api_calculate_rda.Output) {
	// 計算熱量與營養素
	bmrValue := r.CalculateBMR(input.Body.Birthday, input.Body.Weight, input.Body.Height, input.Body.BodyFat, input.Body.Sex)
	tdeeValue := r.CalculateTDEE(bmrValue, input.Body.ActivityLevel, input.Body.ExerciseFeqLevel)
	calorie := r.CalculateCalorie(tdeeValue, input.Body.DietTarget)
	proteinCal := r.CalculateProteinCalorie(calorie, input.Body.DietTarget)
	fatCal := r.CalculateFatCalorie(calorie, input.Body.DietTarget)
	carbsCal := r.CalculateCarbsCalorie(calorie, input.Body.DietTarget)
	dairyAmt := r.CalculateDairyAmount(input.Body.DietType)
	vegetableAmt := r.CalculateVegetableAmount()
	fruitAmt := r.CalculateFruitAmount()
	grainAmt := r.CalculateGrainAmount(r.CalculateCarbsAmount(carbsCal), dairyAmt, vegetableAmt, fruitAmt)
	meatAmt := r.CalculateMeatAmount(r.CalculateProteinAmount(proteinCal), dairyAmt, grainAmt, vegetableAmt)
	nutAmount := r.CalculateNutAmount(r.CalculateFatAmount(fatCal), dairyAmt, meatAmt)

	// Parse Output
	data := api_calculate_rda.Data{}
	data.TDEE = util.PointerInt(tdeeValue)
	data.Calorie = util.PointerInt(calorie)
	data.Protein = util.PointerInt(r.CalculateProteinAmount(proteinCal))
	data.Fat = util.PointerInt(r.CalculateFatAmount(fatCal))
	data.Carbs = util.PointerInt(r.CalculateCarbsAmount(carbsCal))
	data.Dairy = util.PointerInt(dairyAmt)
	data.Vegetable = util.PointerInt(vegetableAmt)
	data.Fruit = util.PointerInt(fruitAmt)
	data.Grain = util.PointerInt(grainAmt)
	data.Meat = util.PointerInt(meatAmt)
	data.Nut = util.PointerInt(nutAmount)
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIUpdateRDA(tx *gorm.DB, input *api_update_rda.Input) (output api_update_rda.Output) {
	defer tx.Rollback()
	// 創建 rda
	rdaTable := model.Table{}
	rdaTable.UserID = util.PointerInt64(input.UserID)
	if err := util.Parser(input.Body, &rdaTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	rdaID, err := r.rdaService.Tx(tx).Create(&rdaTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查找今天之後的飲食紀錄
	dietListInput := dietModel.ListInput{}
	dietListInput.UserID = util.PointerInt64(input.UserID)
	dietListInput.Wheres = []*whereModel.Where{
		{Query: "DATE_FORMAT(diets.schedule_at,'%Y-%m-%d') >= DATE_FORMAT(?,'%Y-%m-%d')", Args: []interface{}{util.PointerString(time.Now().Format("2006-01-02"))}},
	}
	dietOutputs, _, err := r.dietService.Tx(tx).List(&dietListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新今天之後的飲食紀錄
	dietTables := make([]*dietModel.Table, 0)
	for _, dietOutput := range dietOutputs {
		dietTable := dietModel.Table{}
		dietTable.ID = dietOutput.ID
		dietTable.RdaID = util.PointerInt64(rdaID)
		dietTables = append(dietTables, &dietTable)
	}
	if err := r.dietService.Tx(tx).Updates(dietTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	output.Set(code.Success, "success")
	return output
}

// CalculateBMR 計算BMR
func (r *resolver) CalculateBMR(birthday string, weight float64, height float64, bodyFat *int, sex string) int {
	birthdayValue, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return 0
	}
	msjbmr := r.bmrTool.MSJBMR(weight, height, util.GetAge(birthdayValue), sex)
	if bodyFat == nil {
		return int(math.Round(msjbmr))
	}
	kmabmr := r.bmrTool.KMABMR(weight, *bodyFat)
	value := (msjbmr + kmabmr) / 2
	return int(math.Round(value))
}

// CalculateTDEE 計算TDEE
func (r *resolver) CalculateTDEE(bmr int, activityLevel int, exerciseFeqLevel int) int {
	value := r.tdeeTool.TDEE(bmr, activityLevel, exerciseFeqLevel)
	return int(math.Round(value))
}

// CalculateCalorie 計算目標下總熱量
func (r *resolver) CalculateCalorie(tdee int, dietTarget int) int {
	value := r.calorieTool.TargetCalorie(tdee, dietTarget)
	return int(math.Round(value))
}

// CalculateProteinCalorie 計算蛋白質在各種目標下所需熱量
func (r *resolver) CalculateProteinCalorie(calorie int, dietTarget int) int {
	if dietTarget == food_calorie.DietTargetLoseFat || dietTarget == food_calorie.DietTargetBuildMuscle {
		return int(float64(calorie) * 0.2)
	}
	value := float64(calorie) * 0.15
	return int(math.Round(value))
}

// CalculateFatCalorie 計算脂肪在各種目標下所需熱量
func (r *resolver) CalculateFatCalorie(calorie int, dietTarget int) int {
	if dietTarget == food_calorie.DietTargetLoseFat {
		return int(math.Round(float64(calorie) * 0.3))
	}
	return int(math.Round(float64(calorie) * 0.20))
}

// CalculateCarbsCalorie 計算碳水化合物在各種目標下所需熱量
func (r *resolver) CalculateCarbsCalorie(calorie int, dietTarget int) int {
	if dietTarget == food_calorie.DietTargetLoseFat {
		return int(math.Round(float64(calorie) * 0.5))
	}
	if dietTarget == food_calorie.DietTargetBuildMuscle {
		return int(math.Round(float64(calorie) * 0.6))
	}
	return int(math.Round(float64(calorie) * 0.65))
}

// CalculateProteinAmount 計算蛋白質份量(克)
func (r *resolver) CalculateProteinAmount(proteinCal int) int {
	return int(math.Round(float64(proteinCal) / 4))
}

// CalculateCarbsAmount 計算碳水化合物份量(克)
func (r *resolver) CalculateCarbsAmount(carbsCal int) int {
	return int(math.Round(float64(carbsCal) / 4))
}

// CalculateFatAmount 計算脂肪份量(克)
func (r *resolver) CalculateFatAmount(fatCal int) int {
	return int(math.Round(float64(fatCal) / 9))
}

// CalculateDairyAmount 計算乳製品份量
func (r *resolver) CalculateDairyAmount(dietType int) int {
	if dietType == food_calorie.DietTypeAllVegan || dietType == food_calorie.DietTypeOvoVegan {
		return 0
	}
	return 1
}

// CalculateVegetableAmount 計算蔬菜類份量
func (r *resolver) CalculateVegetableAmount() int {
	return 5
}

// CalculateFruitAmount 計算水果類份量
func (r *resolver) CalculateFruitAmount() int {
	return 2
}

// CalculateGrainAmount 計算穀物類份量
func (r *resolver) CalculateGrainAmount(carbsAmt int, dairyAmt int, vegetableAmt int, fruitAmt int) int {
	value := (float64(carbsAmt) - (float64(dairyAmt)*12 + float64(vegetableAmt)*5 + float64(fruitAmt)*15)) / 15
	result := math.Round(value)
	return int(result)
}

// CalculateMeatAmount 計算蛋豆魚肉類份量
func (r *resolver) CalculateMeatAmount(proteinAmt int, dairyAmt int, grainAmt int, vegetableAmt int) int {
	value := (float64(proteinAmt) - (float64(dairyAmt)*8 + float64(grainAmt)*2 + float64(vegetableAmt)*1)) / 7
	result := math.Round(value)
	return int(result)
}

// CalculateNutAmount 計算油脂堅果類份量
func (r *resolver) CalculateNutAmount(fatAmt int, dairyAmt int, meatAmt int) int {
	value := (float64(fatAmt) - (float64(dairyAmt)*4 + float64(meatAmt)*3)) / 5
	result := math.Round(value)
	return int(result)
}
