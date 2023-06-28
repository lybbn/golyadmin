package system

import (
	"gitee.com/lybbn/golyadmin/global"
)

type LyadminRole struct {
	global.GL_BASE_MODEL
	Name       string              `json:"name" form:"name" gorm:"comment:角色名称" binding:"required" msg:"名称不能为空"`
	Key        string              `json:"key" form:"key" gorm:"index;unique;comment:权限字符" binding:"required" msg:"权限字符不能为空"`
	Sort       int                 `json:"sort" form:"sort" gorm:"default:1;comment:显示顺序"`
	DataRange  int                 `json:"data_range" form:"data_range" gorm:"default:0;comment:数据权限范围"` //数据权限范围(0, "仅本人数据权限"),(1, "本部门及以下数据权限"),(2, "本部门数据权限"),(3, "全部数据权限"),(4, "自定数据权限"),
	Status     bool                `json:"status" form:"status" gorm:"default:true;comment:状态"`
	Menu       []LyadminMenu       `json:"menu" form:"menu" gorm:"many2many:lyadmin_role_menu;"`                   //manytomany
	Dept       []LyadminDept       `json:"dept" form:"dept" gorm:"many2many:lyadmin_role_dept;"`                   //manytomany
	Permission []LyadminMenuButton `json:"permission" form:"permission" gorm:"many2many:lyadmin_role_menubutton;"` //manytomany
	global.GL_CONTROL_MODEL
}

func (LyadminRole) TableName() string {
	return "lyadmin_role"
}
