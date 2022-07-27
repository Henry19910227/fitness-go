package line_login

type Tool interface {
	GetUserIDByAccessToken(accessToken string) (string, error)
}
