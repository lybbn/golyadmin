package router

import (
	"gitee.com/lybbn/go-vue-lyadmin/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
