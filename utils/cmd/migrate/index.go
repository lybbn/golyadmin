package migrate

import (
	"fmt"

	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/model/system"
	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:     "migrate",
		Short:   "Initialize the database",
		Example: "golyadmin migrate",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func run() {
	initDB()
}

func migrateModel() error {
	err := global.GVLA_DB.Debug().AutoMigrate(&system.SysAdminUsers{})
	if err != nil {
		return err
	}
	return err
}

func initDB() {
	//4. 数据库迁移
	fmt.Println("数据库迁移开始")
	_ = migrateModel()
	fmt.Println(`数据库迁移成功`)
}
