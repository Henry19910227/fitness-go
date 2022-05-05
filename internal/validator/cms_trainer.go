package validator

type CMSGetTrainersQuery struct {
	UserID        *int64  `form:"user_id" binding:"omitempty" example:"10001"`                  // 用戶ID
	Nickname      *string `form:"nickname" binding:"omitempty,min=1,max=40" example:"Henry"`    // 教練名稱(1~40字元)
	Email         *string `form:"email" binding:"omitempty,email" example:"test@gmail.com"`     // 教練Email
	TrainerStatus *int    `form:"trainer_status" binding:"omitempty,oneof=1 2 3 4" example:"1"` // 教練狀態(1:正常/2:審核中/3:停權/4:未啟用)
}
