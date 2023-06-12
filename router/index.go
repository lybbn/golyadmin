package router

import (
	"gitee.com/lybbn/go-vue-lyadmin/router/example"
	"gitee.com/lybbn/go-vue-lyadmin/router/system"
)

type RouterGroup struct {
	Example example.RouterGroup
	System  system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
