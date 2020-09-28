package learn_gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func Test_Redirects(t *testing.T) {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})

	r.POST("/test", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})

	r.GET("/test7", func(c *gin.Context) {
		c.Request.URL.Path = "/test8"
		fmt.Println(c.Request.URL.Path)
		r.HandleContext(c)
	})

	r.GET("/test8", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	_ = r.Run(":8080")
}
