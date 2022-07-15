package jwt

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var (
	UserTokenPrefix  = "fitness.user.token"
	AdminTokenPrefix = "fitness.admin.token"
)

type tool struct {
	setting setting.JWT
}

func New(setting setting.JWT) Tool {
	return &tool{setting}
}

func (t *tool) GenerateUserToken(uid int64) (string, error) {
	claims := jwt.MapClaims{"uid": strconv.Itoa(int(uid)), "role": "1", "time": time.Now()}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.setting.GetTokenSecret()))
	return token, err
}

func (t *tool) GenerateAdminToken(uid int64, lv int) (string, error) {
	claims := jwt.MapClaims{"uid": strconv.Itoa(int(uid)), "role": "2", "lv": strconv.Itoa(lv)}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.setting.GetTokenSecret()))
	return token, err
}

func (t *tool) VerifyToken(tokenString string) error {
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

func (t *tool) GetIDByToken(tokenString string) (int64, error) {
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

func (t *tool) GetRoleByToken(token string) (int, error) {
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

func (t *tool) GetLvByToken(tokenString string) (int64, error) {
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

func (t *tool) GetExpire() time.Duration {
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
