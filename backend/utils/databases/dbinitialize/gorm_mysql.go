package dbinitialize

import (
	. "gitee.com/lybbn/golyadmin/config"
	"gitee.com/lybbn/golyadmin/utils/databases/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

// GormMysqlByConfig 初始化Mysql数据库传入配置
func GormMysqlByConfig(m Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		// sqlDB.SetConnMaxLifetime(time.Hour)   // SetConnMaxLifetime 设置了连接可复用的最大时间。
		return db
	}
}
