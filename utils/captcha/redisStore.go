package captcha

import (
	"context"
	"time"

	"gitee.com/lybbn/go-vue-lyadmin/global"
	"github.com/mojocn/base64Captcha"

	"go.uber.org/zap"
)

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * global.GVLA_CONFIG.Captcha.CaptchaTimeout,
		PreKey:     "LYADMIN_CAPTCHA_",
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *RedisStore) Set(id string, value string) {
	err := global.GVLA_REDIS.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		global.GVLA_LOG.Error("RedisStoreSetError!", zap.Error(err))
	}
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.GVLA_REDIS.Get(rs.Context, key).Result()
	if err != nil {
		global.GVLA_LOG.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := global.GVLA_REDIS.Del(rs.Context, key).Err()
		if err != nil {
			global.GVLA_LOG.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) VerifyCaptcha(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
