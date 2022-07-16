package fb_login

type Setting interface {
	GetAppID() string
	GetAppSecret() string
	GetDebugTokenURL() string
}
