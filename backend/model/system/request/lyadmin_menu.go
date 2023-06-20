package request

import (
	"gitee.com/lybbn/go-vue-lyadmin/utils/response"
)

type LyadminMenuSearch struct {
	Name    string `json:"name" form:"name"`         //名称
	Visible *bool  `json:"visible" form:"visible"`   //是否显示
	Status  *bool  `json:"status" form:"status"`     //状态
	WebPath string `json:"web_path" form:"web_path"` //路径
	response.StructPageQueryParams
}
