package jwt

import "time"

type Setting interface {
	GetTokenSecret() string
	GetIssuer() string
	GetExpire() time.Duration
}
