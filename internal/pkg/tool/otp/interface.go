package otp

type Tool interface {
	Generate(secret string) (string, error)
	Validate(code string, secret string) bool
}
