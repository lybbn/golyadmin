package superadmin

import (
	"fmt"

	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	"gitee.com/lybbn/golyadmin/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	username string
	password string
	StartCmd = &cobra.Command{
		Use:     "createsuperuser",
		Short:   "create a super admin user",
		Example: "golyadmin createsuperuser -u superadmin -p 123456",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.Flags().StringVarP(&username, "username", "u", "superadmin", "superadmin username")
	StartCmd.Flags().StringVarP(&password, "password", "p", "123456", "superadmin password")
}

func run() {
	if username == "" {
		username = "superadmin"
	}
	if password == "" {
		password = "123456"
	}
	var count int64
	if err := global.GL_DB.Model(&system.LyadminUsers{}).Where("is_superuser = 1").Count(&count).Error; err != nil {
		fmt.Println(err.Error())
		return
	}
	if count > 0 {
		fmt.Println(utils.Red("已存在超级管理员，无需再创建，如需更改超级管理员密码请使用changepassword命令！"))
		return
	}
	user := &system.LyadminUsers{Username: username, Nickname: "超级管理员", Name: "超级管理员", Password: password, IsSuperuser: true, Identity: 1}
	// 加密密码
	user.Password = utils.MakePassowrd(password)
	err := global.GL_DB.Create(&user).Error

	if err != nil {
		fmt.Println(utils.Red("创建超级管理员失败!"), zap.Error(err))
		return
	}
	fmt.Println("创建超级管理员成功！")
	fmt.Println("账号：" + utils.Green(username))
	fmt.Println("密码：" + utils.Green(password))
}
