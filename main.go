package main

import (
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/initialize"
	"gitee.com/lybbn/go-vue-lyadmin/utils/cmd"
	"gitee.com/lybbn/go-vue-lyadmin/utils/core"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GVLA_VP = core.Viper() // 初始化Viper
	global.GVLA_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVLA_LOG)
	initialize.DBInit() // gorm连接数据库
	if global.GVLA_DB != nil {
		global.GVLA_LOG.Info("database connect success")
	}
	cmd.Execute()
}
