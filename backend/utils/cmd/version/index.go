package version

import (
	"fmt"

	"gitee.com/lybbn/golyadmin/global"
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
	fmt.Println(global.GL_VERSION)
}
