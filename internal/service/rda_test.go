package service

import (
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/Henry19910227/fitness-go/internal/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 男性無填寫體脂肪
// 70 * 10 + 176 * 6.25 - 31 * 5 + 5
// => 700 + 1100 - 155 + 5
// = 1650
func TestCalculateBMRCase1(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	value := rdaService.CalculateBMR("1991-02-27", 70, 176, nil, "m")
	assert.Equal(t, 1650, value)

}

// 男性有填寫體脂肪
// ((70 * 10 + 176 * 6.25 - 31 * 5 + 5) + (370 + 21.6 * 70 * (100 - 20) / 100)) / 2
// => ((700 + 1100 - 155 + 5) + (370 + 21.6 * 70 * 0.8)) / 2
// => ((1650 + 1579) / 2
// => 3229 / 2
// = 1614.5
// = 1615 四捨五入
func TestCalculateBMRCase2(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	value := rdaService.CalculateBMR("1991-02-27", 70, 176, util.PointerInt(20), "m")
	assert.Equal(t, 1615, value)
}

// 女性無填寫體脂肪
// 70 * 10 + 176 * 6.25 - 31 * 5 - 161
// => 700 + 1100 - 155 - 161
// = 1484
func TestCalculateBMRCase3(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	value := rdaService.CalculateBMR("1991-02-27", 70, 176, nil, "f")
	assert.Equal(t, 1484, value)
}

// 女性有填寫體脂肪
// ((70 * 10 + 176 * 6.25 - 31 * 5 - 161) + (370 + 21.6 * 70 * (100 - 20) / 100)) / 2
// => ((700 + 1100 - 155 + 5) + (370 + 21.6 * 70 * 0.8)) / 2
// => ((1484 + 1579) / 2
// => 3063 / 2
// = 1531.5
// = 1532 四捨五入
func TestCalculateBMRCase4(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	value := rdaService.CalculateBMR("1991-02-27", 70, 176, util.PointerInt(20), "f")
	assert.Equal(t, 1532, value)
}

// TDEE測試1 植物人
// 1650 * 1.0 + 0 = 1650
func TestCalculateTDEECase1(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	value := rdaService.CalculateTDEE(1650, global.ActivityLevel1, global.ExerciseFeqLevel1)
	assert.Equal(t, 1650, value)
}

// TDEE測試2 每週輕度步行 3-4天 & 一週3-5次，一次45-60分鐘
// 1650 * 1.375 + 300 = 2568.75 = 2569 四捨五入
func TestCalculateTDEECase2(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	value := rdaService.CalculateTDEE(1650, global.ActivityLevel6, global.ExerciseFeqLevel3)
	assert.Equal(t, 2569, value)
}

// 建議熱量測試1 增肌
// 1650 * 1.15 = 1897.5 = 1897
func TestCalculateCalorieCase1(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	value := rdaService.CalculateCalorie(1650, global.DietTargetBuildMuscle)
	assert.Equal(t, 1897, value)
}

// 建議熱量測試2 哺乳者
// 1650 + 600 = 2250
func TestCalculateCalorieCase2(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	value := rdaService.CalculateCalorie(1650, global.DietTargetFeed)
	assert.Equal(t, 2250, value)
}

//減脂時期蛋白質克數 1650 * 0.2 / 4 = 82.5 = 83 四捨五入
func TestCalculateProteinAmount(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	value := rdaService.CalculateProteinCalorie(1650, global.DietTargetLoseFat)
	assert.Equal(t, 83, rdaService.CalculateProteinAmount(value))
}

//減脂時期碳水化合物克數 1650 * 0.5 / 4 = 206.25 = 206
func TestCalculateCarbsAmount(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	value := rdaService.CalculateCarbsCalorie(1650, global.DietTargetLoseFat)
	assert.Equal(t, 206, rdaService.CalculateCarbsAmount(value))
}

//減脂時期脂肪克數 1650 * 0.3 / 9 = 55
func TestCalculateFatAmount(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	value := rdaService.CalculateFatCalorie(1650, global.DietTargetLoseFat)
	assert.Equal(t, 55, rdaService.CalculateFatAmount(value))
}

func TestRda_CalculateGrainAmount(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	value := rdaService.CalculateGrainAmount(188, 1, 5, 2)
	assert.Equal(t, 8, value)
}

func TestRda_CalculateMeatAmount(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	value := rdaService.CalculateMeatAmount(75, 1, 8, 5)
	assert.Equal(t, 7, value)
}

func TestRda_CalculateNutAmount(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	value := rdaService.CalculateNutAmount(50, 1, 7)
	assert.Equal(t, 5, value)
}

//飲食計算機條件：
//標準飲食
//性別：男性
//年齡：30（生日：1992/2/2）
//身高：178cm
//體重：70kg
//體脂：20%
//活動量：每週輕度步行3-4天（1.375）
//運動量：中度 一週3-5次，一次45-60分鐘 （+300）
//飲食目標：增肌
//
//
//計算結果：
//TDEE=2533
//建議每日攝取熱量=2913
//蛋白質146克、脂肪65克、醣類437克
//乳製品1份、蔬菜5份、水果2份、全穀雜糧25份、蛋豆魚肉12份、油脂堅果5份

//—————————————以下為計算公式—————————————
//男性（有體脂）BMR公式={[10*體重(kg)+6.25*身高(cm)-5*年齡+5] + [370+21.6*(100% - 體脂率) * 體重]}/2
//BMR= {[10*70+6.25*178-5*30+5] + [370+21.6*(0.8) * 70]}/2= 1624 (1623.55 四捨五入）
//TDEE公式 = BMR * 活動量因子 + 運動量
//1624 * 1.375 + 300 = 2533
//飲食目標增肌建議每日攝取熱量 = TDEE *1.15
//2533*1.15 = 2913 (2912.95 四捨五入）
//增肌-三大營養素克數: 蛋白質20%、脂肪20%、醣類60%
//2913*0.2/4 = 146
//2913*0.2/9 = 65
//2913*0.6/4 = 437
//蛋白質146克、脂肪65克、醣類437克
//a乳製品 1
//b蔬菜 5
//c水果 2
//d全穀雜糧 = (437 – (1 * 12 + 5 * 5 + 2 * 15)) / 15 = 25
//e蛋豆魚肉 = (146 – (1 * 8 + 25 * 2 + 5 * 1)) / 7 = 12
//f油脂堅果 = (65 – (1 * 4 + 12 * 3)) / 5 = 5

func TestRda_CalculateRDA(t *testing.T) {
	rdaService := NewRDA(nil, tool.NewBMR(), tool.NewTDEE(), tool.NewCalorie(), nil)
	rda := rdaService.CalculateRDA(&dto.CalculateRDAParam{
		DietType:      1,
		Sex:           "m",
		Birthday:      "1992-02-02",
		Height:        178,
		Weight:        70,
		BodyFat:       util.PointerInt(20),
		ActivityLevel: 6,
		ExerciseFeq:   3,
		Target:        2,
	})
	assert.Equal(t, 2533, rda.TDEE)
	assert.Equal(t, 2913, rda.Calorie)
	assert.Equal(t, 146, rda.Protein)
	assert.Equal(t, 65, rda.Fat)
	assert.Equal(t, 437, rda.Carbs)
	assert.Equal(t, 1, rda.Dairy)
	assert.Equal(t, 5, rda.Vegetable)
	assert.Equal(t, 2, rda.Fruit)
	assert.Equal(t, 25, rda.Grain)
	assert.Equal(t, 12, rda.Meat)
	assert.Equal(t, 5, rda.Nut)
}
