package tool

import (
	"database/sql"
	"time"
)

type Mysql interface {
	DB() *sql.DB
}

type Migrate interface {
	Up(step *int) error
	Down(step *int) error
	Force(version int) error
	Migrate(version uint) error
	Version() (uint, bool, error)
}

type JWT interface {
	GenerateUserToken(uid int64) (string, error)
	GenerateAdminToken(uid int64, lv int) (string, error)
	VerifyToken(token string) error
	GetRoleByToken(token string) (int, error)
	GetIDByToken(token string) (int64, error)
	GetLvByToken(token string) (int64, error)
	GetExpire() time.Duration
}