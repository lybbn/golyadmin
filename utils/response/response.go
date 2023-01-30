package response

import (
	"net/http"

	"gitee.com/lybbn/go-vue-lyadmin/utils/pagination"
	"github.com/gin-gonic/gin"
)

type StructResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type StructPageResponse struct {
	Code        int         `json:"code"`
	Msg         string      `json:"msg"`
	CurrentPage int64       `json:"page"`  // 当前页码
	PageSize    int64       `json:"limit"` // 每页条数
	Total       int64       `json:"total"` // 总数据量
	Pages       int64       `json:"pages"` // 总分页数
	Data        interface{} `json:"data"`  // 分页数据
}

const (
	ERROR      = 4000
	SUCCESS    = 2000
	MSGERROR   = "error"
	MSGSUCCESS = "success"
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, StructResponse{
		code,
		data,
		msg,
	})
}

// 分页
func PaginateResponse[T any](data interface{}, page pagination.Page[T], message string, c *gin.Context) {
	if message == "" {
		message = MSGSUCCESS
	}
	var p StructPageResponse
	p.Code = SUCCESS
	p.Msg = message
	p.Data = data
	p.CurrentPage = page.CurrentPage
	p.PageSize = page.PageSize
	p.Total = page.Total
	p.Pages = page.Pages
	c.JSON(http.StatusOK, p)
}

// 详情
func DetailResponse(data interface{}, message string, c *gin.Context) {
	if message == "" {
		message = MSGSUCCESS
	}
	Result(SUCCESS, data, message, c)
}

// 错误
func ErrorResponse(data interface{}, message string, c *gin.Context) {
	if message == "" {
		message = MSGERROR
	}
	Result(ERROR, data, message, c)
}
