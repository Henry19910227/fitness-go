package setting

import "time"

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

type JWT interface {
	GetTokenSecret() string
	GetIssuer() string
	GetExpire() time.Duration
}