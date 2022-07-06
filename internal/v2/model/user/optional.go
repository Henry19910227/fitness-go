package user

type IDOptional struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id"` // 帳戶id
}
type AccountTypeOptional struct {
	AccountType *int `json:"account_type,omitempty" gorm:"column:account_type;default:1"` // 帳號類型 (1:Email註冊/2:FB註冊/3:Google註冊/4:Line註冊)
}
type AccountOptional struct {
	Account *string `json:"account,omitempty" gorm:"column:account;default:''"` // 帳號
}
type PasswordOptional struct {
	Password *string `json:"password,omitempty" gorm:"column:password;default:''"` // 密碼
}
type DeviceTokenOptional struct {
	DeviceToken *string `json:"device_token,omitempty" gorm:"column:device_token;default:''"` // 推播token
}
type UserStatusOptional struct {
	UserStatus *int `json:"user_status,omitempty" gorm:"column:user_status;default:1"` // 用戶狀態 (1:正常/2:違規/3:刪除)
}
type UserTypeOptional struct {
	UserType *int `json:"user_type,omitempty" gorm:"column:user_type;default:1"` // 用戶類型 (1:一般用戶/2:訂閱用戶)
}
type EmailOptional struct {
	Email *string `json:"email,omitempty" gorm:"column:email;default:''"` // 信箱
}
type NicknameOptional struct {
	Nickname *string `json:"nickname,omitempty" gorm:"column:nickname;default:''"` // 暱稱
}
type AvatarOptional struct {
	Avatar *string `json:"avatar,omitempty" gorm:"column:avatar;default:''"` // 用戶大頭貼
}
type SexOptional struct {
	Sex *string `json:"sex,omitempty" gorm:"column:sex;default:m"` // 性別 (m:男/f:女)
}
type BirthdayOptional struct {
	Birthday *string `json:"birthday,omitempty" gorm:"column:birthday;default:1991-02-27"` // 生日
}
type HeightOptional struct {
	Height *float64 `json:"height,omitempty" gorm:"column:height;default:176"` // 身高
}
type WeightOptional struct {
	Weight *float64 `json:"weight,omitempty" gorm:"column:weight;default:70"` // 體重
}
type ExperienceOptional struct {
	Experience *int `json:"experience,omitempty" gorm:"column:experience;default:0"` // 經驗 (0:未指定/1:初學/2:中級/3:中高/4:專業)
}
type TargetOptional struct {
	Target *int `json:"target,omitempty" gorm:"column:target;default:0"` // 目標 (0:未指定/1:減重/2:維持健康/3:增肌)
}
type CreateAtOptional struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at;default:2022-06-12 00:00:00" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtOptional struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at;default:2022-06-12 00:00:00" example:"2022-06-12 00:00:00"` // 更新時間
}
