package initialize

import (
	//"github.com/dzwvip/oracle"
	"gitee.com/lybbn/go-vue-lyadmin/config"
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/initialize/internal"

	//_ "github.com/godror/godror"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormOracle 初始化oracle数据库
// 如果需要Oracle库 放开import里的注释 把下方 mysql.Config 改为 oracle.Config ;  mysql.New 改为 oracle.New
// 运行go run main.go会报错：github.com/godror/godror cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%
// 解决方法：下载x86_64-posix-seh（对应版本就行） https://sourceforge.net/projects/mingw-w64/files/mingw-w64/ 并添加到系统环境变量中
func GormOracle() *gorm.DB {
	m := global.GVLA_CONFIG.Oracle
	if m.Dbname == "" {
		return nil
	}
	oracleConfig := mysql.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	if db, err := gorm.Open(mysql.New(oracleConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormOracleByConfig 初始化Oracle数据库用过传入配置
func GormOracleByConfig(m config.Oracle) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	oracleConfig := mysql.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度

	}
	if db, err := gorm.Open(mysql.New(oracleConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
