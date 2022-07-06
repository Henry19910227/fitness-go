package user

type IDRequired struct {
	ID int64 `json:"id" gorm:"column:id"` // 帳戶id
}
type AccountTypeRequired struct {
	AccountType int `json:"account_type" gorm:"column:account_type;default:1"` // 帳號類型 (1:Email註冊/2:FB註冊/3:Google註冊/4:Line註冊)
}
type AccountRequired struct {
	Account string `json:"account" gorm:"column:account;default:''"` // 帳號
}
type PasswordRequired struct {
	Password string `json:"password" gorm:"column:password;default:''" binding:"required,min=6,max=18" example:"12345678"` // 密碼
}
type DeviceTokenRequired struct {
	DeviceToken string `json:"device_token" gorm:"column:device_token;default:''"` // 推播token
}
type UserStatusRequired struct {
	UserStatus int `json:"user_status" gorm:"column:user_status;default:1"` // 用戶狀態 (1:正常/2:違規/3:刪除)
}
type UserTypeRequired struct {
	UserType int `json:"user_type" gorm:"column:user_type;default:1"` // 用戶類型 (1:一般用戶/2:訂閱用戶)
}
type EmailRequired struct {
	Email string `json:"email" gorm:"column:email;default:''"` // 信箱
}
type NicknameRequired struct {
	Nickname string `json:"nickname" gorm:"column:nickname;default:''"` // 暱稱
}
type AvatarRequired struct {
	Avatar string `json:"avatar" gorm:"column:avatar;default:''"` // 用戶大頭貼
}
type SexRequired struct {
	Sex string `json:"sex" gorm:"column:sex;default:m"` // 性別 (m:男/f:女)
}
type BirthdayRequired struct {
	Birthday string `json:"birthday" gorm:"column:birthday;default:1991-02-27"` // 生日
}
type HeightRequired struct {
	Height float64 `json:"height" gorm:"column:height;default:176"` // 身高
}
type WeightRequired struct {
	Weight float64 `json:"weight" gorm:"column:weight;default:70"` // 體重
}
type ExperienceRequired struct {
	Experience int `json:"experience" gorm:"column:experience;default:0"` // 經驗 (0:未指定/1:初學/2:中級/3:中高/4:專業)
}
type TargetRequired struct {
	Target int `json:"target" gorm:"column:target;default:0"` // 目標 (0:未指定/1:減重/2:維持健康/3:增肌)
}
type CreateAtRequired struct {
	CreateAt string `json:"create_at" gorm:"column:create_at;default:2022-06-12 00:00:00" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtRequired struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at;default:2022-06-12 00:00:00" example:"2022-06-12 00:00:00"` // 更新時間
}
