package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/Henry19910227/fitness-go/internal/util"
	"github.com/gin-gonic/gin"
	"math"
	"time"
)

type rda struct {
	Base
	rdaRepo     repository.RDA
	bmrTool     tool.BMR
	tdeeTool    tool.TDEE
	calorieTool tool.Calorie
	errHandler  errcode.Handler
}

func NewRDA(rdaRepo repository.RDA, bmrTool tool.BMR, tdeeTool tool.TDEE, calorieTool tool.Calorie, errHandler errcode.Handler) RDA {
	return &rda{rdaRepo: rdaRepo, bmrTool: bmrTool, tdeeTool: tdeeTool, calorieTool: calorieTool, errHandler: errHandler}
}

func (r *rda) CreateRDA(c *gin.Context, userID int64, param *dto.RDA) errcode.Error {
	if err := r.rdaRepo.CreateRDA(nil, userID, &model.CreateRDAParam{
		TDEE:      param.TDEE,
		Calorie:   param.Calorie,
		Protein:   param.Protein,
		Fat:       param.Fat,
		Carbs:     param.Carbs,
		Grain:     param.Grain,
		Vegetable: param.Vegetable,
		Fruit:     param.Fruit,
		Meat:      param.Meat,
		Dairy:     param.Dairy,
		Nut:       param.Nut,
	}); err != nil {
		return r.errHandler.Set(c, "iap handler", err)
	}
	return nil
}
func (r *rda) CalculateRDA(param *dto.CalculateRDAParam) *dto.RDA {
	if param == nil {
		return &dto.RDA{}
	}
	bmr := r.CalculateBMR(param.Birthday, param.Weight, param.Height, param.BodyFat, param.Sex)
	tdee := r.CalculateTDEE(bmr, global.ActivityLevel(param.ActivityLevel), global.ExerciseFeqLevel(param.ExerciseFeqLevel))
	calorie := r.CalculateCalorie(tdee, global.DietTarget(param.DietTarget))
	proteinCal := r.CalculateProteinCalorie(calorie, global.DietTarget(param.DietTarget))
	fatCal := r.CalculateFatCalorie(calorie, global.DietTarget(param.DietTarget))
	carbsCal := r.CalculateCarbsCalorie(calorie, global.DietTarget(param.DietTarget))
	dairyAmt := r.CalculateDairyAmount(global.DietType(param.DietType))
	vegetableAmt := r.CalculateVegetableAmount()
	fruitAmt := r.CalculateFruitAmount()
	grainAmt := r.CalculateGrainAmount(r.CalculateCarbsAmount(carbsCal), dairyAmt, vegetableAmt, fruitAmt)
	meatAmt := r.CalculateMeatAmount(r.CalculateProteinAmount(proteinCal), dairyAmt, grainAmt, vegetableAmt)
	nutAmount := r.CalculateNutAmount(r.CalculateFatAmount(fatCal), dairyAmt, meatAmt)
	return &dto.RDA{
		TDEE:      tdee,
		Calorie:   calorie,
		Protein:   r.CalculateProteinAmount(proteinCal),
		Fat:       r.CalculateFatAmount(fatCal),
		Carbs:     r.CalculateCarbsAmount(carbsCal),
		Dairy:     dairyAmt,
		Vegetable: vegetableAmt,
		Fruit:     fruitAmt,
		Grain:     grainAmt,
		Meat:      meatAmt,
		Nut:       nutAmount,
	}
}

// CalculateBMR 計算BMR
func (r *rda) CalculateBMR(birthday string, weight float64, height float64, bodyFat *int, sex string) int {
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
func (r *rda) CalculateTDEE(bmr int, activityLevel global.ActivityLevel, exerciseFeqLevel global.ExerciseFeqLevel) int {
	value := r.tdeeTool.TDEE(bmr, activityLevel, exerciseFeqLevel)
	return int(math.Round(value))
}

// CalculateCalorie 計算目標下總熱量
func (r *rda) CalculateCalorie(tdee int, dietTarget global.DietTarget) int {
	value := r.calorieTool.TargetCalorie(tdee, dietTarget)
	return int(math.Round(value))
}

// CalculateProteinCalorie 計算蛋白質在各種目標下所需熱量
func (r *rda) CalculateProteinCalorie(calorie int, dietTarget global.DietTarget) int {
	if dietTarget == global.DietTargetLoseFat || dietTarget == global.DietTargetBuildMuscle {
		return int(float64(calorie) * 0.2)
	}
	value := float64(calorie) * 0.15
	return int(math.Round(value))
}

// CalculateFatCalorie 計算脂肪在各種目標下所需熱量
func (r *rda) CalculateFatCalorie(calorie int, dietTarget global.DietTarget) int {
	if dietTarget == global.DietTargetLoseFat {
		return int(math.Round(float64(calorie) * 0.3))
	}
	return int(math.Round(float64(calorie) * 0.20))
}

// CalculateCarbsCalorie 計算碳水化合物在各種目標下所需熱量
func (r *rda) CalculateCarbsCalorie(calorie int, dietTarget global.DietTarget) int {
	if dietTarget == global.DietTargetLoseFat {
		return int(math.Round(float64(calorie) * 0.5))
	}
	if dietTarget == global.DietTargetBuildMuscle {
		return int(math.Round(float64(calorie) * 0.6))
	}
	return int(math.Round(float64(calorie) * 0.65))
}

// CalculateProteinAmount 計算蛋白質份量(克)
func (r *rda) CalculateProteinAmount(proteinCal int) int {
	return int(math.Round(float64(proteinCal) / 4))
}

// CalculateCarbsAmount 計算碳水化合物份量(克)
func (r *rda) CalculateCarbsAmount(carbsCal int) int {
	return int(math.Round(float64(carbsCal) / 4))
}

// CalculateFatAmount 計算脂肪份量(克)
func (r *rda) CalculateFatAmount(fatCal int) int {
	return int(math.Round(float64(fatCal) / 9))
}

// CalculateDairyAmount 計算乳製品份量
func (r *rda) CalculateDairyAmount(dietType global.DietType) int {
	if dietType == global.DietTypeAllVegan || dietType == global.DietTypeOvoVegan {
		return 0
	}
	return 1
}

// CalculateVegetableAmount 計算蔬菜類份量
func (r *rda) CalculateVegetableAmount() int {
	return 5
}

// CalculateFruitAmount 計算水果類份量
func (r *rda) CalculateFruitAmount() int {
	return 2
}

// CalculateGrainAmount 計算穀物類份量
func (r *rda) CalculateGrainAmount(carbsAmt int, dairyAmt int, vegetableAmt int, fruitAmt int) int {
	value := (float64(carbsAmt) - (float64(dairyAmt)*12 + float64(vegetableAmt)*5 + float64(fruitAmt)*15)) / 15
	result := math.Round(value)
	return int(result)
}

// CalculateMeatAmount 計算蛋豆魚肉類份量
func (r *rda) CalculateMeatAmount(proteinAmt int, dairyAmt int, grainAmt int, vegetableAmt int) int {
	value := (float64(proteinAmt) - (float64(dairyAmt)*8 + float64(grainAmt)*2 + float64(vegetableAmt)*1)) / 7
	result := math.Round(value)
	return int(result)
}

// CalculateNutAmount 計算油脂堅果類份量
func (r *rda) CalculateNutAmount(fatAmt int, dairyAmt int, meatAmt int) int {
	value := (float64(fatAmt) - (float64(dairyAmt)*4 + float64(meatAmt)*3)) / 5
	result := math.Round(value)
	return int(result)
}
