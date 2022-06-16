package mysql

type Setting interface {
	GetUserName() string
	GetPassword() string
	GetHost() string
	GetDatabase() string
}
