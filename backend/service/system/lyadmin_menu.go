package system

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	systemReq "gitee.com/lybbn/golyadmin/model/system/request"
	systemResp "gitee.com/lybbn/golyadmin/model/system/response"
	"gitee.com/lybbn/golyadmin/utils"
	"gorm.io/gorm"
)

type MenuService struct{}

// 分页获取菜单列表
func (m *MenuService) GetLyadminMenuList(info systemReq.LyadminMenuSearch) *gorm.DB {
	// 创建db
	db := global.GL_DB.Model(&system.LyadminMenu{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Status != "" {
		status, err := utils.FormatString2Bool(info.Status)
		if err == nil {
			db = db.Where("status = ?", status)
		}
	}
	if info.Visible != "" {
		visible, err := utils.FormatString2Bool(info.Visible)
		if err == nil {
			db = db.Where("visible = ?", visible)
		}
	}
	if info.WebPath != "" {
		db = db.Where("web_path LIKE ?", "%"+info.WebPath+"%")
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Search != "" {
		db = db.Where("name LIKE ? or web_path LIKE ?", "%"+info.Search+"%", "%"+info.Search+"%")
	}
	db = db.Order("sort asc")
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
// db.Preload() 该方法请求参数可以添加func方法，在func中传入gorm.DB,可添加其他sql语句，如果多级别预加载，则层级预加载
func (m *MenuService) GetWebRouter(uinfo *utils.CustomClaims) (menus []systemResp.LyadminWebRouterResponse, err error) {
	userid := uinfo.BaseClaims.ID
	identity := uinfo.BaseClaims.Identity
	var mn []system.LyadminMenu
	if identity == 1 {
		err = global.GL_DB.Preload("MenuButtons").Where("status = ?", 1).Find(&mn).Error
		if err != nil {
			return menus, err
		}
		for _, v := range mn {
			var btnValue []string
			for _, vm := range v.MenuButtons {
				btnValue = append(btnValue, vm.Value)
			}
			menus = append(menus, systemResp.LyadminWebRouterResponse{
				ID:             v.ID,
				ParentId:       v.ParentId,
				Name:           v.Name,
				Icon:           v.Icon,
				WebPath:        v.WebPath,
				IsLink:         v.IsLink,
				Visible:        v.Visible,
				Component:      v.Component,
				ComponentName:  v.ComponentName,
				Sort:           v.Sort,
				IsCatalog:      v.IsCatalog,
				KeepAlive:      v.KeepAlive,
				MenuPermission: btnValue,
			})
		}
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
		var rolelist []system.LyadminRole
		err = global.GL_DB.Model(&system.LyadminRole{}).Where("status = ? and id in (?)", 1, roleIds).Preload("Menu", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort asc")
		}).Preload("Permission").Find(&rolelist).Error
		if err != nil {
			return menus, err
		}
		var menubtnlist []system.LyadminMenuButton
		for _, rl := range rolelist {
			for _, vm := range rl.Menu {
				if isContainMenu(mn, vm) == false {
					mn = append(mn, vm)
				}
			}
			for _, mt := range rl.Permission {
				if isContainMenuButton(menubtnlist, mt) == false {
					menubtnlist = append(menubtnlist, mt)
				}
			}
		}

		for _, v := range mn {
			var btnValue []string
			for _, vm := range menubtnlist {
				if vm.MenuID == v.ID {
					btnValue = append(btnValue, vm.Value)
				}
			}
			menus = append(menus, systemResp.LyadminWebRouterResponse{
				ID:             v.ID,
				ParentId:       v.ParentId,
				Name:           v.Name,
				Icon:           v.Icon,
				WebPath:        v.WebPath,
				IsLink:         v.IsLink,
				Visible:        v.Visible,
				Component:      v.Component,
				ComponentName:  v.ComponentName,
				Sort:           v.Sort,
				IsCatalog:      v.IsCatalog,
				KeepAlive:      v.KeepAlive,
				MenuPermission: btnValue,
			})
		}
		return menus, err
	}
}

func isContainMenu(items []system.LyadminMenu, item system.LyadminMenu) bool {
	for _, eachItem := range items {
		if eachItem.ID == item.ID {
			return true
		}
	}
	return false
}

func isContainMenuButton(items []system.LyadminMenuButton, item system.LyadminMenuButton) bool {
	for _, eachItem := range items {
		if eachItem.ID == item.ID {
			return true
		}
	}
	return false
}
