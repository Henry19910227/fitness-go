package redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type Tool interface {
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
