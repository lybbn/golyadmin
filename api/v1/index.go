package v1

import (
	"gitee.com/lybbn/go-vue-lyadmin/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
