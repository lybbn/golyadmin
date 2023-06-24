package system

import (
	"errors"

	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	"gitee.com/lybbn/golyadmin/utils"
)

type UserService struct{}

func (s *UserService) FindUserById(id uint) (user *system.LyadminUsers, err error) {
	var u system.LyadminUsers
	err = global.GL_DB.Where("`id` = ?", id).First(&u).Error
	return &u, err
}

func (s *UserService) ChangePassword(u *system.LyadminUsers, newPassword string) (userInter *system.LyadminUsers, err error) {
	var user system.LyadminUsers
	if err = global.GL_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.CheckPassword(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.MakePassowrd(newPassword)
	err = global.GL_DB.Save(&user).Error
	return &user, err

}
