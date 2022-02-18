package dto

type UserSubscribeInfo struct {
	Status      int    `json:"status" gorm:"column:status"`             // 會員狀態(0:無會員狀態/1:付費會員狀態)
	StartDate   string `json:"start_date" gorm:"column:start_date"`     // 訂閱開始日期
	ExpiresDate string `json:"expires_date" gorm:"column:expires_date"` // 訂閱過期日期
}
