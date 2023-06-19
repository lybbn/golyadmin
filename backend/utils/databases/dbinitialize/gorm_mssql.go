package dbinitialize

import (
	"gitee.com/lybbn/go-vue-lyadmin/config"
	"gitee.com/lybbn/go-vue-lyadmin/utils/databases/internal"

	//"gorm.io/driver/sqlserver"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mssql struct {
	config.GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// dsn := "sqlserver://gorm:root@localhost:1433?database=gorm"
func (m *Mssql) Dsn() string {
	return "sqlserver://" + m.Username + ":" + m.Password + "@" + m.Path + ":" + m.Port + "?database=" + m.Dbname + "&encrypt=disable"
}

// GormMssqlByConfig 初始化Mssql数据库传入配置
// 如果需要sqlserver库 放开import里的注释 把下方 mysql.Config 改为 sqlserver.Config ;  mysql.New 改为 sqlserver.New
func GormMssqlByConfig(m Mssql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mssqlConfig := mysql.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	if db, err := gorm.Open(mysql.New(mssqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
