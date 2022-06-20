package redis

type Setting interface {
	GetHost() string
	GetPwd() string
}
