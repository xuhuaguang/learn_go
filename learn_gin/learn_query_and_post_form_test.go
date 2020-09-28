package learn_gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

//Query and post form

//POST /post?id=1234&page=1 HTTP/1.1
//Content-Type: application/x-www-form-urlencoded
//
//name=manu&message=this_is_great

func Test_Query_and_post_form(t *testing.T) {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	_ = router.Run(":8080")
}
