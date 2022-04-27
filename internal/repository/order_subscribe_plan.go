package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderSubscribePlan struct {
	gorm tool.Gorm
}

func NewOrderSubscribePlan(gorm tool.Gorm) OrderSubscribePlan {
	return &orderSubscribePlan{gorm: gorm}
}

func (o *orderSubscribePlan) SaveOrderSubscribePlan(tx *gorm.DB, orderID string, subscribePlanID int64) error {
	db := o.gorm.DB()
	if tx != nil {
		db = tx
	}
	plan := entity.OrderSubscribePlan{
		OrderID: orderID,
		SubscribePlanID: subscribePlanID,
	}
	if err := db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "order_id"}},
		DoUpdates: clause.AssignmentColumns(
			[]string{"subscribe_plan_id"}),
	}).Create(&plan).Error; err != nil {
		return err
	}
	return nil
}



