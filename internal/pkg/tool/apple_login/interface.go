package apple_login

import "time"

type Tool interface {
	GenerateClientSecret(duration time.Duration) (string, error)
	APIGetRefreshToken(authCode string, clientSecret string) (string, error)
	APIGetUserID(refreshToken string, clientSecret string) (string, error)
}
