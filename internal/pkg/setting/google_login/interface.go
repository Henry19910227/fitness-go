package google_login

type Setting interface {
	GetAndroidClientID() string
	GetIOSClientID() string
	GetIss() string
	GetDebugTokenURL() string
}
