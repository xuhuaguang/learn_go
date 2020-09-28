package learn_gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func Test_Grouping_routes(t *testing.T) {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	_ = router.Run(":8080")
}

func loginEndpoint(c *gin.Context) {
	fmt.Println("login")
}

func submitEndpoint(c *gin.Context) {
	fmt.Println("submit")
}

func readEndpoint(c *gin.Context) {
	fmt.Println("read")
}
