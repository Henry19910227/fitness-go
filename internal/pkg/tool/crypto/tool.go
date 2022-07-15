package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

type tool struct {
}

func New() Tool {
	return &tool{}
}

func (t *tool) MD5Encode(plaintext string) string {
	h := md5.New()
	h.Write([]byte(plaintext))
	return hex.EncodeToString(h.Sum(nil))
}
