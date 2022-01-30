package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
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
		UserID: param.UserID,
		SubscribePlanID: param.SubscribePlanID,
		Status: int(param.Status),
		StartDate: param.StartDate,
		ExpiresDate: param.ExpiresDate,
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"subscribe_plan_id", "status", "start_date", "expires_date", "update_at"}),
	}).Create(&member).Error; err != nil {
		return 0, err
	}
	return member.UserID, nil
}

func (m *userSubscribeInfo) FindSubscribeInfo(uid int64) (*model.UserSubscribeInfo, error) {
	var member model.UserSubscribeInfo
	if err := m.gorm.DB().Find(&member, "user_id = ?", uid).Error; err != nil {
		return nil, err
	}
	return &member, nil
}