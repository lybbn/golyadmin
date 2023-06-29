package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"gitee.com/lybbn/golyadmin/utils"

	"gitee.com/lybbn/golyadmin/global"
	"gitee.com/lybbn/golyadmin/model/system"
	"gitee.com/lybbn/golyadmin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var operationLogService = service.ServiceGroupApp.SystemServiceGroup.OperationLogService

var respPool sync.Pool

func init() {
	respPool.New = func() interface{} {
		return make([]byte, 1024)
	}
}

func OperationLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId int
		reMethod := c.Request.Method
		//不记录GET请求，需要记录请注释以下4行代码
		if reMethod == "GET" {
			c.Next()
			return
		}
		if reMethod != http.MethodGet {
			var err error
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				global.GL_LOG.Error("read body from request error:", zap.Error(err))
			} else {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}
		claims, err := utils.GetClaims(c)
		if err != nil {
			userId = 0
		} else {
			if claims.BaseClaims.ID != 0 {
				userId = int(claims.BaseClaims.ID)
			} else {
				userId = 0
			}
		}
		record := system.LyadminOperationLog{
			Ip:     utils.GetRealClientIP(c),
			Method: reMethod,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
			UserID: userId,
		}

		// 上传文件时候 中间件日志进行裁断操作
		if strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			if len(record.Body) > 1024 {
				// 截断
				newBody := respPool.Get().([]byte)
				copy(newBody, record.Body)
				record.Body = string(newBody)
				defer respPool.Put(newBody[:0])
			}
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		record.Msg = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Code = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		if strings.Contains(c.Writer.Header().Get("Pragma"), "public") ||
			strings.Contains(c.Writer.Header().Get("Expires"), "0") ||
			strings.Contains(c.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/force-download") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/octet-stream") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/download") ||
			strings.Contains(c.Writer.Header().Get("Content-Disposition"), "attachment") ||
			strings.Contains(c.Writer.Header().Get("Content-Transfer-Encoding"), "binary") {
			if len(record.Resp) > 1024 {
				// 截断
				newBody := respPool.Get().([]byte)
				copy(newBody, record.Resp)
				record.Resp = string(newBody)
				defer respPool.Put(newBody[:0])
			}
		}

		if err := operationLogService.CreateLyadminOperationLog(record); err != nil {
			global.GL_LOG.Error("create operation record error:", zap.Error(err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
