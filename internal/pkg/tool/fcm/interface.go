package fcm

import (
	"time"
)

type Tool interface {
	GenerateGoogleOAuth2Token(duration time.Duration) (string, error)
	APIGetGooglePlayToken(oauthToken string) (string, error)
	APISendMessage(token string, message map[string]interface{}) error
}
