package initialize

import (
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/router"
	"github.com/gin-gonic/gin"
)

// 初始化总路由
func Routers() *gin.Engine {
	//设置运行模式
	gin.SetMode(global.GVLA_CONFIG.System.RunMode)
	//创建路由
	Router := gin.Default()
	exampleRouter := router.RouterGroupApp.Example
	PrivateGroup := Router.Group("")
	{
		exampleRouter.InitExampleRouter(PrivateGroup)
	}

	// global.GVLA_LOG.Info("router register success")
	return Router
}