package google_login

type Setting interface {
	GetClientID() string
	GetIss() string
	GetDebugTokenURL() string
}
