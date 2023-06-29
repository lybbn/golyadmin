package system

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	systemReq "gitee.com/lybbn/golyadmin/model/system/request"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	// err = global.GL_DB.Where("id = ?", id).Delete(&system.LyadminRole{}).Error
	// return err
	return global.GL_DB.Transaction(func(tx *gorm.DB) error {
		var model = system.LyadminRole{}
		tx.Preload("Menu").Preload("Dept").Preload("Permission").First(&model, id)
		//删除 LyadminRole 时，同时删除角色所有 关联其它表 记录 (Menu\Dept\Permission)
		db := tx.Select(clause.Associations).Delete(&model)
		if err = db.Error; err != nil {
			return err
		}
		return err
	})
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
func (r *RoleService) UpdateRolePremission(ReqData systemReq.LyadminRoleParams) (err error) {
	return global.GL_DB.Transaction(func(tx *gorm.DB) error {
		tx.Model(&system.LyadminRole{}).Where("id = ?", ReqData.ID).Update("data_range", ReqData.DataRange)
		//many2many中间表更新方式一
		err = tx.Delete(&[]system.LyadminRoleMenu{}, "lyadmin_role_id = ?", ReqData.ID).Error
		if err != nil {
			return err
		}
		var roleMenu []system.LyadminRoleMenu
		for _, v := range ReqData.Menu {
			roleMenu = append(roleMenu, system.LyadminRoleMenu{
				LyadminRoleId: ReqData.ID,
				LyadminMenuId: uint(v),
			})
		}
		if len(roleMenu) > 0 {
			err = tx.Create(&roleMenu).Error
		}
		if err != nil {
			return err
		}
		//many2many中间表更新方式二
		var model = system.LyadminRole{}
		var roleDept []system.LyadminDept
		var rolePermission []system.LyadminMenuButton
		tx.Preload("Dept").Preload("Permission").First(&model, ReqData.ID)
		tx.Where("id in ?", ReqData.Dept).Find(&roleDept)
		tx.Where("id in ?", ReqData.Permission).Find(&rolePermission)
		// 删除LyadminRole 和 Dept 和 Permission 的关联关系
		err = tx.Model(&model).Association("Dept").Delete(model.Dept)
		if err != nil {
			return err
		}
		err = tx.Model(&model).Association("Permission").Delete(model.Permission)
		if err != nil {
			return err
		}
		model.Dept = roleDept
		model.Permission = rolePermission
		// 更新关联的数据，使用 FullSaveAssociations 模式
		db := tx.Model(&model).Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model)
		if err = db.Error; err != nil {
			return err
		}
		return err
	})
}
