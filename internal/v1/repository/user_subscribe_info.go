package repository

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type userSubscribeInfo struct {
	gorm tool.Gorm
}

func NewSubscribeInfo(gorm tool.Gorm) UserSubscribeInfo {
	return &userSubscribeInfo{gorm: gorm}
}

func (m *userSubscribeInfo) SaveSubscribeInfo(tx *gorm.DB, param *model.SaveUserSubscribeInfoParam) (int64, error) {
	if param == nil {
		return 0, nil
	}
	db := m.gorm.DB()
	if tx != nil {
		db = tx
	}
	member := entity.UserSubscribeInfo{
		UserID:      param.UserID,
		OrderID:     param.OrderID,
		Status:      int(param.Status),
		StartDate:   param.StartDate,
		ExpiresDate: param.ExpiresDate,
		UpdateAt:    time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"order_id", "status", "start_date", "expires_date", "update_at"}),
	}).Create(&member).Error; err != nil {
		return 0, err
	}
	return member.UserID, nil
}

func (m *userSubscribeInfo) FindSubscribeInfo(uid int64) (*model.UserSubscribeInfo, error) {
	var member model.UserSubscribeInfo
	if err := m.gorm.DB().Take(&member, "user_id = ?", uid).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (m *userSubscribeInfo) FindSubscribeInfoByOriginalTransactionID(originalTransactionID string) (*model.UserSubscribeInfo, error) {
	var info model.UserSubscribeInfo
	if err := m.gorm.DB().
		Joins("INNER JOIN orders ON user_subscribe_infos.order_id = orders.id").
		Joins("INNER JOIN receipts ON receipts.order_id = orders.id").
		Where("receipts.original_transaction_id = ?", originalTransactionID).
		Take(&info).Error; err != nil {
		return nil, err
	}
	return &info, nil
}
