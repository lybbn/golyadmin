package system

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	systemReq "gitee.com/lybbn/golyadmin/model/system/request"
	"gorm.io/gorm"
)

type RoleService struct{}

// 分页获取角色列表
func (r *RoleService) GetLyadminRoleList(info systemReq.LyadminRoleSearch) *gorm.DB {
	// 创建db
	db := global.GL_DB.Model(&system.LyadminRole{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	db = db.Order("sort asc").Preload("Dept").Preload("Menu").Preload("Permission")
	return db
}

// 获取所有菜单按钮
func (r *RoleService) GetRoleMenuById() (menus []system.LyadminMenu, err error) {
	err = global.GL_DB.Preload("MenuButtons").Order("sort asc").Find(&menus).Error
	if err != nil {
		return menus, err
	}
	return menus, err
}

// 新增角色
func (r *RoleService) CreateRole(ReqData system.LyadminRole) error {
	// if !errors.Is(global.GL_DB.Where("name = ?", role.Name).First(&system.LyadminRole{}).Error, gorm.ErrRecordNotFound) {
	// 	return errors.New("存在重复name，请修改name")
	// }
	return global.GL_DB.Create(&ReqData).Error
}

// 删除角色
func (r *RoleService) DeleteRole(id uint) (err error) {
	err = global.GL_DB.Where("id = ?", id).Delete(&system.LyadminRole{}).Error
	return err
}

// 编辑角色
func (r *RoleService) UpdateRole(ReqData system.LyadminRole) (err error) {
	var oldData system.LyadminRole
	upDateMap := make(map[string]interface{})
	upDateMap["name"] = ReqData.Name
	upDateMap["key"] = ReqData.Key
	upDateMap["sort"] = ReqData.Sort
	upDateMap["status"] = ReqData.Status
	upDateMap["update_by"] = ReqData.UpdateBy

	db := global.GL_DB.Where("id = ?", ReqData.ID).Find(&oldData)
	err = db.Updates(upDateMap).Error
	return err
}

// 更新角色权限
func (r *RoleService) UpdateRolePremission(ReqData system.LyadminRole) error {
	var s system.LyadminRole
	global.GL_DB.Preload("Menu").Preload("Dept").Preload("Permission").First(&s, "id = ?", ReqData.ID)
	global.GL_DB.Model(&s).Update("data_range", ReqData.DataRange)
	err := global.GL_DB.Model(&s).Association("Menu").Replace(&ReqData.Menu)
	if err != nil {
		return err
	}
	err = global.GL_DB.Model(&s).Association("Dept").Replace(&ReqData.Dept)
	if err != nil {
		return err
	}
	err = global.GL_DB.Model(&s).Association("Permission").Replace(&ReqData.Permission)
	return err
}
