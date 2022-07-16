package fb_login

type Tool interface {
	GetFbUidByAccessToken(accessToken string) (string, error)
}