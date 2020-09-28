package learn_gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

//Map as querystring or postform parameters

//POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
//Content-Type: application/x-www-form-urlencoded
//
//names[first]=thinkerou&names[second]=tianou

func Test_map_query(t *testing.T) {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})
	_ = router.Run(":8080")
}
