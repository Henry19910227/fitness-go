package registerdto

type OTP struct {
	Code string `json:"otp_code" example:"254235"` // 信箱驗證碼
}

type Register struct {
	UserID int64 `json:"user_id" example:"10001"` // 用戶ID
}