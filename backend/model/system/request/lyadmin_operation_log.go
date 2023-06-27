package request

import (
	"gitee.com/lybbn/golyadmin/model/system"
	"gitee.com/lybbn/golyadmin/utils/response"
)

type LyadminOperationLogSearch struct {
	Search  string `json:"search" form:"search"`   //搜索关键词
	BeginAt string `json:"beginAt" form:"beginAt"` //搜索开始时间
	EndAt   string `json:"endAt" form:"endAt"`     //搜索结束时间
	system.LyadminOperationLog
	response.StructPageQueryParams
}
