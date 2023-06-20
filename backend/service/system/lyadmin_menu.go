package system

import (
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/model/system"
	systemReq "gitee.com/lybbn/go-vue-lyadmin/model/system/request"
	"gorm.io/gorm"
)

type MenuService struct{}

// 分页获取菜单列表
func (m *MenuService) GetLyadminMenuList(info systemReq.LyadminMenuSearch) *gorm.DB {
	// 创建db
	db := global.GVLA_DB.Model(&system.LyadminMenu{})
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
	db = db.Order("id desc")
	return db
}
