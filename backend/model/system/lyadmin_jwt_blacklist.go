package system

import "gitee.com/lybbn/golyadmin/global"

type LyadminJwtBlacklist struct {
	global.GL_BASE_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (LyadminJwtBlacklist) TableName() string {
	return "lyadmin_jwt_blacklist"
}
