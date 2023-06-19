package request

import (
	"gitee.com/lybbn/go-vue-lyadmin/model/system"
	"gitee.com/lybbn/go-vue-lyadmin/utils/response"
)

type LyadminOperationLogSearch struct {
	system.LyadminOperationLog
	response.StructPageQueryParams
}
