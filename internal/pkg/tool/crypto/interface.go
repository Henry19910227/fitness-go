package crypto

type Tool interface {
	MD5Encode(plaintext string) string
}
