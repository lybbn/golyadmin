package initdb

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
		Use:     "initdb",
		Short:   "initialize database with schema and seed data",
		Example: "golyadmin initdb",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.Flags().StringVarP(&database, "database", "d", "default", "database alias-name")
}

func run() {
	var db *gorm.DB
	if database == "default" || database == "" {
		db = global.GL_DB
	} else {
		db = global.GetGlobalDBByName(database)
	}
	if db == nil {
		panic("数据库连接为空，请检查配置文件")
	}

	// 表结构迁移（与 migrate 命令保持一致）
	fmt.Println("数据库表结构迁移开始")
	for _, t := range mmodel.MigrateModelList {
		err := db.Debug().Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&t)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("数据库表结构迁移成功")

	// 写入种子数据
	SeedAll(db)
}
