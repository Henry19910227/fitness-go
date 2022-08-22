package otp

import (
	"encoding/base32"
	otpSetting "github.com/Henry19910227/fitness-go/internal/pkg/setting/otp"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"time"
)

type tool struct {
	setting otpSetting.Setting
}

func New(setting otpSetting.Setting) Tool {
	return &tool{setting: setting}
}

// Generate 生成 OTP 碼
func (t *tool) Generate(secret string) (string, error) {
	newSecret := base32.StdEncoding.EncodeToString([]byte(secret))
	return totp.GenerateCodeCustom(newSecret, time.Now().UTC(), totp.ValidateOpts{
		Period:    uint(t.setting.GetPeriod()),
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
}

// Validate 驗證 OTP 碼
func (t *tool) Validate(code string, secret string) bool {
	newSecret := base32.StdEncoding.EncodeToString([]byte(secret))
	verify, _ := totp.ValidateCustom(code, newSecret, time.Now().UTC(), totp.ValidateOpts{
		Period:    uint(t.setting.GetPeriod()),
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	return verify
}
