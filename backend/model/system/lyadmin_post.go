package system

import (
	"gitee.com/lybbn/golyadmin/global"
)

type LyadminPost struct {
	global.GVLA_BASE_MODEL
	Name   string `json:"name" gorm:"comment:岗位名称"`
	Code   string `json:"code" gorm:"size:100;comment:岗位编码"`
	Sort   int    `json:"sort" gorm:"default:1;comment:显示顺序"`
	Status bool   `json:"status" gorm:"default:true;comment:状态"`
}

func (LyadminPost) TableName() string {
	return "lyadmin_post"
}
