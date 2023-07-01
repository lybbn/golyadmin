package system

import (
	"strings"

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

type RoleApi struct{}

// @Tags      Role
// @Summary   获取角色全部列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  query       systemReq.LyadminRoleSearch false "搜索字段"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "获取角色全部列表"
// @Router    /system/role/menu [get]
func (r *RoleApi) GetRole(c *gin.Context) {
	var req systemReq.LyadminRoleSearch
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := roleService.GetLyadminRole(req).Scopes(utils.DataLevelPermissionsFilter(system.LyadminRole{}, c))
	var data []system.LyadminRole
	err = query.Find(&data).Error
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(data, "获取成功", c)
}

// @Tags      Role
// @Summary   分页获取角色列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  query       systemReq.LyadminRoleSearch false "分页参数"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "分页获取角色列表"
// @Router    /system/role/menulist [get]
func (r *RoleApi) GetRoleList(c *gin.Context) {
	var pageInfo systemReq.LyadminRoleSearch
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := roleService.GetLyadminRoleList(pageInfo).Scopes(utils.DataLevelPermissionsFilter(system.LyadminRole{}, c))
	p := pagination.Page[system.LyadminRole]{}
	p.PaginateQuery(query, c)
	response.PaginateResponse(p.Data, p, "获取成功", c)
}

// @Tags      Role
// @Summary   通过角色id获取该角色所有的菜单
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "通过角色id获取该角色所有的菜单"
// @Router    /system/role/role_id_to_menu/:id [get]
func (r *RoleApi) GetRoleMenuById(c *gin.Context) {
	data, err := roleService.GetRoleMenuById()
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(data, "获取成功", c)
}

// @Tags      Role
// @Summary   新增角色
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body       system.LyadminRole true "LyadminRole模型"
// @Success   200   {object}  response.StructResponse{data=string,msg=string}  "新增角色"
// @Router    /system/role/role [post]
func (r *RoleApi) CreateRole(c *gin.Context) {
	var req system.LyadminRole
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	req.CreateBy = utils.GetUserID(c)
	req.BelongDept = utils.GetDeptIdDB(c)
	err = roleService.CreateRole(req)
	if err != nil {
		global.GL_LOG.Error("添加失败!", zap.Error(err))
		msg := err.Error()
		if strings.Contains(err.Error(), "Error 1062 (23000): Duplicate entry") {
			msg = "存在相同权限字符，请更换！"
		}
		response.ErrorResponse(msg, c)
		return
	}
	response.SuccessResponse(nil, "添加成功", c)
}

// @Tags      Role
// @Summary   根据ID删除角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Id      true  "ID"
// @Success   200   {object}  response.StructResponse{msg=string}  "删除角色"
// @Router    /system/role/role [delete]
func (r *RoleApi) DeleteRole(c *gin.Context) {
	var req request.Id
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	err = roleService.DeleteRole(uint(req.Id))
	if err != nil {
		global.GL_LOG.Error("删除失败!", zap.Error(err))
		response.ErrorResponse("删除失败", c)
		return
	}
	response.SuccessResponse(nil, "删除成功", c)
}

// @Tags      Role
// @Summary   编辑角色
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body       system.LyadminRole true "分页"
// @Success   200   {object}  response.StructResponse{data=string,msg=string}  "编辑菜单"
// @Router    /system/role/role/:id [put]
func (r *RoleApi) UpdateRole(c *gin.Context) {
	var req system.LyadminRole
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	req.UpdateBy = utils.GetUserID(c)
	err = roleService.UpdateRole(req)
	if err != nil {
		global.GL_LOG.Error("修改失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(nil, "修改成功", c)
}

// @Tags      Role
// @Summary   更新角色权限
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      systemReq.LyadminRoleParams true "原LyadminRoleParams参数"
// @Success   200   {object}  response.StructResponse{data=string,msg=string}  "更新角色权限"
// @Router    /system/role/permission/:id [put]
func (r *RoleApi) UpdateRolePremission(c *gin.Context) {
	var req systemReq.LyadminRoleParams
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	err = roleService.UpdateRolePremission(req)
	if err != nil {
		global.GL_LOG.Error("修改失败!", zap.Error(err))
		response.ErrorResponse(err.Error(), c)
		return
	}
	response.SuccessResponse(nil, "修改成功", c)
}
