package global

import (
	"sync"

	"gitee.com/lybbn/go-vue-lyadmin/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	// Version go-vue-lyadmin version info
	GVLA_VERSION = "1.0.1"
)

var (
	GVLA_DB        *gorm.DB            //默认数据库
	GVLA_DATABASES map[string]*gorm.DB //多数据源
	GVLA_REDIS     *redis.Client       //Redis客户端
	GVLA_CONFIG    config.Server       //服务端配置
	GVLA_LOG       *zap.Logger
	GVLA_VP        *viper.Viper
	lock           sync.RWMutex
)

// 通过数据库别名获取GVLA_DATABASES中的db
func GetGlobalDBByName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVLA_DATABASES[dbname]
	if !ok || db == nil {
		panic("the db no does not init")
	}
	return db
}
