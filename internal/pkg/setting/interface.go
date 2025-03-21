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

type Resource interface {
	RootPath() string
}

type Upload interface {
	ImageAllowExts() []string
	AudioAllowExts() []string
	VideoAllowExts() []string
	ImageMaxSize() int
	AudioMaxSize() int
	VideoMaxSize() int
}

type IAP interface {
	GetSandboxURL() string
	GetProductURL() string
	GetAppServerAPIURL() string
	GetPassword() string
	GetKeyPath() string
	GetKeyID() string
	GetBundleID() string
	GetIssuer() string
}

type IAB interface {
	GetURL() string
	GetScope() string
	GetJsonFilePath() string
	GetPackageName() string
}
