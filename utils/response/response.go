package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 4000
	SUCCESS = 2000
	MSGERROR = "error"
	MSGSUCCESS = "success"
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func SuccessResponse(data interface{},message string, c *gin.Context) {
	if message == ""{
		message = MSGSUCCESS
	}
	Result(SUCCESS, data, message, c)
}

func DetailResponse(data interface{}, message string, c *gin.Context) {
	if message == ""{
		message = MSGSUCCESS
	}
	Result(SUCCESS, data, message, c)
}

func ErrorResponse(data interface{}, message string, c *gin.Context) {
	if message == ""{
		message = MSGERROR
	}
	Result(ERROR, data, message, c)
}