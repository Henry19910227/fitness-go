package repository

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
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
		UserID:         param.UserID,
		FoodCategoryID: param.FoodCategoryID,
		Source:         param.Source,
		Name:           param.Name,
		Calorie:        param.Calorie,
		AmountDesc:     param.AmountDesc,
		CreateAt:       time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:       time.Now().Format("2006-01-02 15:04:05"),
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

func (f *food) UpdateFood(param *model.UpdateFoodParam) error {
	if param == nil {
		return nil
	}
	selects := make([]interface{}, 0)
	if param.IsDeleted != nil {
		selects = append(selects, "is_deleted")
	}
	selects = append(selects, "update_at")
	param.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	if err := f.gorm.DB().
		Table("foods").
		Where("id = ?", param.FoodID).
		Select("", selects...).
		Updates(param).Error; err != nil {
		return err
	}
	return nil
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

func (f *food) FindRecentFoods(param *model.FindRecentFoodsParam) ([]*model.RecentFood, error) {
	query := "1=1 "
	params := make([]interface{}, 0)
	if param.UserID != nil {
		query += "AND (foods.user_id = ? OR foods.user_id IS NULL) "
		params = append(params, *param.UserID)
	}
	if param.IsDeleted != nil {
		query += "AND foods.is_deleted = ? "
		params = append(params, *param.IsDeleted)
	}
	foods := make([]*model.RecentFood, 0)
	db := f.gorm.DB().
		Table("foods").
		Select("MAX(foods.id) AS id",
			"MAX(foods.food_category_id) AS food_category_id",
			"MAX(foods.user_id) AS user_id",
			"MAX(meals.id) AS meal_id",
			"MAX(foods.source) AS source",
			"MAX(foods.`name`) AS `name`",
			"MAX(foods.calorie) AS `calorie`",
			"MAX(foods.amount_desc) AS amount_desc",
			"MAX(foods.is_deleted) AS is_deleted",
			"MAX(foods.create_at) AS create_at",
			"MAX(foods.update_at) AS update_at")
	if len(param.Preloads) > 0 {
		for _, preload := range param.Preloads {
			db = db.Preload(preload.Field)
		}
	}
	if err := db.
		Joins("INNER JOIN meals ON meals.food_id = foods.id").
		Where(query, params...).
		Group("meals.food_id,meals.type").
		Order("meals.type ASC").
		Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}

func (f *food) FindFoodOwner(foodID int64) (int64, error) {
	var userID int64
	if err := f.gorm.DB().
		Table("foods").
		Select("IFNULL(user_id,0) AS user_id").
		Where("id = ?", foodID).
		Take(&userID).Error; err != nil {
		return 0, err
	}
	return userID, nil
}
