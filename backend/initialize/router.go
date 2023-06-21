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
	gin.SetMode(global.GL_CONFIG.System.RunMode)
	//创建路由
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System

	// 集成部署web端（不使用nginx）
	// 前端执行打包命令 npm run build。把打包后的dist目录放入backend，然后在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/static")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	// 采用https访问，使用SSL证书
	if global.GL_CONFIG.Ssl.Enable {
		Router.Use(middleware.LoadSSL())
	}

	// 跨域
	if global.GL_CONFIG.System.IsCors {
		// Router.Use(middleware.CorsAllowAll())        // 直接放行全部跨域请求
		Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
		global.GL_LOG.Info("use middleware cors")
	}

	//错误异常捕获
	Router.Use(middleware.GinRecovery(false))

	//swagger文档（正式环境可注释掉）
	if global.GL_CONFIG.System.IsSwagger {
		docs.SwaggerInfo.BasePath = global.GL_CONFIG.System.RouterPrefix
		Router.GET(global.GL_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		global.GL_LOG.Info("register swagger handler")
	}

	//不需要认证
	PublicGroup := Router.Group(global.GL_CONFIG.System.RouterPrefix)
	{
		systemRouter.InitBaseRouter(PublicGroup)
	}

	//需要认证
	PrivateGroup := Router.Group(global.GL_CONFIG.System.RouterPrefix)
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

	global.GL_LOG.Info("router register success")
	return Router
}
