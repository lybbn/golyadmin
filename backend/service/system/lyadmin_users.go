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
	"gorm.io/gorm/clause"
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

// 新增管理员
func (s *UserService) CreateAdminUser(reqData systemReq.CreateUserRequestParams) (err error) {
	return global.GL_DB.Transaction(func(tx *gorm.DB) error {
		var i int64
		err = global.GL_DB.Model(&system.LyadminUsers{}).Where("username = ?", reqData.Username).Count(&i).Error
		if err != nil {
			global.GL_LOG.Error("创建用户失败!", zap.Error(err))
			return err
		}
		if i > 0 {
			err = errors.New("用户名已存在！")
			return err
		}
		model := &system.LyadminUsers{Name: reqData.Name, Username: reqData.Username, Nickname: reqData.Nickname, Password: reqData.Password, Avatar: reqData.Avatar, Mobile: reqData.Mobile, Gender: reqData.Gender, Email: reqData.Email, DeptId: reqData.DeptId, IsActive: reqData.IsActive}
		model.CreateBy = reqData.CreateBy
		model.BelongDept = reqData.BelongDept
		// 加密密码
		model.Password = utils.MakePassowrd(reqData.Password)
		err = global.GL_DB.Create(&model).Error

		if err != nil {
			return err
		}
		//更新关联角色
		var userRoles []system.LyadminUsersRole
		for _, v := range reqData.RoleIds {
			userRoles = append(userRoles, system.LyadminUsersRole{
				LyadminUsersId: model.ID,
				LyadminRoleId:  uint(v),
			})
		}
		if len(userRoles) > 0 {
			err = tx.Create(&userRoles).Error
		}
		if err != nil {
			return err
		}
		return nil
	})
}

// 编辑管理员
func (s *UserService) UpdateAdminUser(reqData systemReq.UpdateUsersRequestParams, c *gin.Context) (err error) {
	return global.GL_DB.Transaction(func(tx *gorm.DB) error {
		var model system.LyadminUsers
		db := tx.Scopes(utils.DataLevelPermissionsFilter(system.LyadminUsers{}, c)).First(&model, reqData.ID)
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
		model.Mobile = reqData.Mobile
		model.DeptId = reqData.DeptId
		model.IsActive = reqData.IsActive
		model.UpdateBy = reqData.UpdateBy
		if reqData.Password != "" {
			model.Password = utils.MakePassowrd(reqData.Password)
		}
		db.Save(&model)
		if db.RowsAffected == 0 {
			err = errors.New("update UpdateAdminUser error")
			return err
		}
		//更新关联角色
		err = tx.Delete(&[]system.LyadminUsersRole{}, "lyadmin_users_id = ?", reqData.ID).Error
		if err = tx.Error; err != nil {
			global.GL_LOG.Error("db error", zap.Error(err))
			return err
		}
		var userRoles []system.LyadminUsersRole
		for _, v := range reqData.RoleIds {
			userRoles = append(userRoles, system.LyadminUsersRole{
				LyadminUsersId: reqData.ID,
				LyadminRoleId:  uint(v),
			})
		}
		if len(userRoles) > 0 {
			err = tx.Create(&userRoles).Error
		}
		if err != nil {
			return err
		}
		return nil
	})
}

// 删除管理员
func (s *UserService) DeleteAdminUser(id uint, c *gin.Context) (err error) {
	return global.GL_DB.Transaction(func(tx *gorm.DB) error {
		var model = system.LyadminUsers{}
		res := tx.Scopes(utils.DataLevelPermissionsFilter(system.LyadminUsers{}, c)).Preload("Dept").Preload("Role").First(&model, id)
		if err = res.Error; err != nil {
			global.GL_LOG.Error("Service DeleteAdminUser error", zap.Error(err))
			return err
		}
		if res.RowsAffected == 0 {
			return errors.New("无权删除该数据")
		}
		//删除 LyadminUsers 时，同时删除用户所关联其它表 记录 (Dept\Role)
		db := tx.Select(clause.Associations).Delete(&model)
		if err = db.Error; err != nil {
			return err
		}
		return err
	})
}

func (s *UserService) ChangePassword(u *system.LyadminUsers, newPassword string, c *gin.Context) (userInter *system.LyadminUsers, err error) {
	var user system.LyadminUsers
	res := global.GL_DB.Where("id = ?", u.ID).Scopes(utils.DataLevelPermissionsFilter(system.LyadminUsers{}, c)).First(&user)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("无权删除该数据")
	}
	if ok := utils.CheckPassword(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.MakePassowrd(newPassword)
	err = global.GL_DB.Save(&user).Error
	return &user, err

}

// 获取前台用户信息列表
func (userService *UserService) GetUserInfoList(info systemReq.LyadminUserSearch) *gorm.DB {
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
	db = db.Where("Identity = ?", 3).Preload("Role").Preload("Dept").Order("id desc")
	return db
}

// 删除用户
func (s *UserService) DeleteUser(id uint, c *gin.Context) (err error) {
	return global.GL_DB.Transaction(func(tx *gorm.DB) error {
		var model = system.LyadminUsers{}
		res := tx.Scopes(utils.DataLevelPermissionsFilter(system.LyadminUsers{}, c)).First(&model, id)
		if err := res.Error; err != nil {
			return err
		}
		if res.RowsAffected == 0 {
			return errors.New("无权删除该数据")
		}
		//删除 LyadminUsers
		db := tx.Select(clause.Associations).Delete(&model)
		if err = db.Error; err != nil {
			return err
		}
		return err
	})
}

// 修改用户状态
func (s *UserService) DisableUser(reqData systemReq.DisableUserReq, c *gin.Context) (err error) {
	var model system.LyadminUsers
	res := global.GL_DB.Where("id = ?", reqData.ID).Scopes(utils.DataLevelPermissionsFilter(system.LyadminUsers{}, c)).First(&model)
	if err := res.Error; err != nil {
		return err
	}
	if res.RowsAffected == 0 {
		return errors.New("无该数据权限")
	}
	model.IsActive = reqData.IsActive
	err = global.GL_DB.Save(&model).Error
	return err

}

// 新增用户
func (s *UserService) CreateUser(reqData systemReq.CreateUserRequestParams) (err error) {
	return global.GL_DB.Transaction(func(tx *gorm.DB) error {
		var i int64
		err = global.GL_DB.Model(&system.LyadminUsers{}).Where("username = ?", reqData.Username).Count(&i).Error
		if err != nil {
			global.GL_LOG.Error("创建用户失败!", zap.Error(err))
			return err
		}
		if i > 0 {
			err = errors.New("用户名已存在！")
			return err
		}
		model := &system.LyadminUsers{Name: reqData.Name, Username: reqData.Username, Nickname: reqData.Nickname, Password: reqData.Password, Avatar: reqData.Avatar, Mobile: reqData.Mobile, Gender: reqData.Gender, Email: reqData.Email, IsActive: reqData.IsActive}
		model.CreateBy = reqData.CreateBy
		model.BelongDept = reqData.BelongDept
		model.Identity = 3
		model.IsStaff = false
		model.IsSuperuser = false
		// 加密密码
		model.Password = utils.MakePassowrd(reqData.Password)
		err = global.GL_DB.Create(&model).Error

		if err != nil {
			return err
		}
		return nil
	})
}

// 编辑用户
func (s *UserService) UpdateUser(reqData systemReq.UpdateUsersRequestParams, c *gin.Context) (err error) {
	return global.GL_DB.Transaction(func(tx *gorm.DB) error {
		var model system.LyadminUsers
		db := tx.Scopes(utils.DataLevelPermissionsFilter(system.LyadminUsers{}, c)).First(&model, reqData.ID)
		if err = db.Error; err != nil {
			global.GL_LOG.Error("Service UpdateUser error", zap.Error(err))
			return err
		}
		if db.RowsAffected == 0 {
			err = errors.New("无该数据权限")
			return err
		}
		model.Avatar = reqData.Avatar
		model.Username = reqData.Username
		model.Name = reqData.Name
		model.Mobile = reqData.Mobile
		model.IsActive = reqData.IsActive
		model.UpdateBy = reqData.UpdateBy
		if reqData.Password != "" {
			model.Password = utils.MakePassowrd(reqData.Password)
		}
		db.Save(&model)
		if db.RowsAffected == 0 {
			err = errors.New("update UpdateUser error")
			return err
		}
		return nil
	})
}
