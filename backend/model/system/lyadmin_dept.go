package system

import (
	"gitee.com/lybbn/go-vue-lyadmin/global"
)

type LyadminDept struct {
	global.GVLA_BASE_MODEL
	ParentId uint   `json:"parent_id" gorm:"comment:上级部门"`
	Name     string `json:"name" gorm:"comment:部门名称"`
	Sort     int    `json:"sort" gorm:"default:1;comment:显示顺序"`
	Status   bool   `json:"status" gorm:"default:true;comment:状态"`
	Owner    string `json:"owner" gorm:"size:100;comment:负责人"`
	Phone    string `json:"phone" gorm:"size:32;comment:手机"`
	Email    string `json:"email" gorm:"size:64;comment:邮箱"`
}

func (LyadminDept) TableName() string {
	return "lyadmin_dept"
}
