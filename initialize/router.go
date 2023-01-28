package initialize

import (
	"net/http"
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/router"
	"github.com/gin-gonic/gin"
)

// 初始化总路由

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	exampleRouter := router.RouterGroupApp.Example

	InstallPlugin(Router) // 安装插件

	global.GVLA_LOG.Info("router register success")
	return Router
}