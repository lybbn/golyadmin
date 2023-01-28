package global

import (
	"gitee.com/lybbn/go-vue-lyadmin/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var (
	GVLA_REDIS  *redis.Client //Redis客户端
	GVLA_CONFIG config.Server //服务端配置
	GVLA_LOG    *zap.Logger
)
