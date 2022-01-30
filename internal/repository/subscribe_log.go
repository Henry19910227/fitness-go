package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type subscribeLog struct {
	gorm tool.Gorm
}

func NewSubscribeLog(gorm tool.Gorm) SubscribeLog {
	return &subscribeLog{gorm: gorm}
}

func (s *subscribeLog) CreateSubscribeLog(tx *gorm.DB, param *model.CreateSubscribeLogParam) (int64, error) {
	db := s.gorm.DB()
	if tx != nil {
		db = tx
	}
	log := entity.SubscribeLog{
		OriginalTransactionID: param.OriginalTransactionID,
		TransactionID: param.TransactionID,
		PurchaseDate:  param.PurchaseDate,
		ExpiresDate:   param.ExpiresDate,
		Type:          param.Type,
		Msg:           param.Msg,
		CreateAt:      time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Create(&log).Error; err != nil {
		return 0, err
	}
	return log.ID, nil
}
