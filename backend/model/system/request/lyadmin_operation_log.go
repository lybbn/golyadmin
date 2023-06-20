package request

import (
	"gitee.com/lybbn/golyadmin/model/system"
	"gitee.com/lybbn/golyadmin/utils/response"
)

type LyadminOperationLogSearch struct {
	system.LyadminOperationLog
	response.StructPageQueryParams
}
