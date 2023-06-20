package system

import (
	"gitee.com/lybbn/golyadmin/model/system"
	systemReq "gitee.com/lybbn/golyadmin/model/system/request"
	"gitee.com/lybbn/golyadmin/utils/pagination"
	"gitee.com/lybbn/golyadmin/utils/response"
	"github.com/gin-gonic/gin"
)

type MenuApi struct{}

// @Tags      Menu
// @Summary   获取菜单列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  query       systemReq.LyadminMenuSearch true "名称、是否显示、状态、路径"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "分页获取菜单列表"
// @Router    /menu/menu [get]
func (a *MenuApi) GetMenuList(c *gin.Context) {
	var pageInfo systemReq.LyadminMenuSearch
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := menuService.GetLyadminMenuList(pageInfo)
	p := pagination.Page[system.LyadminMenu]{}
	p.PaginateQuery(query, c)
	response.PaginateResponse(p.Data, p, "获取成功", c)
}
