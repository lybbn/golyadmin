package start

import (
	"fmt"
	"time"

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
	address := fmt.Sprintf("%s:%d", global.GL_CONFIG.System.Host, global.GL_CONFIG.System.HttpPort)
	newaddress := utils.GetLocalIpAddr() + utils.GetServerPort() + global.GL_CONFIG.System.RouterPrefix
	time.Sleep(10 * time.Microsecond)
	global.GL_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Println("\r\n欢迎使用 golyadmin")
	fmt.Printf("当前版本:v%s\r\n", global.GL_VERSION)
	tip()
	fmt.Println(utils.Green("Server run at:"))
	fmt.Printf("-  Local:   http://127.0.0.1%s \r\n", utils.GetServerPort())
	fmt.Printf("-  Network: http://%s%s \r\n", utils.GetLocalIpAddr(), utils.GetServerPort())
	if global.GL_CONFIG.System.IsSwagger {
		fmt.Printf("默认接口文档地址:http://%s/swagger/index.html \r\n", newaddress)
	}
	fmt.Printf("默认接口基础地址:http://%s \r\n", newaddress)
	if global.GL_CONFIG.Ssl.Enable {
		if err := r.RunTLS(address, global.GL_CONFIG.Ssl.CertFile, global.GL_CONFIG.Ssl.KeyFile); err != nil {
			global.GL_LOG.Error(err.Error())
		}
	} else {
		if err := r.Run(address); err != nil {
			global.GL_LOG.Error(err.Error())
		}
	}
}

func tip() {
	usageStr := `主命令 ` + utils.Green(`golyadmin `) + ` 可以使用 ` + utils.Red(`-h`) + ` 查看帮助`
	fmt.Printf("%s\n", usageStr)
}
