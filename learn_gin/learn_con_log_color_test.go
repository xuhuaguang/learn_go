package learn_gin

import (
	"github.com/gin-gonic/gin"
	"testing"
)

//Controlling Log output coloring

func Test_Controller_Color_Log(t *testing.T) {
	// Disable log's color
	gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	_ = router.Run(":8080")
}

func Test_Controller_Color_Log2(t *testing.T) {
	// Force log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/ping", handle2)

	_ = router.Run(":8080")
}

func handle2(c *gin.Context)  {
	c.String(200, "pong")
}