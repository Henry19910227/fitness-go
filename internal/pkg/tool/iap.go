package tool

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"io/ioutil"
	"time"
)

type iap struct {
	setting setting.IAP
}

func NewIAP(setting setting.IAP) IAP {
	return &iap{setting}
}

func (i *iap) SandboxURL() string {
	return i.setting.GetSandboxURL()
}

func (i *iap) ProductURL() string {
	return i.setting.GetProductURL()
}

func (i *iap) AppServerAPIURL() string {
	return i.setting.GetAppServerAPIURL()
}

func (i *iap) Password() string {
	return i.setting.GetPassword()
}

func (i *iap) GenerateAppleStoreAPIToken(duration time.Duration) (string, error) {
	p8bytes, err := ioutil.ReadFile(i.setting.GetKeyPath())
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode(p8bytes)
	if block == nil || block.Type != "PRIVATE KEY" {
		return "", errors.New("generate apple token error")
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	ecdsaPrivateKey, ok := parsedKey.(*ecdsa.PrivateKey)
	if !ok {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": i.setting.GetIssuer(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(duration).Unix(),
		"aud": "appstoreconnect-v1",
		"nonce": uuid.New().String(),
		"bid": i.setting.GetBundleID(),
	})
	token.Header["kid"] = i.setting.GetKeyID()
	tokenString, err := token.SignedString(ecdsaPrivateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}