package redis

import (
	"context"
	redisSetting "github.com/Henry19910227/fitness-go/internal/pkg/setting/redis"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type tool struct {
	client *redis.Client
}

func New(setting redisSetting.Setting) Tool {
	client := redis.NewClient(&redis.Options{
		Addr:     setting.GetHost(),
		Password: setting.GetPwd(), // no password set
		DB:       0,                // use default DB
	})
	return &tool{client}
}

func (t *tool) Get(key string) (string, error) {
	return t.client.Get(ctx, key).Result()
}

func (t *tool) Del(key string) error {
	return t.client.Del(ctx, key).Err()
}

func (t *tool) SetEX(key string, value interface{}, expiration time.Duration) error {
	return t.client.SetEX(ctx, key, value, expiration).Err()
}

func (t *tool) XRange(key string, start string, end string, count *int64) ([]redis.XMessage, error) {
	if count != nil {
		return t.client.XRangeN(ctx, key, start, end, *count).Result()
	}
	return t.client.XRange(ctx, key, start, end).Result()
}

func (t *tool) LRange(listName string, start int, stop int) ([]string, error) {
	return t.client.LRange(ctx, listName, int64(start), int64(stop)).Result()
}

func (t *tool) LLEN(listName string) (int64, error) {
	return t.client.LLen(ctx, listName).Result()
}

func (t *tool) Keys(patten string) ([]string, error) {
	return t.client.Keys(ctx, patten).Result()
}

func (t *tool) NewPipeliner() redis.Pipeliner {
	return t.client.Pipeline()
}

func (t *tool) PipLRange(pip redis.Pipeliner, listName string, start int, stop int) *redis.StringSliceCmd {
	return pip.LRange(ctx, listName, int64(start), int64(stop))
}

func (t *tool) PipXRange(pip redis.Pipeliner, key string, start string, end string, count *int64) *redis.XMessageSliceCmd {
	if count != nil {
		return pip.XRangeN(ctx, key, start, end, *count)
	}
	return pip.XRange(ctx, key, start, end)
}

func (t *tool) PipXRevRange(pip redis.Pipeliner, key string, start string, end string, count *int64) *redis.XMessageSliceCmd {
	if count != nil {
		return pip.XRevRangeN(ctx, key, start, end, *count)
	}
	return pip.XRevRange(ctx, key, start, end)
}

func (t *tool) PipXLen(pip redis.Pipeliner, key string) *redis.IntCmd {
	return pip.XLen(ctx, key)
}

func (t *tool) PipExec(pip redis.Pipeliner) error {
	_, err := pip.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
