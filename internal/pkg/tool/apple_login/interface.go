package apple_login

import "time"

type Tool interface {
	GenerateClientSecret(duration time.Duration) (string, error)
	GetUserID(authCode string, clientSecret string) (string, error)
}
