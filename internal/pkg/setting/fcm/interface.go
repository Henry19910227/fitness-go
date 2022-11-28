package fcm

type Setting interface {
	GetURL() string
	GetTokenURL() string
	GetScope() string
	GetProjectID() string
	GetKeyName() string
}
