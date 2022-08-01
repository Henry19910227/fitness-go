package fb_login

type Tool interface {
	GetUserID(authCode string, clientSecret string) (string, error)
}
