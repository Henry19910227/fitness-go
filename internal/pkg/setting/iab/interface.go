package iab

type Setting interface {
	GetURL() string
	GetTokenURL() string
	GetScope() string
	GetPackageName() string
	GetKeyName() string
}
