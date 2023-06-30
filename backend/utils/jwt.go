package utils

import (
	"errors"
	"strings"
	"time"

	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims,jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UUID           string
	ID             uint
	Username       string
	Nickname       string
	Identity       int
	DeptId         uint
	RoleIds        []int //角色数组
	RoleDeptIds    []int //角色关联部门数组
	RoleDataScopes []int //数据权限类型数组
}

type JWT struct {
	SecretKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GL_CONFIG.JWT.SecretKey),
	}
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	bf, _ := ParseDuration(global.GL_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.GL_CONFIG.JWT.ExpiresTime)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"lyadmin"},               // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间 7天  配置文件
			IssuedAt:  jwt.NewNumericDate(time.Now()),            //签发时间
			Issuer:    "lybbn",                                   // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建新的 Token (HS256)
func (j *JWT) CreateToken(claims CustomClaims) (token string, err error) {
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return withClaims.SignedString(j.SecretKey)
}

// RefreshTokenByOldToken 旧token 换新token 使用Singleflight避免并发问题
func (j *JWT) RefreshTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	v, err, _ := global.GL_Singleflight.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// VerifyToken 验证 Token
func (j *JWT) VerifyToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	var mc = new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, e error) {
		return j.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

func GetClaims(c *gin.Context) (*CustomClaims, error) {
	authHeader := c.Request.Header.Get("Authorization")
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
	token := parts[1]
	j := NewJWT()
	claims, err := j.VerifyToken(token)
	if err != nil {
		global.GL_LOG.Error("获取请求头Authorization的jwt解析信息失败, 请检查请求头是否存在Authorization且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUser := claims.(*CustomClaims)
		return waitUser.BaseClaims.ID
	}
}

// GetUserIdentity 从Gin的Context中获取从jwt解析出来的用户身份
func GetUserIdentity(c *gin.Context) int {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.Identity
		}
	} else {
		waitUser := claims.(*CustomClaims)
		return waitUser.BaseClaims.Identity
	}
}

// GetDeptID 从Gin的Context中获取从jwt解析出来的用户部门ID
func GetDeptID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.DeptId
		}
	} else {
		waitUser := claims.(*CustomClaims)
		return waitUser.BaseClaims.DeptId
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户信息
func GetUserInfo(c *gin.Context) *CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUser := claims.(*CustomClaims)
		return waitUser
	}
}

// GetUserInfoDB 从Gin的Context中获取从数据库查询的用户信息
func GetUserInfoDB(c *gin.Context) system.LyadminUsers {
	if userinfo, exists := c.Get("golyadmin_userinfo"); !exists {
		var uinfo = system.LyadminUsers{}
		return uinfo
	} else {
		waitUser := userinfo.(system.LyadminUsers)
		return waitUser
	}
}

// GetDeptIdDB 从Gin的Context中获取从数据库查询的用户部门ID
func GetDeptIdDB(c *gin.Context) uint {
	return GetUserInfoDB(c).DeptId
}
