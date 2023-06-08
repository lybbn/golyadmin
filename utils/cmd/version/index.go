package version

import (
	"fmt"

	"gitee.com/lybbn/go-vue-lyadmin/global"
	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: "golyadmin version",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func run() {
	fmt.Println(global.GVLA_VERSION)
}
