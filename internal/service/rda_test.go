package service

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 男性無填寫體脂肪
// 70 * 10 + 176 * 6.25 - 31 * 5 + 5
// => 700 + 1100 - 155 + 5
// = 1650
func TestCalculateBMRCase1(t *testing.T) {
	rdaService := NewRDAService()
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
	rdaService := NewRDAService()
	value := rdaService.CalculateBMR("1991-02-27", 70, 176, util.PointerInt(20), "m")
	assert.Equal(t, 1615, value)
}

// 女性無填寫體脂肪
// 70 * 10 + 176 * 6.25 - 31 * 5 - 161
// => 700 + 1100 - 155 - 161
// = 1484
func TestCalculateBMRCase3(t *testing.T) {
	rdaService := NewRDAService()
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
	rdaService := NewRDAService()
	value := rdaService.CalculateBMR("1991-02-27", 70, 176, util.PointerInt(20), "f")
	assert.Equal(t, 1532, value)
}

// TDEE測試1 植物人
// 1650 * 1.0 + 0 = 1650
func TestCalculateTDEECase1(t *testing.T) {
	rdaService := NewRDAService()
	value := rdaService.CalculateTDEE(1650, global.ActivityLevel1, global.ExerciseFeqLevel1)
	assert.Equal(t, 1650, value)
}

// TDEE測試2 每週輕度步行 3-4天 & 一週3-5次，一次45-60分鐘
// 1650 * 1.375 + 300 = 2568.75 = 2569 四捨五入
func TestCalculateTDEECase2(t *testing.T) {
	rdaService := NewRDAService()
	value := rdaService.CalculateTDEE(1650, global.ActivityLevel6, global.ExerciseFeqLevel3)
	assert.Equal(t, 2569, value)
}

// 建議熱量測試1 增肌
// 1650 * 1.15 = 1897.5 = 1897
func TestCalculateCalorieCase1(t *testing.T) {
	rdaService := NewRDAService()
	value := rdaService.CalculateCalorie(1650, global.DietTargetBuildMuscle)
	assert.Equal(t, 1897, value)
}

// 建議熱量測試2 哺乳者
// 1650 + 600 = 2250
func TestCalculateCalorieCase2(t *testing.T) {
	rdaService := NewRDAService()
	value := rdaService.CalculateCalorie(1650, global.DietTargetFeed)
	assert.Equal(t, 2250, value)
}

//減脂時期蛋白質克數 1650 * 0.2 / 4 = 82.5 = 83 四捨五入
func TestCalculateProteinAmount(t *testing.T) {
	rdaService := NewRDAService()
	value := rdaService.CalculateProteinCalorie(1650, global.DietTargetLoseFat)
	assert.Equal(t, 83, rdaService.CalculateProteinAmount(value))
}

//減脂時期碳水化合物克數 1650 * 0.5 / 4 = 206.25 = 206
func TestCalculateCarbsAmount(t *testing.T) {
	rdaService := NewRDAService()
	value := rdaService.CalculateCarbsCalorie(1650, global.DietTargetLoseFat)
	assert.Equal(t, 206, rdaService.CalculateCarbsAmount(value))
}

//減脂時期脂肪克數 1650 * 0.3 / 9 = 55
func TestCalculateFatAmount(t *testing.T) {
	rdaService := NewRDAService()
	value := rdaService.CalculateFatCalorie(1650, global.DietTargetLoseFat)
	assert.Equal(t, 55, rdaService.CalculateFatAmount(value))
}

func TestRda_CalculateGrainAmount(t *testing.T) {
	rdaService := NewRDAService()
	value := rdaService.CalculateGrainAmount(750, 1, 5, 2)
	assert.Equal(t, 8, value)
}

func TestRda_CalculateMeatAmount(t *testing.T) {
	rdaService := NewRDAService()
	value := rdaService.CalculateMeatAmount(300, 1, 8, 5)
	assert.Equal(t, 7, value)
}

func TestRda_CalculateNutAmount(t *testing.T) {
	rdaService := NewRDAService()
	value := rdaService.CalculateNutAmount(450, 1, 7)
	assert.Equal(t, 5, value)
}
