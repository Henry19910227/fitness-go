package tool

import (
	"context"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type redisTool struct {
	client *redis.Client
}

func NewRedis(setting setting.Redis) Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     setting.GetHost(),
		Password: setting.GetPwd(), // no password set
		DB:       0,  // use default DB
	})
	return &redisTool{client}
}

func (t *redisTool) Get(key string) (string, error) {
	return t.client.Get(ctx, key).Result()
}

func (t *redisTool) Del(key string) error {
	return t.client.Del(ctx, key).Err()
}

func (t *redisTool) SetEX(key string, value interface{}, expiration time.Duration) error {
	return t.client.SetEX(ctx, key, value, expiration).Err()
}

func (t *redisTool) XRange(key string, start string, end string, count *int64) ([]redis.XMessage, error) {
	if count != nil {
		return t.client.XRangeN(ctx, key, start, end, *count).Result()
	}
	return t.client.XRange(ctx, key, start, end).Result()
}


func (t *redisTool) LRange(listName string, start int, stop int) ([]string, error) {
	return t.client.LRange(ctx, listName, int64(start), int64(stop)).Result()
}

func (t *redisTool) LLEN(listName string) (int64, error) {
	return t.client.LLen(ctx, listName).Result()
}

func (t *redisTool) Keys(patten string) ([]string, error) {
	return t.client.Keys(ctx, patten).Result()
}

func (t *redisTool) NewPipeliner() redis.Pipeliner {
	return t.client.Pipeline()
}

func (t *redisTool) PipLRange(pip redis.Pipeliner, listName string, start int, stop int) *redis.StringSliceCmd {
	return pip.LRange(ctx, listName, int64(start), int64(stop))
}

func (t *redisTool) PipXRange(pip redis.Pipeliner, key string, start string, end string, count *int64) *redis.XMessageSliceCmd {
	if count != nil {
		return pip.XRangeN(ctx, key, start, end, *count)
	}
	return pip.XRange(ctx, key, start, end)
}

func (t *redisTool) PipXRevRange(pip redis.Pipeliner, key string, start string, end string, count *int64) *redis.XMessageSliceCmd {
	if count != nil {
		return pip.XRevRangeN(ctx, key, start, end, *count)
	}
	return pip.XRevRange(ctx, key, start, end)
}

func (t *redisTool) PipXLen(pip redis.Pipeliner, key string) *redis.IntCmd {
	return pip.XLen(ctx, key)
}

func (t *redisTool) PipExec(pip redis.Pipeliner) error {
	_, err := pip.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
