package system

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	systemReq "gitee.com/lybbn/golyadmin/model/system/request"
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
