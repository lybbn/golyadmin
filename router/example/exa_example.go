package example

import (
	v1 "gitee.com/lybbn/go-vue-lyadmin/api/v1"
	"github.com/gin-gonic/gin"
)

type ExampleRouter struct{}

func (e *ExampleRouter) InitExampleRouter(Router *gin.RouterGroup) {
	customerRouterWithoutRecord := Router.Group("example")
	exaExampleApi := v1.ApiGroupApp.ExampleApiGroup.ExampleApi
	{
		customerRouterWithoutRecord.GET("example", exaExampleApi.GetExaExample)
	}
}
