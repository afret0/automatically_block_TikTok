package source

//func LoggerMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		startT := time.Now()
//		c.Next()
//		endT := time.Now()
//		latencyT := endT.Sub(startT)
//		reqMethod := c.Request.Method
//		reqUri := c.Request.RequestURI
//		clientIP := c.ClientIP()
//		middlewareLogger.WithFields(logrus.Fields{
//			"reqTime":  startT.Format("2006-01-02 15:04:05"),
//			"latencyT": latencyT,
//			"method":   reqMethod,
//			"uri":      reqUri,
//			"clientIP": clientIP,
//		}).Info("")
//		//})
//	}
//}