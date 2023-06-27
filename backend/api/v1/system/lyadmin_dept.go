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

type DeptApi struct{}

// @Tags      Dept
// @Summary   获取部门全部列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  query       systemReq.LyadminDeptSearch false "名称、方法"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "获取部门全部列表"
// @Router    /system/menu_button/menu_button [get]
func (a *DeptApi) GetDept(c *gin.Context) {
	var req systemReq.LyadminDeptSearch
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := deptService.GetLyadminDeptList(req)
	var data []system.LyadminDept
	err = query.Find(&data).Error
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(data, "获取成功", c)
}

// @Tags      Dept
// @Summary   分页获取部门列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  query       systemReq.LyadminDeptSearch false "名称"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "分页获取部门列表"
// @Router    /system/menu_button/pagelist [get]
func (a *DeptApi) GetDeptList(c *gin.Context) {
	var pageInfo systemReq.LyadminDeptSearch
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := deptService.GetLyadminDeptList(pageInfo)
	p := pagination.Page[system.LyadminDept]{}
	p.PaginateQuery(query, c)
	response.PaginateResponse(p.Data, p, "获取成功", c)
}

// @Tags      Dept
// @Summary   新增部门
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body       system.LyadminDept true "名称"
// @Success   200   {object}  response.StructResponse{data=string,msg=string}  "新增部门"
// @Router    /system/menu_button/menu_button [post]
func (a *DeptApi) CreateDept(c *gin.Context) {
	var req system.LyadminDept
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	req.CreateBy = utils.GetUserID(c)
	req.BelongDept = utils.GetDeptID(c)
	err = deptService.CreateDept(req)
	if err != nil {
		global.GL_LOG.Error("添加失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(nil, "添加成功", c)
}

// @Tags      Dept
// @Summary   根据ID删除部门
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Id      true  "LyadminMenuButton模型"
// @Success   200   {object}  response.StructResponse{msg=string}  "删除LyadminMenuButton"
// @Router    /system/menu_button/menu_button [delete]
func (a *DeptApi) DeleteDept(c *gin.Context) {
	var req request.Id
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	err = deptService.DeleteDept(uint(req.Id))
	if err != nil {
		global.GL_LOG.Error("删除失败!", zap.Error(err))
		response.ErrorResponse("删除失败", c)
		return
	}
	response.SuccessResponse(nil, "删除成功", c)
}

// @Tags      Dept
// @Summary   编辑部门
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body       system.LyadminDept true "名称"
// @Success   200   {object}  response.StructResponse{data=string,msg=string}  "编辑部门"
// @Router    /system/menu_button/menu_button/:id [put]
func (a *DeptApi) UpdateDept(c *gin.Context) {
	var req system.LyadminDept
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	req.UpdateBy = utils.GetUserID(c)
	err = deptService.UpdateDept(req)
	if err != nil {
		global.GL_LOG.Error("修改失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(nil, "修改成功", c)
}
