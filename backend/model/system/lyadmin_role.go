package system

import (
	"gitee.com/lybbn/golyadmin/global"
)

type LyadminRole struct {
	global.GL_BASE_MODEL
	Name       string              `json:"name" gorm:"comment:角色名称"`
	Key        string              `json:"key" gorm:"comment:权限字符"`
	Sort       int                 `json:"sort" gorm:"default:1;comment:显示顺序"`
	Status     bool                `json:"status" gorm:"default:true;comment:状态"`
	Menu       []LyadminMenu       `json:"menu" gorm:"many2many:lyadmin_role_menu;"`             //manytomany
	Dept       []LyadminDept       `json:"dept" gorm:"many2many:lyadmin_role_dept;"`             //manytomany
	Permission []LyadminMenuButton `json:"permission" gorm:"many2many:lyadmin_role_menubutton;"` //manytomany
	global.GL_CONTROL_MODEL
}

func (LyadminRole) TableName() string {
	return "lyadmin_role"
}
