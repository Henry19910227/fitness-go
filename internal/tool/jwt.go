package tool

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"github.com/google/uuid"
	"io/ioutil"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type jwtTool struct {
	setting setting.JWT
}

func NewJWT(setting setting.JWT) JWT {
	return &jwtTool{setting}
}

func (t *jwtTool) GenerateUserToken(uid int64) (string, error) {
	claims := jwt.MapClaims{"uid": strconv.Itoa(int(uid)), "role": "1", "time": time.Now()}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.setting.GetTokenSecret()))
	return token, err
}

func (t *jwtTool) GenerateTrainerToken(uid int64) (string, error) {
	claims := jwt.MapClaims{"uid": strconv.Itoa(int(uid)), "role": "2"}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.setting.GetTokenSecret()))
	return token, err
}

func (t *jwtTool) GenerateAdminToken(uid int64, lv int) (string, error) {
	claims := jwt.MapClaims{"uid": strconv.Itoa(int(uid)), "lv": strconv.Itoa(lv)}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.setting.GetTokenSecret()))
	return token, err
}

func (t *jwtTool) GenerateAppleStoreServerAPIToken(keyPath string, iss string, bid string, kid string) (string, error) {
	p8bytes, err := ioutil.ReadFile("./config/SubscriptionKey_Q2N9VXCN4F.p8")
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
		"iss": "69a6de92-1fe6-47e3-e053-5b8c7c11a4d1",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
		"aud": "appstoreconnect-v1",
		"nonce": uuid.New().String(),
		"bid": "com.henry.PurchaseDemo"})
	token.Header["kid"] = "Q2N9VXCN4F"
	tokenString, err := token.SignedString(ecdsaPrivateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *jwtTool) VerifyToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.setting.GetTokenSecret()), nil
	})
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			return errors.New("Timeout")
		default:
			return err
		}
	}
	return nil
}

func (t *jwtTool) GetIDByToken(tokenString string) (int64, error) {
	claims, err := getJWTClaims(tokenString, t.setting.GetTokenSecret())
	if err != nil {
		return 0, err
	}
	uidStr, ok := claims["uid"].(string)
	if !ok {
		return 0, errors.New("assertion error")
	}
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return 0, errors.New("strconv error")
	}
	return int64(uid), nil
}

func (t *jwtTool) GetRoleByToken(token string) (int, error) {
	claims, err := getJWTClaims(token, t.setting.GetTokenSecret())
	if err != nil {
		return 0, err
	}
	roleStr, ok := claims["role"].(string)
	if !ok {
		return 0, errors.New("assertion error")
	}
	role, err := strconv.Atoi(roleStr)
	if err != nil {
		return 0, errors.New("strconv error")
	}
	return role, nil
}

func (t *jwtTool) GetLvByToken(tokenString string) (int64, error) {
	claims, err := getJWTClaims(tokenString, t.setting.GetTokenSecret())
	if err != nil {
		return 0, err
	}
	lvStr, ok := claims["lv"].(string)
	if !ok {
		return 0, errors.New("assertion error")
	}
	lv, err := strconv.Atoi(lvStr)
	if err != nil {
		return 0, errors.New("strconv error")
	}
	return int64(lv), nil
}

func (t *jwtTool) GetExpire() time.Duration {
	return t.setting.GetExpire()
}

func getJWTClaims(tokenString string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("assertion error")
	}
	return claims, nil
}