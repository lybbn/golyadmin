package example

import (
	"fmt"

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
// @Success 2000 {object} response.StructResponse{data=ExampleService}
// @Router    /example/example [get]
func (e *ExampleApi) GetExaExample(c *gin.Context) {
	// 详情不带分页
	var result []ExampleService
	global.GVLA_DB.Table("lyadmin_users").Select("id", "name", "username").Scan(&result)
	response.SuccessResponse(result, "获取成功", c)
}

// 请求参数结构
type ExampleQueryParmas struct {
	response.StructPageQueryParams
	Name string `json:"name" form:"name"` //查询参数
}

// GetExaExampleList
// @Tags      ExaExample
// @Summary   分页获取信息列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     ExampleQueryParmas                                        true  "页码, 每页大小"
// @Success 2000 {object} response.StructPageResponse{data=ExampleService}
// @Router    /example/exampleList [get]
func (e *ExampleApi) GetExaExampleList(c *gin.Context) {
	//接收请求参数
	var pageParams ExampleQueryParmas
	err := c.ShouldBindQuery(&pageParams)
	fmt.Println("====== Only Bind By Query String ======")
	fmt.Println(pageParams)
	fmt.Println(pageParams.Name)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}

	query := global.GVLA_DB.Table("lyadmin_users").Select("id", "name", "username")
	p := pagination.Page[ExampleService]{}
	p.PaginateQuery(query, c)
	response.PaginateResponse(p.Data, p, "获取成功", c)
}
