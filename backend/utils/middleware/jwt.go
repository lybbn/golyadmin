package middleware

import (
	"strconv"
	"strings"
	"time"

	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	"gitee.com/lybbn/golyadmin/service"
	"gitee.com/lybbn/golyadmin/utils"
	"gitee.com/lybbn/golyadmin/utils/response"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

// jwt认证头部Authorization格式： JWT xxxxxxx
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.ErrorCodeResponse(4001, "需要认证才能访问！", c)
			c.Abort() //终止退出
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "JWT") {
			response.ErrorResponse("无效的token", c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		token := parts[1]
		if jwtService.IsInBlacklist(token) {
			response.ErrorResponse("您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}
		// VerifyToken 解析token包含的信息
		claims, err := j.VerifyToken(token)
		if err != nil {
			if strings.Contains(err.Error(), jwt.ErrTokenExpired.Error()) {
				response.ErrorCodeResponse(4001, "登录授权已过期", c)
				c.Abort()
				return
			}
			response.ErrorResponse("无效的token,请登录!", c)
			c.Abort()
			return
		}

		// 已登录用户被管理员禁用 需要使该用户的jwt失效 此处比较消耗性能 自行根据项目需要选择是否打开
		user, err := userService.GetUserInfoById(claims.BaseClaims.ID)
		if err != nil {
			// _ = jwtService.JoinBlacklist(system.LyadminJwtBlacklist{Jwt: token})
			global.GL_LOG.Error("jwt查询用户信息失败：" + err.Error())
			response.ErrorResponse(err.Error(), c)
			c.Abort()
			return
		}
		if !user.IsActive {
			// _ = jwtService.JoinBlacklist(system.LyadminJwtBlacklist{Jwt: token})
			response.ErrorResponse("该账号已被禁用，请联系管理员！", c)
			c.Abort()
			return
		}

		// token快过期小于1天（BufferTime），在header设置头部new-token设置新的token
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.GL_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.RefreshTokenByOldToken(token, *claims)
			newClaims, _ := j.VerifyToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			if !global.GL_CONFIG.System.UseMultipoint {
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.GL_LOG.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JoinBlacklist(system.LyadminJwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Set("claims", claims)
		c.Set("golyadmin_userinfo", user)
		c.Next() // 后续的处理函数可以用c.Get("claims")来获取当前请求的用户信息
	}
}
