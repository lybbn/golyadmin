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
// @Summary   不分页获取信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Router    /api/v1/example/example [get]
func (e *ExampleApi) GetExaExample(c *gin.Context) {
	// 详情不带分页
	var result []ExampleService
	global.GVLA_DB.Table("lyadmin_users").Select("id", "name", "username").Scan(&result)
	response.SuccessResponse(result, "获取成功", c)
}

// GetExaExampleList
// @Tags      ExaExample
// @Summary   分页获取信息列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Router    /api/v1/example/exampleList [get]
func (e *ExampleApi) GetExaExampleList(c *gin.Context) {
	query := global.GVLA_DB.Table("lyadmin_users").Select("id", "name", "username")
	p := pagination.Page[ExampleService]{}
	p.PaginateQuery(query, c)
	response.PaginateResponse(p.Data, p, "获取成功", c)
}
