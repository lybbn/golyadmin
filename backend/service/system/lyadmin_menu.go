package system

import (
	"fmt"

	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	systemReq "gitee.com/lybbn/golyadmin/model/system/request"
	systemResp "gitee.com/lybbn/golyadmin/model/system/response"
	"gorm.io/gorm"
)

type MenuService struct{}

// 分页获取菜单列表
func (m *MenuService) GetLyadminMenuList(info systemReq.LyadminMenuSearch) *gorm.DB {
	// 创建db
	db := global.GL_DB.Model(&system.LyadminMenu{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.Visible != nil {
		db = db.Where("visible = ?", info.Visible)
	}
	if info.WebPath != "" {
		db = db.Where("web_path LIKE ?", "%"+info.WebPath+"%")
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	db = db.Order("sort aesc")
	return db
}

// 新增菜单
func (m *MenuService) CreateMenu(menu system.LyadminMenu) error {
	// if !errors.Is(global.GL_DB.Where("name = ?", menu.Name).First(&system.LyadminMenu{}).Error, gorm.ErrRecordNotFound) {
	// 	return errors.New("存在重复name，请修改name")
	// }
	return global.GL_DB.Create(&menu).Error
}

// 删除菜单
func (m *MenuService) DeleteMenu(id uint) (err error) {
	err = global.GL_DB.Where("id = ?", id).Delete(&system.LyadminMenu{}).Error
	return err
}

// 编辑菜单
func (m *MenuService) UpdateMenu(ReqData system.LyadminMenu) (err error) {
	var oldData system.LyadminMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = ReqData.KeepAlive
	upDateMap["parent_id"] = ReqData.ParentId
	upDateMap["web_path"] = ReqData.WebPath
	upDateMap["icon"] = ReqData.Icon
	upDateMap["name"] = ReqData.Name
	upDateMap["visible"] = ReqData.Visible
	upDateMap["component"] = ReqData.Component
	upDateMap["component_name"] = ReqData.ComponentName
	upDateMap["is_link"] = ReqData.IsLink
	upDateMap["sort"] = ReqData.Sort
	upDateMap["is_catalog"] = ReqData.IsCatalog
	upDateMap["status"] = ReqData.Status
	upDateMap["update_by"] = ReqData.UpdateBy

	db := global.GL_DB.Where("id = ?", ReqData.ID).Find(&oldData)
	err = db.Updates(upDateMap).Error
	return err
}

// 根据角色获取菜单权限资源
func (m *MenuService) GetWebRouter(userid uint, identity int) (menus []systemResp.LyadminWebRouterResponse, err error) {
	if identity == 1 {
		return menus, err
	} else {
		var roleIds []uint
		err = global.GL_DB.Model(&system.LyadminUsersRole{}).Where("lyadmin_users_id = ?", userid).Pluck("lyadmin_role_id", &roleIds).Error
		if err != nil {
			return menus, err
		}
		if len(roleIds) < 1 {
			return menus, err
		}
		var menuIds []uint
		err = global.GL_DB.Model(&system.LyadminRoleMenu{}).Where("lyadmin_role_id in (?)", roleIds).Distinct("lyadmin_menu_id").Pluck("lyadmin_menu_id", &menuIds).Error
		if err != nil {
			return menus, err
		}
		var mn system.LyadminMenu
		err = global.GL_DB.Where("id in (?)", menuIds).Find(&mn).Error
		if err != nil {
			return menus, err
		}
		fmt.Println(mn)
		return menus, err
	}
}
