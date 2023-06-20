package captcha

import (
	"context"
	"time"

	"gitee.com/lybbn/golyadmin/global"

	"go.uber.org/zap"
)

var ctx = context.Background()

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180, //180s
		PreKey:     "LYADMIN_CAPTCHA_",
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
}

func (rs *RedisStore) Set(id string, value string) error {
	key := rs.PreKey + id
	err := global.GVLA_REDIS.Set(ctx, key, value, rs.Expiration).Err()
	if err != nil {
		global.GVLA_LOG.Error("RedisStoreSetError!", zap.Error(err))
	}
	return err
}

func (rs *RedisStore) Get(id string, clear bool) string {
	key := rs.PreKey + id
	val, err := global.GVLA_REDIS.Get(ctx, key).Result()
	if err != nil {
		global.GVLA_LOG.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := global.GVLA_REDIS.Del(ctx, key).Err()
		if err != nil {
			global.GVLA_LOG.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
