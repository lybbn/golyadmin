package request

import (
	"gitee.com/lybbn/golyadmin/utils/response"
)

type LyadminButtonSearch struct {
	Name  string `json:"name" form:"name"`   //名称
	Value string `json:"value" form:"value"` //值
	response.StructPageQueryParams
}
