package dto

type CMSTrainerSummary struct {
	UserID        int64  `json:"user_id" gorm:"column:user_id" example:"1001"`                    // 用戶id
	Nickname      string `json:"nickname" gorm:"column:nickname" example:"Henry"`                 // 教練暱稱
	TrainerStatus int    `json:"trainer_status" gorm:"column:trainer_status" example:"1"`         // 教練帳戶狀態 (1:正常/2:審核中/3:停權)
	Email         string `json:"email" gorm:"column:email" example:"henry@gmail.com"`             // 信箱
	CreateAt      string `json:"create_at" gorm:"column:create_at" example:"2021-06-01 12:00:00"` // 創建日期
	UpdateAt      string `json:"update_at" gorm:"column:update_at" example:"2021-06-01 12:00:00"` // 修改日期
}

func (CMSTrainerSummary) TableName() string {
	return "trainers"
}

type FinsCMSTrainersParam struct {
	UserID        *int64  // 用戶ID
	NickName      *string // 教練暱稱(1~40字元)
	Email         *string // 用戶Email
	TrainerStatus *int    // 教練狀態 教練狀態(1:正常/2:審核中/3:停權/4:未啟用)
}
