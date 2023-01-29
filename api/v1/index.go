package v1

import (
	"gitee.com/lybbn/go-vue-lyadmin/api/v1/example"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
