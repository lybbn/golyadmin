package initialize

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/utils/databases/dbinitialize"
	"gorm.io/gorm"
)

const defaultDBName = "default"

// Gorm 初始化数据库并产生数据库全局变量
func DBInit() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.GL_CONFIG.Databases {
		switch info.DbType {
		case "mysql":
			dbMap[info.AliasName] = dbinitialize.GormMysqlByConfig(dbinitialize.Mysql{GeneralDB: info})
		case "mssql":
			dbMap[info.AliasName] = dbinitialize.GormMssqlByConfig(dbinitialize.Mssql{GeneralDB: info})
		case "pgsql":
			dbMap[info.AliasName] = dbinitialize.GormPgSqlByConfig(dbinitialize.Pgsql{GeneralDB: info})
		case "oracle":
			dbMap[info.AliasName] = dbinitialize.GormOracleByConfig(dbinitialize.Oracle{GeneralDB: info})
		case "sqlite":
			dbMap[info.AliasName] = dbinitialize.GormSqliteByConfig(dbinitialize.Sqlite{GeneralDB: info})
		default:
			continue
		}
	}
	//默认数据库
	if dfDB, ok := dbMap[defaultDBName]; ok {
		global.GL_DB = dfDB
	}
	global.GL_DATABASES = dbMap
}
