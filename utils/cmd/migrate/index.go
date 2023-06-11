package migrate

import (
	"fmt"

	"gitee.com/lybbn/go-vue-lyadmin/global"
	mmodel "gitee.com/lybbn/go-vue-lyadmin/model"
	"github.com/spf13/cobra"
)

var (
	database string
	StartCmd = &cobra.Command{
		Use:     "migrate",
		Short:   "migrate the modles to database",
		Example: "golyadmin migrate",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.Flags().StringVarP(&database, "database", "d", "default", "database alias-name")
}

func run() {
	initDB()
}

// 迁移数据库表
func migrateModel() {
	db := global.GVLA_DB
	if database == "" {

	} else {
		db = global.GetGlobalDBByName(database)
	}

	tables := mmodel.MigrateModelList
	for _, t := range tables {
		_ = db.Debug().AutoMigrate(&t)
	}
}

func initDB() {
	//数据库迁移
	fmt.Println("数据库迁移开始")
	migrateModel()
	fmt.Println(`数据库迁移成功`)
}
