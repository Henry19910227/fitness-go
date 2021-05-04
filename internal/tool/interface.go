package tool

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
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

type Redis interface {
	Get(key string) (string, error)
	SetEX(key string, value interface{}, expiration time.Duration) error
	Del(key string) error
	XRange(key string, start string, end string, count *int64) ([]redis.XMessage, error)
	LRange(listName string, start int, stop int) ([]string, error)
	LLEN(listName string) (int64, error)
	Keys(patten string) ([]string, error)
	NewPipeliner() redis.Pipeliner
	PipLRange(pip redis.Pipeliner, listName string, start int, stop int) *redis.StringSliceCmd
	PipXRange(pip redis.Pipeliner, key string, start string, end string, count *int64) *redis.XMessageSliceCmd
	PipXRevRange(pip redis.Pipeliner, key string, start string, end string, count *int64) *redis.XMessageSliceCmd
	PipXLen(pip redis.Pipeliner, key string) *redis.IntCmd

	PipExec(pip redis.Pipeliner) error
}