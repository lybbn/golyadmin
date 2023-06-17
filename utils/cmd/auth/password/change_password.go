package password

import (
	"fmt"

	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/model/system"
	"gitee.com/lybbn/go-vue-lyadmin/utils"
	"github.com/spf13/cobra"
)

var (
	username string
	password string
	StartCmd = &cobra.Command{
		Use:     "changepassword",
		Short:   "change a user password",
		Example: "golyadmin changepassword -u superadmin -p 123456",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.Flags().StringVarP(&username, "username", "u", "", "superadmin username")
	StartCmd.Flags().StringVarP(&password, "password", "p", "", "superadmin password")
}

func run() {
	if username == "" {
		fmt.Println(utils.Red("请使用-u 指定要修改密码的用户名"))
		return
	}
	if password == "" {
		fmt.Println(utils.Red("请使用-p 指定要修改的密码"))
		return
	}
	// 加密密码
	enpassword := utils.MakePassowrd(password)
	result := global.GVLA_DB.Model(&system.LyadminUsers{}).Where("username = ?", username).Update("password", enpassword)

	if result.RowsAffected <= 0 {
		fmt.Println(utils.Red("用户 " + username + " 不存在"))
		return
	}

	fmt.Println("修改用户密码成功！")
	fmt.Println("账号：" + utils.Green(username))
	fmt.Println("新密码：" + utils.Green(password))
}
