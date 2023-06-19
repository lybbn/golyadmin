package system

import (
	"context"

	"gitee.com/lybbn/go-vue-lyadmin/global"
	"gitee.com/lybbn/go-vue-lyadmin/model/system"
	"gitee.com/lybbn/go-vue-lyadmin/utils"
)

type JwtService struct{}

// 获取黑名单缓存 key
func (js *JwtService) getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + utils.MD5([]byte(tokenStr))
}

func (js *JwtService) JoinBlacklist(jwtList system.LyadminJwtBlacklist) (err error) {
	err = global.GVLA_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	dr, err := utils.ParseDuration(global.GVLA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.GVLA_REDIS.Set(context.Background(), js.getBlackListKey(jwtList.Jwt), 1, timer).Err()
	return err
}

// IsInBlacklist 判断JWT是否在黑名单中
func (js *JwtService) IsInBlacklist(tokenStr string) bool {
	redisJWT, err := global.GVLA_REDIS.Get(context.Background(), js.getBlackListKey(tokenStr)).Result()
	if redisJWT == "" || err != nil {
		return false
	}
	return true
}

// 从redis获取jwt
func (js *JwtService) GetRedisJWT(username string) (redisJWT string, err error) {
	redisJWT, err = global.GVLA_REDIS.Get(context.Background(), username).Result()
	return redisJWT, err
}

// jwt存入redis缓存并设置过期时间
func (js *JwtService) SetRedisJWT(jwt string, username string) (err error) {
	dr, err := utils.ParseDuration(global.GVLA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.GVLA_REDIS.Set(context.Background(), username, jwt, timer).Err()
	return err
}
