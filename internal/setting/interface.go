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

type Redis interface {
	GetHost() string
	GetPwd() string
}

type User interface {

}

type Logger interface {
	GetLogFilePath() string
	GetLogFileName() string
	GetLogFileExt() string
	GetLogMaxAge() time.Duration
	GetLogRotationTime() time.Duration
	GetRunMode() string
}

type Uploader interface {
	GetUploadSavePath() string
	GetUploadImageAllowExts() []string
	GetUploadImageMaxSize() int
}

type UploadLimit interface {
	ImageAllowExts() []string
	ImageMaxSize() int
}