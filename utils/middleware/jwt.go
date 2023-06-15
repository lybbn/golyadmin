package middleware

import (
	"strings"

	"gitee.com/lybbn/go-vue-lyadmin/utils"
	"gitee.com/lybbn/go-vue-lyadmin/utils/response"

	"github.com/gin-gonic/gin"
)

// jwt认证头部Authorization格式： JWT xxxxxxx
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.ErrorCodeResponse(4001, "未登录或非法访问", c)
			c.Abort() //终止退出
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "JWT") {
			response.ErrorResponse("访问失败,无效的token,请登录!", c)
			c.Abort()
			return
		}
		// if jwtService.IsBlacklist(token) {
		// 	response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
		// 	c.Abort()
		// 	return
		// }
		j := utils.NewJWT()
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		token := parts[1]
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			response.ErrorResponse("访问失败,无效的token,请登录!", c)
			c.Abort()
			return
		}

		// 已登录用户被管理员禁用 需要使该用户的jwt失效 此处比较消耗性能 如果需要 请自行打开
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开

		//if user, err := userService.FindUserByUuid(claims.UUID.String()); err != nil || user.Enable == 2 {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}

		// if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
		// 	dr, _ := utils.ParseDuration(global.GVLA_CONFIG.JWT.ExpiresTime)
		// 	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
		// 	newToken, _ := j.CreateTokenByOldToken(token, *claims)
		// 	newClaims, _ := j.ParseToken(newToken)
		// 	c.Header("new-token", newToken)
		// 	c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
		// 	if global.GVLA_CONFIG.System.UseMultipoint {
		// 		RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
		// 		if err != nil {
		// 			global.GVLA_LOG.Error("get redis jwt failed", zap.Error(err))
		// 		} else { // 当之前的取成功时才进行拉黑操作
		// 			_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
		// 		}
		// 		// 无论如何都要记录当前的活跃状态
		// 		_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
		// 	}
		// }
		c.Set("claims", claims)
		c.Next() // 后续的处理函数可以用过c.Get("claims")来获取当前请求的用户信息
	}
}
