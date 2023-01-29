package example

import (
	"fmt"

	"gitee.com/lybbn/go-vue-lyadmin/global"
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
	rows, _ := global.GVLA_DB.Raw("select * from lyadmin_users").Rows()
	var (
		_id       string
		_name     string
		_username string
	)
	for rows.Next() {
		fmt.Printf("lyadmin_users -> id=%v,name=%v,username=%v", _id, _name, _username)
	}
	response.DetailResponse(nil, "获取成功", c)
}
