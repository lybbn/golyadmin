package system

import "gitee.com/lybbn/go-vue-lyadmin/global"

type LyadminJwtBlacklist struct {
	global.GVLA_BASE_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (LyadminJwtBlacklist) TableName() string {
	return "lyadmin_jwt_blacklist"
}
