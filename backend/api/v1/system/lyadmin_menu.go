package system

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	systemReq "gitee.com/lybbn/golyadmin/model/system/request"
	"gitee.com/lybbn/golyadmin/utils"
	"gitee.com/lybbn/golyadmin/utils/pagination"
	"gitee.com/lybbn/golyadmin/utils/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenuApi struct{}

// @Tags      Menu
// @Summary   获取菜单全部列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  query       systemReq.LyadminMenuSearch true "名称、是否显示、状态、路径"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "获取菜单全部列表"
// @Router    /menu/menu [get]
func (a *MenuApi) GetMenu(c *gin.Context) {
	var req systemReq.LyadminMenuSearch
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := menuService.GetLyadminMenuList(req)
	var data []system.LyadminMenu
	err = query.Find(&data).Error
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(data, "获取成功", c)
}

// @Tags      Menu
// @Summary   分页获取菜单列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  query       systemReq.LyadminMenuSearch true "名称、是否显示、状态、路径"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "分页获取菜单列表"
// @Router    /menu/menulist [get]
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

// @Tags      Menu
// @Summary   新增菜单
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body       system.LyadminMenu true "名称"
// @Success   200   {object}  response.StructResponse{data=string,msg=string}  "新增菜单"
// @Router    /menu/menu [post]
func (a *MenuApi) CreateMenu(c *gin.Context) {
	var req system.LyadminMenu
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	err = menuService.CreateMenu(req)
	if err != nil {
		global.GL_LOG.Error("添加失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(nil, "添加成功", c)
}

// @Tags      Menu
// @Summary   编辑菜单
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body       system.LyadminMenu true "名称"
// @Success   200   {object}  response.StructResponse{data=string,msg=string}  "编辑菜单"
// @Router    /menu/menu/:id [put]
func (a *MenuApi) UpdateMenu(c *gin.Context) {
	var req system.LyadminMenu
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	err = menuService.UpdateMenu(req)
	if err != nil {
		global.GL_LOG.Error("添加失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(nil, "添加成功", c)
}
