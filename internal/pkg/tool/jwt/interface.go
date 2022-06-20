package jwt

import "time"

type Tool interface {
	GenerateUserToken(uid int64) (string, error)
	GenerateAdminToken(uid int64, lv int) (string, error)
	VerifyToken(token string) error
	GetRoleByToken(token string) (int, error)
	GetIDByToken(token string) (int64, error)
	GetLvByToken(token string) (int64, error)
	GetExpire() time.Duration
}
