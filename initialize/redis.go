/**
 * @Description 初始化redis
 **/

 package initialize

 import (
    "context"
    "gitee.com/lybbn/go-vue-lyadmin/global"
    "github.com/go-redis/redis/v8"
    "go.uber.org/zap"
    "time"
)

// 初始化redis客户端
func InitRedis() {
    redisCfg := global.GVLA_CONFIG.Redis
    client := redis.NewClient(&redis.Options{
        Addr:     redisCfg.Addr,
        Password: redisCfg.Password,
        DB:       redisCfg.DB,
        DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
        ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
        WriteTimeout: 3 * time.Second, //写超时，默认等于读超时
    })
    pong, err := client.Ping(context.Background()).Result()
    if err != nil {
        global.GVLA_LOG.Error("redis连接错误, err:", zap.Error(err))
    } else {
        global.GVLA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
        global.GVLA_REDIS = client
    }
}