package iab

import "time"

type Setting interface {
	GetURL() string
	GetTokenURL() string
	GetScope() string
	GetExpire() time.Duration
	GetPackageName() string
	GetKeyName() string
}
