package line_login

type Tool interface {
	GetUserID(authCode string) (string, error)
}
