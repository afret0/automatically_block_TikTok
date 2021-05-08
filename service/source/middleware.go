package source

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// TODO 测试
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if len(token) < 1 {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "未携带 token, 请先登录"})
			c.Abort()
			return
		}
		claims, err := GetJWT().ParseToken(token)
		if err != nil {
			logger.Errorln(token, err)
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startT := time.Now()
		c.Next()
		endT := time.Now()
		latencyT := endT.Sub(startT)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		clientIP := c.ClientIP()
		middlewareLogger.WithFields(logrus.Fields{
			"reqTime":  startT.Format("2006-01-02 15:04:05"),
			"latencyT": latencyT,
			"method":   reqMethod,
			"uri":      reqUri,
			"clientIP": clientIP,
		}).Info("")
		//})
	}
}
