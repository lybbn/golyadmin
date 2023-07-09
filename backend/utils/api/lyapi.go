// 数据库操作工具类
package api

import (
	"gitee.com/lybbn/golyadmin/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DefaultPrimaryKeyName = "id"
)

type LyApi struct {
	Context *gin.Context
	Orm     *gorm.DB
	Errors  error
}

// SetContext 设置gin http上下文Context
func (a *LyApi) SetContext(c *gin.Context) *LyApi {
	a.Context = c
	return a
}

// SetOrm 自定义设置*gorm.DB数据库连接
func (a *LyApi) SetOrm(db *gorm.DB) *LyApi {
	a.Orm = db
	return a
}

// DefaultOrm 设置默认*gorm.DB数据库连接
func (a *LyApi) SetDefaultOrm() *LyApi {
	a.Orm = global.GL_DB
	return a
}
