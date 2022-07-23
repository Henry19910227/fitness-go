package google_login

type Tool interface {
	GetGoogleUidByIDToken(accessToken string) (string, error)
}