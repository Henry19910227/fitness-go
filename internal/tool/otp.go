package tool

import (
	"encoding/base32"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type otpTool struct {
}

func NewOTP() OTP {
	return &otpTool{}
}

// Generate 生成 OTP 碼
func (tool *otpTool) Generate(secret string) (string, error) {
	newSecret := base32.StdEncoding.EncodeToString([]byte(secret))
	return totp.GenerateCodeCustom(newSecret, time.Now().UTC(), totp.ValidateOpts{
		Period:    60,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
}

// Validate 驗證 OTP 碼
func (tool *otpTool) Validate(code string, secret string) bool {
	newSecret := base32.StdEncoding.EncodeToString([]byte(secret))
	verify, _ := totp.ValidateCustom(code, newSecret, time.Now().UTC(), totp.ValidateOpts{
		Period:    60,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	return verify
}
