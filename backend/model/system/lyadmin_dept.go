package system

import (
	"gitee.com/lybbn/golyadmin/global"
)

type LyadminDept struct {
	global.GL_BASE_MODEL
	ParentId uint   `json:"parent_id" gorm:"comment:上级部门"`
	Name     string `json:"name" gorm:"comment:部门名称" binding:"required" msg:"名称不能为空"`
	Sort     int    `json:"sort" gorm:"default:1;comment:显示顺序"`
	Status   bool   `json:"status" gorm:"default:true;comment:状态" binding:"required" msg:"状态不能为空"`
	Owner    string `json:"owner" gorm:"size:100;comment:负责人"`
	Phone    string `json:"phone" gorm:"size:32;comment:手机"`
	Email    string `json:"email" gorm:"size:64;comment:邮箱"`
	global.GL_CONTROL_MODEL
}

func (LyadminDept) TableName() string {
	return "lyadmin_dept"
}
