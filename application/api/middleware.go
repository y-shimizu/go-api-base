package api

import (
	"fmt"
	"time"

	"github.com/y-shimizu/go-api-base/util"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(util.NewAppLogWriter())
}

func AccessLogger() gin.HandlerFunc {
	return AccessLoggerWithWriter(NewAccessLogger())
}

// Access log logging middleware
func AccessLoggerWithWriter(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()
		ua := c.Request.Header.Get("User-Agent")

		header := ""
		for k, v := range c.Request.Header {
			if k == "User-Agent" {
				continue
			}
			header = header + fmt.Sprintf("%s=%s", k, v)
		}

		logger.WithFields(logrus.Fields{
			"status":         statusCode,
			"latency":        latency,
			"cllient_ip":     clientIP,
			"ua":             ua,
			"method":         method,
			"path":           path,
			"request_header": header,
			"request_uri":    c.Request.RequestURI,
		}).Info(comment)
	}
}

// CORS(Cross-Origin Resource Sharing) を制御するための Middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prod以外の場合はデバックしやすいようにすべてのOriginからのアクセスを許可する
		// Allow-Originを*にするとCookieを利用できないためそうしています
		if !util.Env.IsProd() {
			c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
		}

		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "mode, Mode, credentials, Access-Control-Allow-Headers, Content-Type, Cookie, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
