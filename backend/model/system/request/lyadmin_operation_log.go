package request

import (
	"gitee.com/lybbn/golyadmin/utils/response"
)

type LyadminOperationLogSearch struct {
	Search  string `json:"search" form:"search"`                                                   //搜索关键词
	BeginAt string `json:"beginAt" form:"beginAt"`                                                 //搜索开始时间
	EndAt   string `json:"endAt" form:"endAt"`                                                     //搜索结束时间
	Ip      string `json:"ip" form:"ip" gorm:"type:varchar(50);column:ip;comment:请求ip"`            // 请求ip
	Method  string `json:"method" form:"method" gorm:"type:varchar(8);column:method;comment:请求方法"` // 请求方法
	Path    string `json:"path" form:"path" gorm:"column:path;comment:请求路径"`                       // 请求路径
	Code    int    `json:"code" form:"code" gorm:"type:varchar(32);column:code;comment:请求状态"`      // 请求状态
	Agent   string `json:"agent" form:"agent" gorm:"column:agent;comment:UserAgent代理"`             // UserAgent代理
	Msg     string `json:"msg" form:"msg" gorm:"column:msg;comment:返回信息"`                          // 错误信息
	Body    string `json:"body" form:"body" gorm:"type:text;column:body;comment:请求Body"`           // 请求Body
	Resp    string `json:"resp" form:"resp" gorm:"type:text;column:resp;comment:响应Body"`           // 响应Body
	response.StructPageQueryParams
}
