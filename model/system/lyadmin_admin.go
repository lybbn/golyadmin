package system

import (
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/utils"
	"gorm.io/gorm"
)

type SysAdmin struct {
	global.GVLA_MODEL
	UUID     string `gorm:"<-:create;type:varchar(50)" form:"uuid" json:"uuid"` // 允许读和创建
	Username string `json:"username" gorm:"index;comment:用户登录名"`
	Password string `json:"-"  gorm:"comment:用户登录密码"`
	Nickname string `json:"nickname" gorm:"default:系统用户;comment:用户昵称"`
	Mobile   string `json:"mobile"  gorm:"comment:用户手机号"`
	Email    string `json:"email"  gorm:"comment:用户邮箱"`
}

func (SysAdmin) TableName() string {
	return "lyadmin_admin"
}

func (u *SysAdmin) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = utils.MakeUUID()
	return
}
