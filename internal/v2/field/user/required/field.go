package required

type IDField struct {
	ID int64 `json:"id" form:"id" uri:"id" gorm:"column:id" binding:"required" example:"10001"` // 帳戶id
}
type AccountTypeField struct {
	AccountType int `json:"account_type" form:"account_type" gorm:"column:account_type" binding:"required,oneof=1 2 3 4 5" example:"1"` // 帳號類型 (1:Email註冊/2:FB註冊/3:Google註冊/4:Line註冊/5:Apple註冊)
}
type AccountField struct {
	Account string `json:"account" form:"account" gorm:"column:account" binding:"required" example:"test@gmail.com"` // 帳號
}
type PasswordField struct {
	Password string `json:"password" form:"password" gorm:"column:password" binding:"required,min=6,max=18" example:"12345678"` // 密碼
}
type DeviceTokenField struct {
	DeviceToken string `json:"device_token" form:"device_token" gorm:"column:device_token;default:''" binding:"required" example:"d2we12ew3d12we1"` // 推播token
}
type UserStatusField struct {
	UserStatus int `json:"user_status" form:"user_status" gorm:"column:user_status;default:1" binding:"required,oneof=1 2 3" example:"1"` // 用戶狀態 (1:正常/2:違規/3:刪除)
}
type UserTypeField struct {
	UserType int `json:"user_type" form:"user_type" gorm:"column:user_type;default:1" binding:"required,oneof=1 2" example:"1"` // 用戶類型 (1:一般用戶/2:訂閱用戶)
}
type EmailField struct {
	Email string `json:"email" form:"email" gorm:"column:email;default:''" binding:"required,email,max=255" example:"test@gmail.com"` // 信箱
}
type NicknameField struct {
	Nickname string `json:"nickname" form:"nickname" gorm:"column:nickname;default:''" binding:"required,min=1,max=50" example:"henry"` // 暱稱
}
type AvatarField struct {
	Avatar string `json:"avatar" form:"avatar" gorm:"column:avatar;default:''" binding:"required" example:"123.png"` // 用戶大頭貼
}
type SexField struct {
	Sex string `json:"sex" form:"sex" gorm:"column:sex;default:''" binding:"required,oneof=m f" example:"m"` // 性別 (m:男/f:女)
}
type BirthdayField struct {
	Birthday string `json:"birthday" form:"birthday" gorm:"column:birthday" binding:"required,datetime=2006-01-02" example:"1991-02-27"` // 生日
}
type HeightField struct {
	Height float64 `json:"height" form:"height" gorm:"column:height;default:0" binding:"required,min=0.1,max=999.9" example:"165.5"` // 身高
}
type WeightField struct {
	Weight float64 `json:"weight" form:"weight" gorm:"column:weight;default:0" binding:"required,min=0.1,max=999.9" example:"50.5"` // 體重
}
type ExperienceField struct {
	Experience int `json:"experience" form:"experience" gorm:"column:experience;default:0" binding:"required,oneof=0 1 2 3 4" example:"1"` // 經驗 (0:未指定/1:初學/2:中級/3:中高/4:專業)
}
type TargetField struct {
	Target int `json:"target" form:"target" gorm:"column:target;default:0" binding:"required,oneof=0 1 2 3" example:"1"` // 目標 (0:未指定/1:減重/2:維持健康/3:增肌)
}
type IsDeletedField struct {
	IsDeleted int `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;default:0" binding:"required,oneof=0 1" example:"0"` //是否刪除
}
type CreateAtField struct {
	CreateAt string `json:"create_at" form:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" form:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-12 00:00:00"` // 更新時間
}
