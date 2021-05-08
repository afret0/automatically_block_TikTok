package user

import (
	"backend/source"
	"backend/user/tokenManager"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "newtrekWang"
	logger           *logrus.Logger
)

func init() {
	logger = source.GetLogger()
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("tokenManager")
		if len(token) < 1 {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "未携带 tokenManager, 请先登录"})
			c.Abort()
			return
		}
		//j := new(jwt.Claims)
		claims, err := tokenManager.ParseToken(token)
		if err != nil {
			logger.Errorln(token, err)
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
