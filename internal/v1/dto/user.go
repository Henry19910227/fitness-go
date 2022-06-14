package dto

type UpdateUserParam struct {
	Nickname   *string  `gorm:"column:nickname"`
	Sex        *string  `gorm:"column:sex"`
	Birthday   *string  `gorm:"column:birthday"`
	Height     *float64 `gorm:"column:height"`
	Weight     *float64 `gorm:"column:weight"`
	Experience *int     `gorm:"column:experience"`
	Target     *int     `gorm:"column:target"`
	UserStatus *int     `gorm:"column:user_status"`
	Password   *string  `gorm:"column:password"`
}

type User struct {
	ID            int64              `json:"id" gorm:"column:id" example:"10001"`                               // 帳戶id
	AccountType   int                `json:"account_type" gorm:"column:account_type" example:"1"`               // 帳號類型 (1:Email註冊/2:FB註冊/3:Google註冊/4:Line註冊)
	Account       string             `json:"account" gorm:"column:account" example:"henry@gmail.com"`           // 帳號
	DeviceToken   string             `json:"device_token" gorm:"column:device_token" example:"f144b48d9695..."` // 推播 Token
	UserStatus    int                `json:"user_status" gorm:"column:user_status" example:"1"`                 // 用戶狀態 (1:正常/2:違規/3:刪除)
	UserType      int                `json:"user_type" gorm:"column:user_type" example:"1"`                     // 用戶狀態 (1:一般用戶/2:訂閱用戶)
	Email         string             `json:"email" gorm:"column:email" example:"henry@gmail.com"`               // 信箱
	Nickname      string             `json:"nickname" gorm:"column:nickname" example:"Henry"`                   // 暱稱
	Avatar        string             `json:"avatar" gorm:"column:avatar" example:"d2fe5w321a.png"`              // 用戶大頭貼
	Sex           string             `json:"sex" gorm:"column:sex" example:"m"`                                 // 性別 (m:男/f:女)
	Birthday      string             `json:"birthday" gorm:"column:birthday" example:"1991-02-27"`              // 生日
	Height        float64            `json:"height" gorm:"column:height" example:"176.5"`                       // 身高
	Weight        float64            `json:"weight" gorm:"column:weight" example:"72.5"`                        // 體重
	Experience    int                `json:"experience" gorm:"column:experience" example:"2"`                   // 經驗 (0:未指定/1:初學/2:中級/3:中高/4:專業)
	Target        int                `json:"target" gorm:"column:target" example:"3"`                           // 目標 (0:未指定/1:減重/2:維持健康/3:增肌)
	CreateAt      string             `json:"create_at" gorm:"column:create_at" example:"2021-06-01 12:00:00"`   // 創建日期
	UpdateAt      string             `json:"update_at" gorm:"column:update_at" example:"2021-06-01 12:00:00"`   // 修改日期
	TrainerInfo   *Trainer           `json:"trainer_info" gorm:"-"`                                             // 教練資訊
	SubscribeInfo *UserSubscribeInfo `json:"subscribe_info" gorm:"-"`                                           // 用戶訂閱資訊
}

type UserSummary struct {
	ID       int64  `json:"id" example:"1"`                     // 用戶id
	Nickname string `json:"nickname" example:"Henry"`           // 用戶暱稱
	Avatar   string `json:"avatar" example:"dkf2se51fsdds.png"` // 用戶大頭照
}

type UserAvatar struct {
	Avatar string `json:"avatar" example:"dkf2se51fsdds.png"` // 用戶大頭照
}

type FinsCMSUsersParam struct {
	UserID     *int64  // 用戶ID
	Name       *string // 用戶名稱(1~40字元)
	Email      *string // 用戶Email
	UserStatus *int    // 用戶狀態 (1:正常/2:違規/3:刪除)
	UserType   *int    // 用戶類型 (1:一般用戶/2:訂閱用戶)
}
