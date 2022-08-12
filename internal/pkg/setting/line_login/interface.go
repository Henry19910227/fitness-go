package line_login

type Setting interface {
	GetVerifyTokenURL() string
	GetProfileURL() string
	GetClientID() string
}
