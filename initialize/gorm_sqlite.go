package initialize

import (
	"gitee.com/lybbn/go-vue-lyadmin/global"
	//"gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// GormSqlite 初始化Sqlite数据库
func GormSqlite() *gorm.DB {
	m := global.GVLA_CONFIG.Sqlite
	if m.Path == "" {
		return nil
	}

	if db, err := gorm.Open(sqlite.Open(m.Dsn()), &gorm.Config{}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns) // SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxOpenConns(m.MaxOpenConns) // SetMaxOpenConns 设置打开数据库连接的最大数量。
		// sqlDB.SetConnMaxLifetime(time.Hour)   // SetConnMaxLifetime 设置了连接可复用的最大时间。
		return db
	}
}
