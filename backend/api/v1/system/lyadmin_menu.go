package system

import (
	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/common/request"
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
// @Router    /system/menu/menu [get]
func (a *MenuApi) GetMenu(c *gin.Context) {
	var req systemReq.LyadminMenuSearch
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := menuService.GetLyadminMenuList(req).Scopes(utils.DataLevelPermissionsFilter(system.LyadminMenu{}, c))
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
// @Router    /system/menu/menulist [get]
func (a *MenuApi) GetMenuList(c *gin.Context) {
	var pageInfo systemReq.LyadminMenuSearch
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := menuService.GetLyadminMenuList(pageInfo).Scopes(utils.DataLevelPermissionsFilter(system.LyadminMenu{}, c))
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
// @Router    /system/menu/menu [post]
func (a *MenuApi) CreateMenu(c *gin.Context) {
	var req system.LyadminMenu
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	req.CreateBy = utils.GetUserID(c)
	req.BelongDept = utils.GetDeptIdDB(c)
	err = menuService.CreateMenu(req)
	if err != nil {
		global.GL_LOG.Error("添加失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(nil, "添加成功", c)
}

// @Tags      Menu
// @Summary   根据ID删除菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Id      true  "LyadminMenu模型"
// @Success   200   {object}  response.StructResponse{msg=string}  "删除LyadminMenu"
// @Router    /system/menu/menu [delete]
func (a *MenuApi) DeleteMenu(c *gin.Context) {
	var req request.Id
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	err = menuService.DeleteMenu(uint(req.Id))
	if err != nil {
		global.GL_LOG.Error("删除失败!", zap.Error(err))
		response.ErrorResponse("删除失败", c)
		return
	}
	response.SuccessResponse(nil, "删除成功", c)
}

// @Tags      Menu
// @Summary   编辑菜单
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body       system.LyadminMenu true "名称"
// @Success   200   {object}  response.StructResponse{data=string,msg=string}  "编辑菜单"
// @Router    /system/menu/menu/:id [put]
func (a *MenuApi) UpdateMenu(c *gin.Context) {
	var req system.LyadminMenu
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	req.UpdateBy = utils.GetUserID(c)
	err = menuService.UpdateMenu(req)
	if err != nil {
		global.GL_LOG.Error("修改失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(nil, "修改成功", c)
}

// @Tags      Menu
// @Summary   获取菜单全部列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  query       systemReq.LyadminMenuSearch true "名称、是否显示、状态、路径"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "获取菜单全部列表"
// @Router    /system/menu/web_router [get]
func (a *MenuApi) GetWebRouter(c *gin.Context) {
	uinfo := utils.GetUserInfo(c)
	menus, err := menuService.GetWebRouter(uinfo)
	if err != nil {
		global.GL_LOG.Error("获取失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(menus, "获取成功", c)
}
