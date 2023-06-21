package migrate

import (
	"fmt"

	"gitee.com/lybbn/golyadmin/global"
	mmodel "gitee.com/lybbn/golyadmin/model"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
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
	var db *gorm.DB
	if database == "" {
		db = global.GL_DB
	} else {
		db = global.GetGlobalDBByName(database)
	}

	tables := mmodel.MigrateModelList
	for _, t := range tables {
		err := db.Debug().Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&t)
		if err != nil {
			panic(err)
		}
	}
}

func initDB() {
	//数据库迁移
	fmt.Println("数据库迁移开始")
	migrateModel()
	fmt.Println(`数据库迁移成功`)
}
