package apple_login

type Setting interface {
	GetKeyName() string
	GetBundleID() string
	GetDebugTokenURL() string
	GetTeamID() string
	GetKeyID() string
}
