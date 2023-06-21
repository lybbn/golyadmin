package system

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
)

type UserService struct{}

func (s *UserService) FindUserById(id uint) (user *system.LyadminUsers, err error) {
	var u system.LyadminUsers
	err = global.GL_DB.Where("`id` = ?", id).First(&u).Error
	return &u, err
}
