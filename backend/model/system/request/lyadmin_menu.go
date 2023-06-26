package request

import (
	"gitee.com/lybbn/golyadmin/utils/response"
)

type LyadminMenuSearch struct {
	Search  string `json:"search" form:"search"`     //搜索关键词
	Name    string `json:"name" form:"name"`         //名称
	Visible string `json:"visible" form:"visible"`   //是否显示
	Status  string `json:"status" form:"status"`     //状态
	WebPath string `json:"web_path" form:"web_path"` //路径
	response.StructPageQueryParams
}
