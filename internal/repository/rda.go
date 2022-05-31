package repository

import (
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

func (r *rda) CreateRDA(tx *gorm.DB, userID int64, param *model.CreateRDAParam) error {
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
		return err
	}
	return nil
}
