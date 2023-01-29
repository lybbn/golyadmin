package initialize

import (
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/router"
	"gitee.com/lybbn/go-vue-lyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
)

// 初始化总路由
func Routers() *gin.Engine {
	//设置运行模式
	gin.SetMode(global.GVLA_CONFIG.System.RunMode)
	//创建路由
	Router := gin.Default()
	exampleRouter := router.RouterGroupApp.Example
	// 跨域
	if global.GVLA_CONFIG.System.IsCors {
		Router.Use(middleware.Cors()) // 直接放行全部跨域请求
		Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
		global.GVLA_LOG.Info("use middleware cors")
	}

	PrivateGroup := Router.Group("")
	{
		exampleRouter.InitExampleRouter(PrivateGroup)
	}

	global.GVLA_LOG.Info("router register success")
	return Router
}