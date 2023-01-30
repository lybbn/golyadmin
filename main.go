package main

import (
	"fmt"
	"time"

	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/initialize"
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
	if global.GVLA_CONFIG.System.UseMultipoint || global.GVLA_CONFIG.System.UseRedis {
		initialize.InitRedis() // 初始化redis服务
	}
	global.GVLA_DB = initialize.Gorm() // gorm连接数据库
	if global.GVLA_DB != nil {
		global.GVLA_LOG.Info("database connect success")
	}
	r := initialize.Routers() //初始化路由
	address := fmt.Sprintf(":%d", global.GVLA_CONFIG.System.HttpPort)
	//启动
	// r.Run(address)
	s := initialize.InitServer(address, r)
	newaddress := address + global.GVLA_CONFIG.System.RouterPrefix
	time.Sleep(10 * time.Microsecond)
	global.GVLA_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
	欢迎使用 go-vue-lyadmin
	当前版本:v1.0.0
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认接口基础地址:http://127.0.0.1%s
`, newaddress, newaddress)
	global.GVLA_LOG.Error(s.ListenAndServe().Error())
}
