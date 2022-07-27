package fb_login

type Tool interface {
	GetUserIDByAccessToken(accessToken string) (string, error)
}
