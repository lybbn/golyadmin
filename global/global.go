package global

import (
	"gitee.com/lybbn/go-vue-lyadmin/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVLA_DB     *gorm.DB
	GVLA_REDIS  *redis.Client //Redis客户端
	GVLA_CONFIG config.Server //服务端配置
	GVLA_LOG    *zap.Logger
	GVLA_VP     *viper.Viper
)
