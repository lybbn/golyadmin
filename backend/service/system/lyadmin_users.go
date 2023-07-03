package system

import (
	"errors"

	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	systemReq "gitee.com/lybbn/golyadmin/model/system/request"
	"gitee.com/lybbn/golyadmin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct{}

var UserServiceApp = new(UserService)

func (s *UserService) FindUserById(id uint) (user *system.LyadminUsers, err error) {
	var u system.LyadminUsers
	err = global.GL_DB.Where("`id` = ?", id).First(&u).Error
	return &u, err
}

func (s *UserService) GetUserInfoById(id uint) (user system.LyadminUsers, err error) {
	var u system.LyadminUsers
	err = global.GL_DB.Preload("Role").Preload("Role.Dept").First(&u, "id = ?", id).Error
	if err != nil {
		return u, err
	}
	return u, err
}

// 设置用户信息
func (s *UserService) SetUserInfo(req systemReq.ChangeUserInfo, id uint) error {
	return global.GL_DB.Model(&system.LyadminUsers{}).Where("id=?", id).Updates(req).Error
}

// 获取管理员用户信息列表
func (userService *UserService) GetAdminUserInfoList(info systemReq.LyadminUserSearch) *gorm.DB {
	// 创建db
	db := global.GL_DB.Model(&system.LyadminUsers{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.IsActive != "" {
		is_active, err := utils.FormatString2Bool(info.IsActive)
		if err == nil {
			db = db.Where("is_active = ?", is_active)
		}
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Mobile != "" {
		db = db.Where("name LIKE ?", "%"+info.Mobile+"%")
	}
	if info.Username != "" {
		db = db.Where("name LIKE ?", "%"+info.Username+"%")
	}
	if info.Search != "" {
		db = db.Where("name LIKE ? or mobile LIKE ?", "%"+info.Search+"%", "%"+info.Search+"%")
	}
	if info.BeginAt != "" {
		db = db.Where("created_at between ? and ?", info.BeginAt, info.EndAt)
	}
	db = db.Where("Identity = ?", 2).Preload("Role").Preload("Dept").Order("id desc")
	return db
}

// 编辑管理员
func (s *UserService) UpdateAdminUser(reqData system.LyadminUsers, c *gin.Context) (err error) {
	var model system.LyadminUsers
	db := global.GL_DB.Scopes(utils.DataLevelPermissionsFilter(system.LyadminUsers{}, c)).First(&model, reqData.ID)
	if err = db.Error; err != nil {
		global.GL_LOG.Error("Service UpdateAdminUser error", zap.Error(err))
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无该数据权限")
		return err
	}
	model.Username = reqData.Username
	model.Name = reqData.Name
	model.DeptId = reqData.DeptId
	model.IsActive = reqData.IsActive
	model.UpdateBy = reqData.UpdateBy
	db.Save(&model)
	if err = db.Error; err != nil {
		global.GL_LOG.Error("db error", zap.Error(err))
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("update UpdateAdminUser error")
		return err
	}
	return nil
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
