package entity

type User struct {
	ID          int64   `gorm:"column:id"`           // 帳戶id
	AccountType int     `gorm:"column:account_type"` // 帳號類型 (1:Email註冊/2:FB註冊/3:Google註冊/4:Line註冊)
	Account     string  `gorm:"column:account"`      // 帳號
	Password    string  `gorm:"column:password"`     // 密碼
	DeviceToken string  `gorm:"column:device_token"` // 推播token
	UserStatus  int     `gorm:"column:user_status"`  // 用戶狀態 (1:正常/2:違規/3:刪除)
	UserType    int     `gorm:"column:user_type"`    // 用戶類型 (1:一般用戶/2:訂閱用戶)
	CreateAt    string  `gorm:"column:create_at"`    // 創建日期
	UpdateAt    string  `gorm:"column:update_at"`    // 修改日期
	Email       string  `gorm:"column:email"`        // 信箱
	Nickname    string  `gorm:"column:nickname"`     // 暱稱
	Avatar      string  `gorm:"column:avatar"`       // 用戶大頭貼
	Sex         string  `gorm:"column:sex"`          // 性別 (m:男/f:女)
	Birthday    string  `gorm:"column:birthday"`     // 生日
	Height      float64 `gorm:"column:height"`       // 身高
	Weight      float64 `gorm:"column:weight"`       // 體重
	Experience  int     `gorm:"column:experience"`   // 經驗 (0:未指定/1:初學/2:中級/3:中高/4:專業)
	Target      int     `gorm:"column:target"`       // 目標 (0:未指定/1:減重/2:維持健康/3:增肌)
}

func (User) TableName() string {
	return "users"
}

type UserTemplate struct {
	ID          int64            `gorm:"column:id"`                        // 帳戶id
	AccountType int              `gorm:"column:account_type"`              // 帳號類型 (1:Email註冊/2:FB註冊/3:Google註冊/4:Line註冊)
	Account     string           `gorm:"column:account"`                   // 帳號
	Password    string           `gorm:"column:password"`                  // 密碼
	DeviceToken string           `gorm:"column:device_token"`              // 推播token
	UserStatus  int              `gorm:"column:user_status"`               // 用戶狀態 (1:正常/2:違規/3:刪除)
	UserType    int              `gorm:"column:user_type"`                 // 用戶類型 (1:一般用戶/2:訂閱用戶)
	CreateAt    string           `gorm:"column:create_at"`                 // 創建日期
	UpdateAt    string           `gorm:"column:update_at"`                 // 修改日期
	Email       string           `gorm:"column:email"`                     // 信箱
	Nickname    string           `gorm:"column:nickname"`                  // 暱稱
	Avatar      string           `gorm:"column:avatar"`                    // 用戶大頭貼
	Sex         string           `gorm:"column:sex"`                       // 性別 (m:男/f:女)
	Birthday    string           `gorm:"column:birthday"`                  // 生日
	Height      float64          `gorm:"column:height"`                    // 身高
	Weight      float64          `gorm:"column:weight"`                    // 體重
	Experience  int              `gorm:"column:experience"`                // 經驗 (0:未指定/1:初學/2:中級/3:中高/4:專業)
	Target      int              `gorm:"column:target"`                    // 目標 (0:未指定/1:減重/2:維持健康/3:增肌)
	Trainer     *TrainerTemplate `gorm:"foreignkey:user_id;references:id"` // 教練身份
}

func (UserTemplate) TableName() string {
	return "users"
}
