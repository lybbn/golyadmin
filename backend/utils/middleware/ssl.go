package middleware

import (
	"fmt"

	"gitee.com/lybbn/golyadmin/global"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// 用https时把这个中间件在router里面use一下就好

func LoadSSL() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     global.GL_CONFIG.Ssl.Domain,
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Next()
	}
}
