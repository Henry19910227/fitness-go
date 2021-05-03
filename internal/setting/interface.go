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

type Swagger interface {
	GetVersion() string
	GetProtocol() string
	GetHost() string
	GetBasePath() string
}
