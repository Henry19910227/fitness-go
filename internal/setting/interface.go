package setting

type Mysql interface {
	GetUserName() string
	GetPassword() string
	GetHost() string
	GetDatabase() string
}

type Migrate interface {
	DirPathSource() string
}
