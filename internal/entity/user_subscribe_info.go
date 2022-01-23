package entity

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
