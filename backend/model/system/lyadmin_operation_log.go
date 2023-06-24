package system

import (
	"time"

	"gitee.com/lybbn/golyadmin/global"
)

type LyadminOperationLog struct {
	global.GL_BASE_MODEL
	Ip      string        `json:"ip" form:"ip" gorm:"type:varchar(50);column:ip;comment:请求ip"`                  // 请求ip
	Method  string        `json:"method" form:"method" gorm:"type:varchar(8);column:method;comment:请求方法"`       // 请求方法
	Path    string        `json:"path" form:"path" gorm:"column:path;comment:请求路径"`                             // 请求路径
	Code    int           `json:"code" form:"code" gorm:"type:varchar(32);column:code;comment:请求状态"`            // 请求状态
	Latency time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:延迟" swaggertype:"string"` // 延迟
	Agent   string        `json:"agent" form:"agent" gorm:"column:agent;comment:UserAgent代理"`                   // UserAgent代理
	Msg     string        `json:"msg" form:"msg" gorm:"column:msg;comment:返回信息"`                                // 错误信息
	Body    string        `json:"body" form:"body" gorm:"type:text;column:body;comment:请求Body"`                 // 请求Body
	Resp    string        `json:"resp" form:"resp" gorm:"type:text;column:resp;comment:响应Body"`                 // 响应Body
	UserID  int           `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`                    // 用户id Belongs To
	User    LyadminUsers  `json:"user"`
	global.GL_CONTROL_MODEL
}

func (LyadminOperationLog) TableName() string {
	return "lyadmin_operation_log"
}
