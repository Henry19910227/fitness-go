package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type purchaseLog struct {
	gorm  tool.Gorm
}

func NewPurchaseLog(gorm tool.Gorm) PurchaseLog {
	return &purchaseLog{gorm: gorm}
}

func (p *purchaseLog) CreatePurchaseLog(tx *gorm.DB, param *model.CreatePurchaseLogParam) (int64, error) {
	db := p.gorm.DB()
	if tx != nil {
		db = tx
	}
	log := entity.PurchaseLog{
		UserID: param.UserID,
		OrderID: param.OrderID,
		Type: int(param.Type),
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Create(&log).Error; err != nil {
		return 0, err
	}
	return log.ID, nil
}
