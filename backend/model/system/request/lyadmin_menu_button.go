package request

import (
	"gitee.com/lybbn/golyadmin/utils/response"
)

type LyadminMenuButtonSearch struct {
	Name   string `json:"name" form:"name"`       //名称
	Method string `json:"method" form:"method"`   //方法
	MenuID uint   `json:"menu_id" form:"menu_id"` //方法
	response.StructPageQueryParams
}
