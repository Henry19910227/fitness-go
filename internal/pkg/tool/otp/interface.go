package otp

import (
	"encoding/base32"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"time"
)

type Tool interface {
	Generate(secret string) (string, error)
	Validate(code string, secret string) bool
}

type tool struct {
}

func New() Tool {
	return &tool{}
}

// Generate 生成 OTP 碼
func (t *tool) Generate(secret string) (string, error) {
	newSecret := base32.StdEncoding.EncodeToString([]byte(secret))
	return totp.GenerateCodeCustom(newSecret, time.Now().UTC(), totp.ValidateOpts{
		Period:    300,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
}

// Validate 驗證 OTP 碼
func (t *tool) Validate(code string, secret string) bool {
	newSecret := base32.StdEncoding.EncodeToString([]byte(secret))
	verify, _ := totp.ValidateCustom(code, newSecret, time.Now().UTC(), totp.ValidateOpts{
		Period:    300,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	return verify
}
