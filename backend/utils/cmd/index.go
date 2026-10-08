package cmd

import (
	"errors"
	"fmt"
	"os"

	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/initialize"
	"gitee.com/lybbn/golyadmin/utils"
	"gitee.com/lybbn/golyadmin/utils/cmd/auth/password"
	"gitee.com/lybbn/golyadmin/utils/cmd/auth/superadmin"
	"gitee.com/lybbn/golyadmin/utils/cmd/initdb"
	"gitee.com/lybbn/golyadmin/utils/cmd/migrate"
	"gitee.com/lybbn/golyadmin/utils/cmd/start"
	"gitee.com/lybbn/golyadmin/utils/cmd/version"
	"gitee.com/lybbn/golyadmin/utils/core"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var configYaml string
var rootCmd = &cobra.Command{
	Use:          "golyadmin",
	Short:        "golyadmin is a golyadmin manage commond",
	SilenceUsage: true,
	Long:         `golyadmin is a golyadmin manage commond`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(utils.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func initEnvironment() {
	// 配置文件选择优先级: 命令行 -c > 环境变量 GOLYADMIN_ENV=prod > 默认 config.yaml（开发）
	// （cobra flag 解析先于 OnInitialize 执行，此处直接读 Changed 可判断 -c 是否显式传入）
	if configYaml == "" {
		configYaml = "config.yaml"
	}
	explicit := false
	if f := rootCmd.PersistentFlags().Lookup("config"); f != nil {
		explicit = f.Changed
	}
	if !explicit && os.Getenv("GOLYADMIN_ENV") == "prod" {
		configYaml = "config.pro.yaml"
		fmt.Printf("检测到 GOLYADMIN_ENV=prod,使用生产配置文件 %s\n", configYaml)
	}
	global.GL_VP = core.Viper(configYaml) // 初始化Viper
	global.GL_LOG = core.Zap()            // 初始化zap日志库
	zap.ReplaceGlobals(global.GL_LOG)
	initialize.DBInit() // gorm连接数据库
	if global.GL_DB != nil {
		global.GL_LOG.Info("database connect success")
		fmt.Printf("数据库连接成功\n")
	}
}

func tip() {
	usageStr := `主命令 ` + utils.Green(`golyadmin `) + ` 可以使用 ` + utils.Red(`-h`) + ` 查看帮助`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	cobra.OnInitialize(initEnvironment)
	rootCmd.PersistentFlags().StringVarP(&configYaml, "config", "c", "config.yaml", "provided configuration file")
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(initdb.StartCmd)
	rootCmd.AddCommand(start.StartCmd)
	rootCmd.AddCommand(superadmin.StartCmd)
	rootCmd.AddCommand(password.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
