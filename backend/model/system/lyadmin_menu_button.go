package system

import (
	"gitee.com/lybbn/golyadmin/global"
)

type LyadminMenuButton struct {
	global.GVLA_BASE_MODEL
	LyadminMenuID uint   `json:"menu_id" gorm:"comment:关联菜单ID"`
	Name          string `json:"name" gorm:"varchar(30);comment:名称"`
	Value         string `json:"value" gorm:"comment:权限值"`
	Api           string `json:"api" gorm:"comment:接口地址"`
	Method        string `json:"method" gorm:"varchar(30);comment:接口请求方法"`
}

func (LyadminMenuButton) TableName() string {
	return "lyadmin_menu_button"
}

type LyadminButton struct {
	global.GVLA_BASE_MODEL
	Name  string `json:"name" gorm:"varchar(30);comment:按钮名称"`
	Value string `json:"value" gorm:"comment:按钮值"`
}

func (LyadminButton) TableName() string {
	return "lyadmin_button"
}
