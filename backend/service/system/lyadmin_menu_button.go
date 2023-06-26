package system

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	systemReq "gitee.com/lybbn/golyadmin/model/system/request"
	"gorm.io/gorm"
)

type MenuButtonService struct{}

// 分页获取菜单按钮列表
func (m *MenuButtonService) GetLyadminMenuButtonList(info systemReq.LyadminMenuButtonSearch) *gorm.DB {
	// 创建db
	db := global.GL_DB.Model(&system.LyadminMenuButton{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Method != "" {
		db = db.Where("method LIKE ?", "%"+info.Method+"%")
	}
	if info.MenuID != 0 {
		db = db.Where("menu_id = ?", info.MenuID)
	}
	db = db.Order("id desc")
	return db
}

// 新增菜单按钮
func (m *MenuButtonService) CreateMenuButton(ReqData system.LyadminMenuButton) error {
	return global.GL_DB.Create(&ReqData).Error
}

// 删除菜单按钮
func (m *MenuButtonService) DeleteMenuButton(id uint) (err error) {
	err = global.GL_DB.Where("id = ?", id).Delete(&system.LyadminMenuButton{}).Error
	return err
}

// 编辑菜单按钮
func (m *MenuButtonService) UpdateMenuButton(ReqData system.LyadminMenuButton) (err error) {
	var oldData system.LyadminMenuButton
	upDateMap := make(map[string]interface{})
	upDateMap["menu_id"] = ReqData.MenuID
	upDateMap["name"] = ReqData.Name
	upDateMap["value"] = ReqData.Value
	upDateMap["api"] = ReqData.Api
	upDateMap["method"] = ReqData.Method
	upDateMap["update_by"] = ReqData.UpdateBy

	db := global.GL_DB.Where("id = ?", ReqData.ID).Find(&oldData)
	err = db.Updates(upDateMap).Error
	return err
}
