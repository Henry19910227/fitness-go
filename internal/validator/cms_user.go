package validator

type CMSGetUsersQuery struct {
	UserID     *int64  `form:"user_id" binding:"omitempty" example:"10001"`              // 用戶ID
	Name       *string `form:"name" binding:"omitempty,min=1,max=40" example:"Henry"`    // 用戶名稱(1~40字元)
	Email      *string `form:"email" binding:"omitempty,email" example:"test@gmail.com"` // 用戶Email
	UserStatus *int    `form:"user_status" binding:"omitempty,oneof=1 2 3" example:"1"`  // 用戶狀態 (1:正常/2:違規/3:刪除)
	UserType   *int    `form:"user_type" binding:"omitempty,oneof=1 2" example:"1"`      // 用戶類型 (1:一般用戶/2:訂閱用戶)
}
