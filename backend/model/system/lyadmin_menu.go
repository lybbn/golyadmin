package system

import (
	"gitee.com/lybbn/golyadmin/global"
)

type LyadminMenu struct {
	global.GVLA_BASE_MODEL
	ParentId      uint                `json:"parent_id" gorm:"comment:父菜单ID"`
	Name          string              `json:"name" gorm:"comment:菜单名称"`
	Icon          string              `json:"icon" gorm:"comment:菜单图标"`
	WebPath       string              `json:"web_path" gorm:"comment:路由地址"`
	IsLink        bool                `json:"is_link" gorm:"default:false;comment:是否外链"`
	Visible       bool                `json:"visible" gorm:"default:true;comment:是否显示菜单"`
	Component     string              `json:"component" gorm:"comment:对应前端文件路径"`
	ComponentName string              `json:"component_name" gorm:"comment:对应前端文件名称"`
	Sort          int                 `json:"sort" gorm:"default:1;comment:显示顺序"`
	KeepAlive     bool                `json:"keep_alive" gorm:"default:false;comment:是否缓存页面"`
	Status        bool                `json:"status" gorm:"default:true;comment:状态"`
	MenuButton    []LyadminMenuButton `json:"menu_button"` //has many
}

func (LyadminMenu) TableName() string {
	return "lyadmin_menu"
}
