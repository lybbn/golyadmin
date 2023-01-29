package main

import (
	"fmt"
	"time"
	"go.uber.org/zap"
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/initialize"
	"gitee.com/lybbn/go-vue-lyadmin/utils/core"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	// 初始化Viper
	global.GVLA_VP = core.Viper() 
	// 初始化zap日志库
	global.GVLA_LOG = core.Zap()
	zap.ReplaceGlobals(global.GVLA_LOG)
	if global.GVLA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.InitRedis()
	}
	//初始化路由
	r := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVLA_CONFIG.System.HttpPort)
	fmt.Println(global.GVLA_CONFIG.System.UseRedis)
	//启动
	// r.Run(address)
	s := initialize.InitServer(address, r)
	time.Sleep(10 * time.Microsecond)
	global.GVLA_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
	欢迎使用 go-vue-lyadmin
	当前版本:v1.0.0
	默认接口地址:http://127.0.0.1%s
`, address)
	global.GVLA_LOG.Error(s.ListenAndServe().Error())
}
