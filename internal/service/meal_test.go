package service

import (
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMeal_CreateMeals(t *testing.T) {
	if err := PrepareMigrate(); err != nil {
		t.Fatalf(err.Error())
	}
	gormTool, sqlDB, err := PrepareGorm()
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer sqlDB.Close()
	// 創建user
	users := entity.NewMockUsers()
	if err := gormTool.DB().Create(&users).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建rda
	rda := entity.RDA{
		ID:       1,
		UserID:   users[0].ID,
		TDEE:     2000,
		CreateAt: "2022-05-10 00:00:00",
	}
	if err := gormTool.DB().Create(&rda).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建diet
	diet := entity.Diet{
		ID:         1,
		UserID:     users[0].ID,
		RdaID:      rda.ID,
		ScheduleAt: "2022-01-01",
		CreateAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := gormTool.DB().Create(&diet).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 創建food
	foods := entity.NewMockSystemFoods()
	if err := gormTool.DB().Create(&foods).Error; err != nil {
		t.Fatalf(err.Error())
	}
	// 驗證 create meals
	mealService := NewMockMealService(gormTool)
	items := make([]*dto.MealParamItem, 0)
	item := dto.MealParamItem{
		DietID: diet.ID,
		FoodID: foods[0].ID,
		Type:   1,
		Amount: 0.5,
	}
	items = append(items, &item)
	if err := mealService.CreateMeals(nil, &dto.CreateMealsParam{
		MealParamItems: items,
	}); err != nil {
		t.Fatalf(err.Msg())
	}
	var result entity.Meal
	if err := gormTool.DB().
		Table("meals").
		Take(&result).Error; err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, int64(1), result.DietID)
	assert.Equal(t, int64(1), result.FoodID)
	assert.Equal(t, 1, result.Type)
	assert.Equal(t, 0.5, result.Amount)
}
