package feedback

type IDRequired struct {
	ID int64 `json:"id" binding:"required" example:"1"` //主鍵id
}
type UserIDRequired struct {
	UserID int64 `json:"user_id" binding:"required" example:"10001"` //用戶id
}
type VersionRequired struct {
	Version string `json:"version" binding:"required" example:"1.0.0"` //軟體版本
}
type PlatformRequired struct {
	Platform string `json:"platform" binding:"required" example:"ios"` //平台(ios/android)
}
type OSVersionRequired struct {
	OSVersion string `json:"os_version" binding:"required" example:"14.0"` //OS版本
}
type PhoneModelRequired struct {
	PhoneModel string `json:"phone_model" binding:"required" example:"iPhoneX"` //手機型號
}
type BodyRequired struct {
	Body string `json:"body" form:"body" binding:"required,min=1,max=300" example:"我遇到了bug!!"` //反饋內文
}
