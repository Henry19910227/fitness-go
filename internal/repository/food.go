package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"time"
)

type food struct {
	gorm tool.Gorm
}

func NewFood(gorm tool.Gorm) Food {
	return &food{gorm: gorm}
}

func (f *food) CreateFood(param *model.CreateFoodParam) (int64, error) {
	food := entity.Food{
		UserID: param.UserID,
		FoodCategoryID: param.FoodCategoryID,
		Source: param.Source,
		Name: param.Name,
		Calorie: param.Calorie,
		AmountDesc: param.AmountDesc,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := f.gorm.DB().Create(&food).Error; err != nil {
		return 0, err
	}
	return food.ID, nil
}

func (f *food) FindFood(foodID int64, preloads []*model.Preload) (*model.Food, error) {
	//設置表
	db := f.gorm.DB().Model(&model.Food{})
	//關聯加載
	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = db.Preload(preload.Field)
		}
	}
	//查找數據
	var food model.Food
	if err := db.Where("id = ?", foodID).Take(&food).Error; err != nil {
		return nil, err
	}
	return &food, nil
}


func (f *food) FindFoods(param *model.FindFoodsParam) ([]*model.Food, error) {
	//設置篩選
	query := "1=1 "
	params := make([]interface{}, 0)
	if param.IsDeleted != nil {
		query += "AND foods.is_deleted = ? "
		params = append(params, *param.IsDeleted)
	}
	if param.UserID != nil {
		query += "AND (foods.user_id = ? OR foods.user_id IS NULL) "
		params = append(params, *param.UserID)
	}
	if param.Tag != nil {
		query += "AND food_categories.tag = ? "
		params = append(params, *param.Tag)
	}
	//設置表
	db := f.gorm.DB().Model(&model.Food{})
	//關聯加載
	if len(param.Preloads) > 0 {
		for _, preload := range param.Preloads {
			db = db.Preload(preload.Field)
		}
	}
	//查找數據
	foods := make([]*model.Food, 0)
	if err := db.
		Joins("INNER JOIN food_categories ON foods.food_category_id = food_categories.id").
		Where(query, params...).
		Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}