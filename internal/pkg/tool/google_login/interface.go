package google_login

type Tool interface {
	GetGoogleUidByAccessToken(accessToken string) (string, error)
}