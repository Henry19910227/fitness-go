package user

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"10001"` // 帳戶id
}
type AccountTypeField struct {
	AccountType *int `json:"account_type,omitempty" gorm:"column:account_type" example:"1"` // 帳號類型 (1:Email註冊/2:FB註冊/3:Google註冊/4:Line註冊/5:Apple註冊)
}
type AccountField struct {
	Account *string `json:"account,omitempty" gorm:"column:account" example:"test@gmail.com"` // 帳號
}
type PasswordField struct {
	Password *string `json:"password,omitempty" gorm:"column:password" example:"12345678"` // 密碼
}
type DeviceTokenField struct {
	DeviceToken *string `json:"device_token,omitempty" gorm:"column:device_token;default:''" example:"d2we12ew3d12we1"` // 推播token
}
type UserStatusField struct {
	UserStatus *int `json:"user_status,omitempty" gorm:"column:user_status;default:1" example:"1"` // 用戶狀態 (1:正常/2:違規/3:刪除)
}
type UserTypeField struct {
	UserType *int `json:"user_type,omitempty" gorm:"column:user_type;default:1" example:"1"` // 用戶類型 (1:一般用戶/2:訂閱用戶)
}
type EmailField struct {
	Email *string `json:"email,omitempty" gorm:"column:email;default:''" example:"test@gmail.com"` // 信箱
}
type NicknameField struct {
	Nickname *string `json:"nickname,omitempty" gorm:"column:nickname;default:''" example:"henry"` // 暱稱
}
type AvatarField struct {
	Avatar *string `json:"avatar,omitempty" gorm:"column:avatar;default:''" example:"123.png"` // 用戶大頭貼
}
type SexField struct {
	Sex *string `json:"sex,omitempty" gorm:"column:sex;default:''" example:"m"` // 性別 (m:男/f:女)
}
type BirthdayField struct {
	Birthday *string `json:"birthday,omitempty" gorm:"column:birthday" example:"1991-02-27"` // 生日
}
type HeightField struct {
	Height *float64 `json:"height,omitempty" gorm:"column:height;default:0" example:"165.5"` // 身高
}
type WeightField struct {
	Weight *float64 `json:"weight,omitempty" gorm:"column:weight;default:0" example:"50.5"` // 體重
}
type ExperienceField struct {
	Experience *int `json:"experience,omitempty" gorm:"column:experience;default:0" example:"1"` // 經驗 (0:未指定/1:初學/2:中級/3:中高/4:專業)
}
type TargetField struct {
	Target *int `json:"target,omitempty" gorm:"column:target;default:0" example:"1"` // 目標 (0:未指定/1:減重/2:維持健康/3:增肌)
}
type IsDeletedField struct {
	IsDeleted *int `json:"is_deleted,omitempty" gorm:"column:is_deleted;default:0" example:"0"` //是否刪除
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-12 00:00:00"` // 更新時間
}

type Table struct {
	IDField
	AccountTypeField
	AccountField
	PasswordField
	DeviceTokenField
	UserStatusField
	UserTypeField
	EmailField
	NicknameField
	AvatarField
	SexField
	BirthdayField
	HeightField
	WeightField
	ExperienceField
	TargetField
	IsDeletedField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "users"
}
