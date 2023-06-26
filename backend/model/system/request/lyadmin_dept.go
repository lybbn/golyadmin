package request

import (
	"gitee.com/lybbn/golyadmin/utils/response"
)

type LyadminDeptSearch struct {
	Search string `json:"search" form:"search"` //搜索关键词
	Name   string `json:"name" form:"name"`     //名称
	Status string `json:"status" form:"status"` //状态
	response.StructPageQueryParams
}
