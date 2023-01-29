package example

import (
	"gitee.com/lybbn/go-vue-lyadmin/utils/response"
	"github.com/gin-gonic/gin"
)

type ExampleApi struct{}

// GetExaExample
// @Tags      ExaExample
// @Summary   获取信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     example.ExaExample                                               
// @Success   2000   {object}  response.Response{data=exampleRes.ExaCustomerResponse,msg=string}
// @Router    /example/example [get]
func (e *ExampleApi) GetExaExample(c *gin.Context) {
	response.DetailResponse(nil, "获取成功", c)
}