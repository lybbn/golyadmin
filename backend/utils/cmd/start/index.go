package start

import (
	"fmt"
	"time"

	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/initialize"
	"gitee.com/lybbn/go-vue-lyadmin/utils"
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
	address := fmt.Sprintf("%s:%d", global.GVLA_CONFIG.System.Host, global.GVLA_CONFIG.System.HttpPort)
	s := initialize.InitServer(address, r)
	newaddress := utils.GetLocalIpAddr() + utils.GetServerPort() + global.GVLA_CONFIG.System.RouterPrefix
	time.Sleep(10 * time.Microsecond)
	global.GVLA_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Println("\r\n欢迎使用 go-vue-lyadmin")
	fmt.Printf("当前版本:v%s\r\n", global.GVLA_VERSION)
	tip()
	fmt.Println(utils.Green("Server run at:"))
	fmt.Printf("-  Local:   http://127.0.0.1%s \r\n", utils.GetServerPort())
	fmt.Printf("-  Network: http://%s%s \r\n", utils.GetLocalIpAddr(), utils.GetServerPort())
	if global.GVLA_CONFIG.System.IsSwagger {
		fmt.Printf("默认接口文档地址:http://%s/swagger/index.html \r\n", newaddress)
	}
	fmt.Printf("默认接口基础地址:http://%s \r\n", newaddress)
	global.GVLA_LOG.Error(s.ListenAndServe().Error())
	//gin启动
	// r.Run(address)
}

func tip() {
	usageStr := `主命令 ` + utils.Green(`golyadmin `) + ` 可以使用 ` + utils.Red(`-h`) + ` 查看帮助`
	fmt.Printf("%s\n", usageStr)
}
