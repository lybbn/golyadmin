package example

import (
	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/utils/pagination"
	"gitee.com/lybbn/go-vue-lyadmin/utils/response"
	"github.com/gin-gonic/gin"
)

type ExampleApi struct {
}

type ExampleService struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

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
	// 详情不带分页
	// var result []ExampleService
	// global.GVLA_DB.Table("lyadmin_users").Select("id", "name", "username").Scan(&result)
	// response.DetailResponse(result, "获取成功", c)

	//分页
	query := global.GVLA_DB.Table("lyadmin_users").Select("id", "name", "username")
	p := pagination.Page[ExampleService]{}
	p.PaginateQuery(query, c)
	println(p.PageSize)
	response.PaginateResponse(p.Data, p, "", c)
}
