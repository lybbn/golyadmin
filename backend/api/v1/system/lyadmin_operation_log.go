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

type OperationLogApi struct{}

// @Tags      OperationLog
// @Summary   根据ID删除日志
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Id      true  "LyadminOperationLog模型"
// @Success   200   {object}  response.StructResponse{msg=string}  "删除LyadminOperationLog"
// @Router    /system/operationlog/log [delete]
func (s *OperationLogApi) DeleteLyadminOperationLog(c *gin.Context) {
	var req request.Id
	err := c.ShouldBind(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	err = operationLogService.DeleteLyadminOperationLog(uint(req.Id))
	if err != nil {
		global.GL_LOG.Error("删除失败!", zap.Error(err))
		response.ErrorResponse("删除失败", c)
		return
	}
	response.SuccessResponse(nil, "删除成功", c)
}

// @Tags      OperationLog
// @Summary   清除全部日志
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Empty      false  "无"
// @Success   200   {object}  response.StructResponse{msg=string}  "清除全部日志"
// @Router    /system/operationlog/log [delete]
func (s *OperationLogApi) DeleteAllLyadminOperationLog(c *gin.Context) {
	identity := utils.GetUserIdentity(c)
	if identity != 1 {
		response.ErrorResponse("该账号权限不足！", c)
		return
	}
	err := operationLogService.DeleteAllLyadminOperationLog()
	if err != nil {
		global.GL_LOG.Error("清空失败", zap.Error(err))
		response.ErrorResponse("清空失败", c)
		return
	}
	response.SuccessResponse(nil, "清空成功", c)
}

// @Tags      OperationLog
// @Summary   批量删除日志
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Ids                 true  "批量删除日志"
// @Success   200   {object}  response.StructResponse{msg=string}  "批量删除日志"
// @Router    /system/operationlog/deletelogbyids [delete]
func (s *OperationLogApi) DeleteLyadminOperationLogByIds(c *gin.Context) {
	var IDS request.Ids
	err := c.ShouldBind(&IDS)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	err = operationLogService.DeleteLyadminOperationLogByIds(IDS)
	if err != nil {
		global.GL_LOG.Error("批量删除失败!", zap.Error(err))
		response.ErrorResponse("批量删除失败", c)
		return
	}
	response.SuccessResponse(nil, "批量删除成功", c)
}

// @Tags      OperationLog
// @Summary   用id查询OperationLog
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.Id                                  true  "Id"
// @Success   200   {object}  response.StructResponse{data=system.LyadminOperationLog,msg=string}  "用id查询LyadminOperationLog"
// @Router    /system/operationlog/log [get]
func (s *OperationLogApi) GetLyadminOperationLogDetail(c *gin.Context) {
	var req request.Id
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.ErrorResponse(utils.GetValidMsg(err, &req), c)
		return
	}
	data, err := operationLogService.GetLyadminOperationLogDetail(uint(req.Id))
	if err != nil {
		global.GL_LOG.Error("查询失败!", zap.Error(err))
		response.ErrorResponse("查询失败", c)
		return
	}
	response.SuccessResponse(data, "查询成功", c)
}

// @Tags      OperationLog
// @Summary   分页获取LyadminOperationLog列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     systemReq.LyadminOperationLogSearch                        true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  response.StructResponse{data=map[string]interface{},msg=string}  "分页获取LyadminOperationLog列表,返回包括列表,总数,页码,每页数量"
// @Router    /system/operationlog/pagelist [get]
func (s *OperationLogApi) GetLyadminOperationLogList(c *gin.Context) {
	var pageInfo systemReq.LyadminOperationLogSearch
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		response.ErrorResponse(err.Error(), c)
		return
	}
	query := operationLogService.GetLyadminOperationLogList(pageInfo)
	p := pagination.Page[system.LyadminOperationLog]{}
	p.PaginateQuery(query, c)
	response.PaginateResponse(p.Data, p, "获取成功", c)
}
