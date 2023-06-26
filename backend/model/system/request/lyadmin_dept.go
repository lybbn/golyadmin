package request

import (
	"gitee.com/lybbn/golyadmin/utils/response"
)

type LyadminDeptSearch struct {
	Name string `json:"name" form:"name"` //名称
	response.StructPageQueryParams
}
