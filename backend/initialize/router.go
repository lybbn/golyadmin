package initialize

import (
	"net/http"

	docs "gitee.com/lybbn/go-vue-lyadmin/docs"
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/router"
	"gitee.com/lybbn/go-vue-lyadmin/utils/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化总路由
func Routers() *gin.Engine {
	//设置运行模式
	gin.SetMode(global.GVLA_CONFIG.System.RunMode)
	//创建路由
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System
	// 跨域
	if global.GVLA_CONFIG.System.IsCors {
		// Router.Use(middleware.CorsAllowAll())        // 直接放行全部跨域请求
		Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
		global.GVLA_LOG.Info("use middleware cors")
	}
	//错误异常捕获
	Router.Use(middleware.GinRecovery(false))
	docs.SwaggerInfo.BasePath = global.GVLA_CONFIG.System.RouterPrefix
	Router.GET(global.GVLA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	global.GVLA_LOG.Info("register swagger handler")
	//不需要认证
	PublicGroup := Router.Group(global.GVLA_CONFIG.System.RouterPrefix)
	{
		systemRouter.InitBaseRouter(PublicGroup)
	}
	//需要认证
	PrivateGroup := Router.Group(global.GVLA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuthMiddleware())
	{
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitOperationLogRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
	}
	//不存在路由
	Router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "404 not found",
		})
	})

	global.GVLA_LOG.Info("router register success")
	return Router
}
