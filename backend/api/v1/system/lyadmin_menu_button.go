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

type MenuButtonApi struct{}

// @Tags      MenuButton
// @Summary   获取菜单按钮全部列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  query       systemReq.LyadminMenuButtonSearch false "名称、方法"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "获取菜单按钮全部列表"
// @Router    /system/menu_button/menu_button [get]
func (a *MenuButtonApi) GetMenuButton(c *gin.Context) {
	var req systemReq.LyadminMenuButtonSearch
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := menuButtonService.GetLyadminMenuButtonList(req)
	var data []system.LyadminMenuButton
	err = query.Find(&data).Error
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(data, "获取成功", c)
}

// @Tags      MenuButton
// @Summary   分页获取菜单按钮列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  query       systemReq.LyadminMenuButtonSearch false "名称、方法"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "分页获取菜单按钮列表"
// @Router    /system/menu_button/pagelist [get]
func (a *MenuButtonApi) GetMenuButtonList(c *gin.Context) {
	var pageInfo systemReq.LyadminMenuButtonSearch
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := menuButtonService.GetLyadminMenuButtonList(pageInfo)
	p := pagination.Page[system.LyadminMenuButton]{}
	p.PaginateQuery(query, c)
	response.PaginateResponse(p.Data, p, "获取成功", c)
}

// @Tags      MenuButton
// @Summary   新增菜单按钮
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body       system.LyadminMenuButton true "名称"
// @Success   200   {object}  response.StructResponse{data=string,msg=string}  "新增菜单按钮"
// @Router    /system/menu_button/menu_button [post]
func (a *MenuButtonApi) CreateMenuButton(c *gin.Context) {
	var req system.LyadminMenuButton
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	req.CreateBy = utils.GetUserID(c)
	req.BelongDept = utils.GetDeptIdDB(c)
	err = menuButtonService.CreateMenuButton(req)
	if err != nil {
		global.GL_LOG.Error("添加失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(nil, "添加成功", c)
}

// @Tags      MenuButton
// @Summary   根据ID删除菜单按钮
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Id      true  "LyadminMenuButton模型"
// @Success   200   {object}  response.StructResponse{msg=string}  "删除LyadminMenuButton"
// @Router    /system/menu_button/menu_button [delete]
func (a *MenuButtonApi) DeleteMenuButton(c *gin.Context) {
	var req request.Id
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	err = menuButtonService.DeleteMenuButton(uint(req.Id))
	if err != nil {
		global.GL_LOG.Error("删除失败!", zap.Error(err))
		response.ErrorResponse("删除失败", c)
		return
	}
	response.SuccessResponse(nil, "删除成功", c)
}

// @Tags      MenuButton
// @Summary   编辑菜单按钮
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body       system.LyadminMenuButton true "名称"
// @Success   200   {object}  response.StructResponse{data=string,msg=string}  "编辑菜单按钮"
// @Router    /system/menu_button/menu_button/:id [put]
func (a *MenuButtonApi) UpdateMenuButton(c *gin.Context) {
	var req system.LyadminMenuButton
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	req.UpdateBy = utils.GetUserID(c)
	err = menuButtonService.UpdateMenuButton(req)
	if err != nil {
		global.GL_LOG.Error("修改失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(nil, "修改成功", c)
}
