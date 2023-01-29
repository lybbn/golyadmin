package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR      = 4000
	SUCCESS    = 2000
	MSGERROR   = "error"
	MSGSUCCESS = "success"
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

/**
 * 分页
 * query 为查询要分页的记录
 * 需要传入当前页码和每页条数
 */
func PageResponse(current int64, limit int64, query interface{}, message string, c *gin.Context) {
	if message == "" {
		message = MSGSUCCESS
	}
	page := NewPage(current, limit)
	var count int64 // 统计总的记录数
	query.Count(&count)
	if count > 0 {
		result := query.Limit(int(page.GetSize())).Offset(int(page.GetOffset())).Find(&sysUserList)
		// 返回 error
		if result.Error != nil {
			fmt.Println("数据分页查询异常：", result.Error)
			return nil, result.Error
		}
	}

	page.SetTotal(count)
	page.SetRecords(sysUserList)
	Result(SUCCESS, data, message, c)
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
