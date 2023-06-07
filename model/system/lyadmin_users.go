package system

import (
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/utils"
	"gorm.io/gorm"
)

type SysUsers struct {
	global.GVLA_MODEL
	UUID string `gorm:"<-:create;type:varchar(50)" form:"uuid" json:"uuid"`
}

func (SysUsers) TableName() string {
	return "lyadmin_users"
}

func (u *SysUsers) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = utils.MakeUUID()
	return
}
