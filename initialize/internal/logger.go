package internal

import (
	"fmt"

	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.GVLA_CONFIG.System.DbType {
	case "mysql":
		logZap = global.GVLA_CONFIG.Mysql.LogZap
	case "pgsql":
		logZap = global.GVLA_CONFIG.Pgsql.LogZap
	}
	if logZap {
		global.GVLA_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
