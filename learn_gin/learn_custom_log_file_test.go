package learn_gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
	"time"
)

func Test_Custom_Log_File(t *testing.T) {
	router := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(logFile))
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	_ = router.Run(":8080")
}

func logFile(param gin.LogFormatterParams) string {
	// your custom format
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

//::1 - [Tue, 15 Sep 2020 12:08:44 CST] "GET /ping HTTP/1.1 200 0s "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36" "
