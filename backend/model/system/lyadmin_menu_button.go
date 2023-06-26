package system

import (
	"gitee.com/lybbn/golyadmin/global"
)

type LyadminMenuButton struct {
	global.GL_BASE_MODEL
	MenuID uint   `json:"menu_id" form:"menu_id" gorm:"comment:关联菜单ID"`
	Name   string `json:"name" form:"name" gorm:"varchar(30);comment:名称"`
	Value  string `json:"value" form:"value" gorm:"comment:权限值"`
	Api    string `json:"api" form:"api" gorm:"comment:接口地址"`
	Method string `json:"method" form:"method" gorm:"varchar(30);comment:接口请求方法"`
	global.GL_CONTROL_MODEL
}

func (LyadminMenuButton) TableName() string {
	return "lyadmin_menu_button"
}

type LyadminButton struct {
	global.GL_BASE_MODEL
	Name  string `json:"name" form:"name" gorm:"varchar(30);comment:按钮名称"`
	Value string `json:"value" form:"value" gorm:"comment:按钮值"`
	global.GL_CONTROL_MODEL
}

func (LyadminButton) TableName() string {
	return "lyadmin_button"
}
