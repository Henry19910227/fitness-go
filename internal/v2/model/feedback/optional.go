package feedback

type IDOptional struct {
	ID *int64 `json:"id,omitempty" binding:"omitempty" example:"1"` //主鍵id
}
type UserIDOptional struct {
	UserID *int64 `json:"user_id,omitempty" binding:"omitempty" example:"10001"` //用戶id
}
type VersionOptional struct {
	Version *string `json:"version,omitempty" form:"version" binding:"omitempty,max=50" example:"1.0.0"` //軟體版本
}
type PlatformOptional struct {
	Platform *string `json:"platform,omitempty" form:"platform" binding:"omitempty,oneof=ios android" example:"ios"` //平台(ios/android)
}
type OSVersionOptional struct {
	OSVersion *string `json:"os_version,omitempty" form:"os_version" binding:"omitempty,max=50" example:"14.0"` //OS版本
}
type PhoneModelOptional struct {
	PhoneModel *string `json:"phone_model,omitempty" form:"phone_model" binding:"omitempty,max=50" example:"iPhoneX"` //手機型號
}
type BodyOptional struct {
	Body *string `json:"body,omitempty" form:"body" binding:"omitempty,max=300" example:"我遇到了bug!!"` //反饋內文
}
