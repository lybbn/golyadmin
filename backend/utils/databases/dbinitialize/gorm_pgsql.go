package dbinitialize

import (
	"gitee.com/lybbn/golyadmin/config"
	"gitee.com/lybbn/golyadmin/utils/databases/internal"

	//"gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Pgsql struct {
	config.GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// Dsn 基于配置文件获取 dsn
func (p *Pgsql) Dsn() string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port + " " + p.Config
}

// LinkDsn 根据 dbname 生成 dsn
func (p *Pgsql) LinkDsn(dbname string) string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + dbname + " port=" + p.Port + " " + p.Config
}

// GormPgSqlByConfig 初始化 Postgresql 数据库传入配置
// 如果需要sqlserver库 放开import里的注释 把下方 mysql.Config 改为 postgres.Config ;  mysql.New 改为 postgres.New
// 并把PreferSimpleProtocol注释放开
func GormPgSqlByConfig(p Pgsql) *gorm.DB {
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := mysql.Config{
		DSN: p.Dsn(), // DSN data source name
		//PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(mysql.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}
