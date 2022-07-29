package iap

type Setting interface {
	GetSandboxURL() string
	GetProductURL() string
	GetAppServerAPIURL() string
	GetPassword() string
	GetKeyPath() string
	GetKeyName() string
	GetKeyID() string
	GetBundleID() string
	GetIssuer() string
}
