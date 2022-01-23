package model

import "github.com/Henry19910227/fitness-go/internal/global"

type UserSubscribeInfo struct {
	UserID int64 `gorm:"column:user_id"` // 用戶id
	Status int `gorm:"column:status"` // 會員狀態(0:無會員狀態/1:付費會員狀態)
	StartDate string `gorm:"column:start_date"` // 訂閱開始日期
	ExpiresDate string `gorm:"column:expires_date"` // 訂閱過期日期
	UpdateAt string `gorm:"column:update_at"` // 更新時間
}

func (UserSubscribeInfo) TableName() string {
	return "user_subscribe_infos"
}

type SaveUserSubscribeInfoParam struct {
	UserID int64 // 用戶id
	Status global.SubscribeStatus // 會員狀態(1:正常/2:到期/3:取消)
	StartDate string `gorm:"column:start_date"` // 訂閱開始日期
	ExpiresDate string // 訂閱過期日期
}