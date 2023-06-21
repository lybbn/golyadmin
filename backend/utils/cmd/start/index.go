package start

import (
	"fmt"

	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/initialize"
	"gitee.com/lybbn/golyadmin/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	StartCmd = &cobra.Command{
		Use:     "start",
		Short:   "start backend service",
		Example: "golyadmin start",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func run() {
	fmt.Println("正在启动后端服务")
	initialize.InitRedis()    // 初始化redis服务
	r := initialize.Routers() //初始化路由
	assginAddr := global.GL_CONFIG.System.Host
	address := fmt.Sprintf("%s:%d", assginAddr, global.GL_CONFIG.System.HttpPort)
	newaddress := assginAddr + utils.GetServerPort()
	if assginAddr == "0.0.0.0" {
		newaddress = utils.GetLocalIpAddr() + utils.GetServerPort()
	}
	// 与优雅的重启停止服务一起配合使用
	// srv := &http.Server{
	// 	Addr:    address,
	// 	Handler: r,
	// }
	global.GL_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Println("\r\n欢迎使用 golyadmin")
	fmt.Printf("当前版本:v%s\r\n", global.GL_VERSION)
	tip()
	fmt.Println(utils.Green("Server run at:"))
	fmt.Printf("-  Local:   http://127.0.0.1%s \r\n", utils.GetServerPort())
	fmt.Printf("-  Network: http://%s \r\n", newaddress)
	if global.GL_CONFIG.System.IsSwagger {
		fmt.Printf("默认接口文档地址:http://%s/swagger/index.html \r\n", newaddress+global.GL_CONFIG.System.RouterPrefix)
	}
	fmt.Printf("默认接口基础地址:http://%s \r\n", newaddress+global.GL_CONFIG.System.RouterPrefix)
	// 启动服务（使用优雅重启停止服务需要注释掉以下9行）
	if global.GL_CONFIG.Ssl.Enable {
		if err := r.RunTLS(address, global.GL_CONFIG.Ssl.CertFile, global.GL_CONFIG.Ssl.KeyFile); err != nil {
			global.GL_LOG.Error(err.Error())
		}
	} else {
		if err := r.Run(address); err != nil {
			global.GL_LOG.Error(err.Error())
		}
	}
	// 启动服务 优雅重启停止服务（要使用请打开以下所有注释）。global.GL_CONFIG.Ssl.KeyFile 中的文件路径要改为证书key的内容
	// go func() {
	// 	if global.GL_CONFIG.Ssl.Enable {
	// 		if err := srv.ListenAndServeTLS(global.GL_CONFIG.Ssl.CertFile, global.GL_CONFIG.Ssl.KeyFile); err != nil && err != http.ErrServerClosed {
	// 			global.GL_LOG.Error("listen:", zap.Error(err))
	// 		}
	// 	} else {
	// 		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 			global.GL_LOG.Error("listen:", zap.Error(err))
	// 		}
	// 	}
	// }()
	// // 优雅重启停止服务等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	// <-quit

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// global.GL_LOG.Info("Shutdown Server ...")
	// fmt.Printf("%s Shutdown Server ... \r\n", utils.GetNowTimeFormatStr())

	// if err := srv.Shutdown(ctx); err != nil {
	// 	global.GL_LOG.Error("Server Shutdown:", zap.Error(err))
	// }
	// global.GL_LOG.Info("Server exiting")
	// fmt.Printf("Server exiting \r\n")
}

func tip() {
	usageStr := `主命令 ` + utils.Green(`golyadmin `) + ` 可以使用 ` + utils.Red(`-h`) + ` 查看帮助`
	fmt.Printf("%s\n", usageStr)
}
