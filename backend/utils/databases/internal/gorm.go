package internal

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/schema"

	"gitee.com/lybbn/golyadmin/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	defaultLog := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond, //慢sql的标准为 200毫秒
		LogLevel:      logger.Warn,
		Colorful:      true, //彩色输出
	})

	switch global.GL_CONFIG.System.GormLogMode {
	case "silent", "Silent":
		config.Logger = defaultLog.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = defaultLog.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = defaultLog.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = defaultLog.LogMode(logger.Info)
	default:
		config.Logger = defaultLog.LogMode(logger.Info)
	}
	return config
}
