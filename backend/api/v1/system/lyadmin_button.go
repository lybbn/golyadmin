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

type ButtonApi struct{}

// @Tags      Button
// @Summary   获取按钮全部列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  query       systemReq.LyadminButtonSearch false "名称、方法"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "获取按钮全部列表"
// @Router    /system/button/button [get]
func (a *ButtonApi) GetButton(c *gin.Context) {
	var req systemReq.LyadminButtonSearch
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := buttonService.GetLyadminButtonList(req)
	var data []system.LyadminButton
	err = query.Find(&data).Error
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(data, "获取成功", c)
}

// @Tags      Button
// @Summary   分页获取按钮列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  query       systemReq.LyadminButtonSearch false "名称、方法"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "分页获取按钮列表"
// @Router    /system/button/pagelist [get]
func (a *ButtonApi) GetButtonList(c *gin.Context) {
	var pageInfo systemReq.LyadminButtonSearch
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := buttonService.GetLyadminButtonList(pageInfo)
	p := pagination.Page[system.LyadminButton]{}
	p.PaginateQuery(query, c)
	response.PaginateResponse(p.Data, p, "获取成功", c)
}

// @Tags      Button
// @Summary   新增按钮
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body       system.LyadminMenuButton true "名称"
// @Success   200   {object}  response.StructResponse{data=string,msg=string}  "新增按钮"
// @Router    /system/button/button [post]
func (a *ButtonApi) CreateButton(c *gin.Context) {
	var req system.LyadminButton
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	req.CreateBy = utils.GetUserID(c)
	req.BelongDept = utils.GetDeptID(c)
	err = buttonService.CreateButton(req)
	if err != nil {
		global.GL_LOG.Error("添加失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(nil, "添加成功", c)
}

// @Tags      Button
// @Summary   根据ID删除按钮
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Id      true  "LyadminMenuButton模型"
// @Success   200   {object}  response.StructResponse{msg=string}  "删除LyadminMenuButton"
// @Router    /system/button/button [delete]
func (a *ButtonApi) DeleteButton(c *gin.Context) {
	var req request.Id
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	err = buttonService.DeleteButton(uint(req.Id))
	if err != nil {
		global.GL_LOG.Error("删除失败!", zap.Error(err))
		response.ErrorResponse("删除失败", c)
		return
	}
	response.SuccessResponse(nil, "删除成功", c)
}

// @Tags      Button
// @Summary   编辑按钮
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body       system.LyadminMenuButton true "名称"
// @Success   200   {object}  response.StructResponse{data=string,msg=string}  "编辑按钮"
// @Router    /system/button/button/:id [put]
func (a *ButtonApi) UpdateButton(c *gin.Context) {
	var req system.LyadminButton
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	req.UpdateBy = utils.GetUserID(c)
	err = buttonService.UpdateButton(req)
	if err != nil {
		global.GL_LOG.Error("添加失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(nil, "添加成功", c)
}
