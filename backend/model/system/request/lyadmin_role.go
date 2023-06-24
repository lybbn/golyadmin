package request

import (
	"gitee.com/lybbn/golyadmin/utils/response"
)

type LyadminRoleSearch struct {
	Name   string `json:"name" form:"name"`     //名称
	Status *bool  `json:"status" form:"status"` //状态
	response.StructPageQueryParams
}
