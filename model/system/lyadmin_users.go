package system

import (
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/utils/uuid"
	"gorm.io/gorm"
)

type LyadminUsers struct {
	global.GVLA_MODEL
	UUID string `gorm:"<-:create;type:varchar(50)" form:"uuid" json:"uuid"`
}

func (LyadminUsers) TableName() string {
	return "lyadmin_users"
}

func (u *LyadminUsers) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.MakeUUID()
	return
}
