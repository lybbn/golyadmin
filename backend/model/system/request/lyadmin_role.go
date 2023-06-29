package request

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/utils/response"
)

type LyadminRoleSearch struct {
	Name   string `json:"name" form:"name"`     //名称
	Status *bool  `json:"status" form:"status"` //状态
	response.StructPageQueryParams
}

type LyadminRoleParams struct {
	global.GL_BASE_MODEL
	Name       string `json:"name" form:"name" gorm:"comment:角色名称"`
	Key        string `json:"key" form:"key" gorm:"index;unique;comment:权限字符"`
	Sort       int    `json:"sort" form:"sort" gorm:"default:1;comment:显示顺序"`
	DataRange  int    `json:"data_range" form:"data_range" gorm:"default:0;comment:数据权限范围"` //数据权限范围(0, "仅本人数据权限"),(1, "本部门及以下数据权限"),(2, "本部门数据权限"),(3, "全部数据权限"),(4, "自定数据权限"),
	Status     bool   `json:"status" form:"status" gorm:"default:true;comment:状态"`
	Menu       []int  `json:"menu" form:"menu"`
	Dept       []int  `json:"dept" form:"dept"`             //manytomany
	Permission []int  `json:"permission" form:"permission"` //manytomany
	global.GL_CONTROL_MODEL
}
