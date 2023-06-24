package request

import (
	"gitee.com/lybbn/golyadmin/utils/response"
)

type LyadminMenuButtonSearch struct {
	Name   string `json:"name" form:"name"`     //名称
	Method string `json:"method" form:"method"` //方法
	response.StructPageQueryParams
}
