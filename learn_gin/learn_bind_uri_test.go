package learn_gin

import (
	"github.com/gin-gonic/gin"
	"testing"
)

type Person2 struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func Test_Bind_Uri(t *testing.T) {
	route := gin.Default()
	route.GET("/:name/:id", handle)
	_ = route.Run(":8088")
}

//$ curl -v localhost:8088/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
//$ curl -v localhost:8088/thinkerou/not-uuid

func handle(c *gin.Context) {
	var person Person2
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
}
