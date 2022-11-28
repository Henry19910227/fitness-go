package fcm

import "time"

type Setting interface {
	GetURL() string
	GetTokenURL() string
	GetScope() string
	GetExpire() time.Duration
	GetProjectID() string
	GetKeyName() string
}
