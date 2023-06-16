package system

import "gitee.com/lybbn/go-vue-lyadmin/global"

type JwtBlacklist struct {
	global.GVLA_BASE_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
