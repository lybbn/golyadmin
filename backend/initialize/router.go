package initialize

import (
	"net/http"

	docs "gitee.com/lybbn/golyadmin/docs"
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/router"
	"gitee.com/lybbn/golyadmin/utils/middleware"
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
	// 集成部署web端（不使用nginx）
	// 前端执行打包命令 npm run build。把打包后的dist目录放入backend，然后在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/static")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面
	// 采用https访问，使用SSL证书。如需要使用https 请打开此中间件注释 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	//Router.Use(middleware.LoadSSL())
	// 跨域
	if global.GVLA_CONFIG.System.IsCors {
		// Router.Use(middleware.CorsAllowAll())        // 直接放行全部跨域请求
		Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
		global.GVLA_LOG.Info("use middleware cors")
	}
	//错误异常捕获
	Router.Use(middleware.GinRecovery(false))
	//swagger文档（正式环境可注释掉）
	if global.GVLA_CONFIG.System.IsSwagger {
		docs.SwaggerInfo.BasePath = global.GVLA_CONFIG.System.RouterPrefix
		Router.GET(global.GVLA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		global.GVLA_LOG.Info("register swagger handler")
	}
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
