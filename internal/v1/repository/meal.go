package repository

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"gorm.io/gorm/clause"
	"time"
)

type meal struct {
	gorm tool.Gorm
}

func NewMeal(gorm tool.Gorm) Meal {
	return &meal{gorm: gorm}
}

func (m *meal) SaveMeals(param *model.SaveMealsParam) ([]int64, error) {
	if len(param.MealItems) == 0 {
		return []int64{}, nil
	}
	meals := make([]*entity.Meal, 0)
	for _, item := range param.MealItems {
		meal := entity.Meal{
			DietID:   item.DietID,
			FoodID:   item.FoodID,
			Amount:   item.Amount,
			Type:     item.Type,
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		meals = append(meals, &meal)
	}
	if err := m.gorm.DB().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "diet_id"}, {Name: "food_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"amount"}),
	}).Create(&meals).Error; err != nil {
		return nil, err
	}
	mealIDs := make([]int64, 0)
	for _, meal := range meals {
		mealIDs = append(mealIDs, meal.ID)
	}
	return mealIDs, nil
}

func (m *meal) FindMeals(param *model.FindMealsParam) ([]*model.Meal, error) {
	meals := make([]*model.Meal, 0)
	if err := m.gorm.DB().Find(&meals, param.MealIDs).Error; err != nil {
		return nil, err
	}
	return meals, nil
}

func (m *meal) FindMealOwner(mealID int64) (int64, error) {
	var userID int64
	if err := m.gorm.DB().
		Table("meals").
		Select("diets.user_id").
		Joins("INNER JOIN diets ON meals.diet_id = diets.id").
		Where("meals.id = ?", mealID).
		Take(&userID).Error; err != nil {
		return 0, err
	}
	return userID, nil
}

func (m *meal) DeleteMeal(mealID int64) error {
	if err := m.gorm.DB().Delete(&entity.Meal{}, mealID).Error; err != nil {
		return err
	}
	return nil
}
