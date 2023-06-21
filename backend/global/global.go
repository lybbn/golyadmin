package global

import (
	"sync"

	"gitee.com/lybbn/golyadmin/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

const (
	// Version golyadmin version info
	GL_VERSION = "1.0.1"
)

var (
	GL_DB           *gorm.DB            //默认数据库
	GL_DATABASES    map[string]*gorm.DB //多数据源,访问使用GL_DATABASES[数据库别名]形式指定数据源
	GL_REDIS        *redis.Client       //Redis客户端
	GL_CONFIG       config.Server       //服务端配置
	GL_LOG          *zap.Logger
	GL_VP           *viper.Viper
	GL_Singleflight = &singleflight.Group{} //处理并发，合并相同请求（防缓存击穿）
	lock            sync.RWMutex
)

// 通过数据库别名获取GL_DATABASES中的db
func GetGlobalDBByName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GL_DATABASES[dbname]
	if !ok || db == nil {
		panic("the db no does not init")
	}
	return db
}
