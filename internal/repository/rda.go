package repository

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type rda struct {
	gorm tool.Gorm
}

func NewRDA(gorm tool.Gorm) RDA {
	return &rda{gorm: gorm}
}

func (r *rda) CreateRDA(tx *gorm.DB, userID int64, param *model.CreateRDAParam) (int64, error) {
	db := r.gorm.DB()
	if tx != nil {
		db = tx
	}
	rda := entity.RDA{
		UserID:    userID,
		TDEE:      param.TDEE,
		Calorie:   param.Calorie,
		Protein:   param.Protein,
		Fat:       param.Protein,
		Carbs:     param.Carbs,
		Grain:     param.Grain,
		Vegetable: param.Vegetable,
		Fruit:     param.Fruit,
		Meat:      param.Meat,
		Dairy:     param.Dairy,
		Nut:       param.Nut,
		CreateAt:  time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Create(&rda).Error; err != nil {
		return 0, err
	}
	return rda.ID, nil
}

func (r *rda) FindRDA(tx *gorm.DB, param *model.FindRDAParam, orderBy *model.OrderBy, output interface{}) error {
	db := r.gorm.DB()
	if tx != nil {
		db = tx
	}
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 id 篩選條件
	if param.ID != nil {
		query += "AND id = ? "
		params = append(params, *param.ID)
	}
	//加入 user_id 篩選條件
	if param.UserID != nil {
		query += "AND user_id = ? "
		params = append(params, *param.UserID)
	}
	db = db.Model(&entity.RDA{})
	if orderBy != nil {
		db = db.Order(fmt.Sprintf("%s %s", orderBy.Field, orderBy.OrderType))
	}
	if err := db.Where(query, params...).Take(output).Error; err != nil {
		return err
	}
	return nil
}
