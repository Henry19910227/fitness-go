package feedback

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //主鍵id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` //用戶id
}
type VersionField struct {
	Version *string `json:"version,omitempty" gorm:"column:version;default:''" example:"1.0.0"` //軟體版本
}
type PlatformField struct {
	Platform *string `json:"platform,omitempty" gorm:"column:platform;default:''" example:"ios"` //平台(ios/android)
}
type OSVersionField struct {
	OSVersion *string `json:"os_version,omitempty" gorm:"column:os_version;default:''" example:"14.0"` //OS版本
}
type PhoneModelField struct {
	PhoneModel *string `json:"phone_model,omitempty" gorm:"column:phone_model;default:''" example:"iPhoneX"` //手機型號
}
type BodyField struct {
	Body *string `json:"body,omitempty" gorm:"column:body" example:"我遇到了bug!!"` //反饋內文
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	IDField
	UserIDField
	VersionField
	PlatformField
	OSVersionField
	PhoneModelField
	BodyField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "feedbacks"
}
