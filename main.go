package main

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/initialize"
	"go.uber.org/zap"
)

func main() {
	//设置运行模式
	gin.SetMode(global.GVLA_CONFIG.System.RunMode)
	if global.GVLA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
	//初始化路由
	initialize.Routers()
	//run
	r.Run(":"+string(global.GVLA_CONFIG.System.HttpPort))
}
