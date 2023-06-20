package dbinitialize

import (
	. "gitee.com/lybbn/golyadmin/config"

	//"gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Sqlite) Dsn() string {
	return m.Path
}

// GormSqliteByConfig 初始化Sqlite数据库传入配置
func GormSqliteByConfig(s Sqlite) *gorm.DB {
	if s.Path == "" {
		return nil
	}

	if db, err := gorm.Open(sqlite.Open(s.Dsn()), &gorm.Config{}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns) // SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxOpenConns(s.MaxOpenConns) // SetMaxOpenConns 设置打开数据库连接的最大数量。
		// sqlDB.SetConnMaxLifetime(time.Hour)   // SetConnMaxLifetime 设置了连接可复用的最大时间。
		return db
	}
}
