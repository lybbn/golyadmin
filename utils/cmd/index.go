package cmd

import (
	"errors"
	"fmt"
	"os"

	"gitee.com/lybbn/go-vue-lyadmin/utils"
	"gitee.com/lybbn/go-vue-lyadmin/utils/cmd/migrate"
	"gitee.com/lybbn/go-vue-lyadmin/utils/cmd/start"
	"gitee.com/lybbn/go-vue-lyadmin/utils/cmd/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "golyadmin",
	Short:        "golyadmin is a go-vue-lyadmin manage commond",
	SilenceUsage: true,
	Long:         `golyadmin is a go-vue-lyadmin manage commond`,
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

func tip() {
	usageStr := `主命令 ` + utils.Green(`golyadmin `) + ` 可以使用 ` + utils.Red(`-h`) + ` 查看帮助`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(start.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
