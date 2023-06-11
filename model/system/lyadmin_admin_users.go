package system

import (
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/utils/uuid"
	"gorm.io/gorm"
)

type LyadminAdminUsers struct {
	global.GVLA_MODEL
	UUID     string `gorm:"<-:create;type:varchar(50);comment:uuid" form:"uuid" json:"uuid"` // 允许读和创建
	Username string `json:"username" gorm:"type:varchar(50);not null;index;comment:用户名"`
	Password string `json:"-"  gorm:"type:varchar(128);comment:密码"`
	Nickname string `json:"nickname" gorm:"type:varchar(20);default:系统用户;comment:昵称"`
	Mobile   string `json:"mobile"  gorm:"type:char(25);comment:手机号"`
	Email    string `json:"email"  gorm:"type:varchar(100);comment:邮箱"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255);comment:头像"`
	Gender   string `json:"gender" gorm:"type:varchar(10);comment:性别"`
}

func (LyadminAdminUsers) TableName() string {
	return "lyadmin_admin_users"
}

func (u *LyadminAdminUsers) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.MakeUUID()
	return
}
