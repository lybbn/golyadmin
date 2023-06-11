package v1

import (
	"gitee.com/lybbn/go-vue-lyadmin/api/v1/example"
	"gitee.com/lybbn/go-vue-lyadmin/api/v1/system"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
	SystemApiGroup  system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
