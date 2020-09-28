package learn_gin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"testing"
)

//Only bind query string

type Person3 struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func Test_only_bind_query(t *testing.T) {
	route := gin.Default()
	route.Any("/testing", startPage2)
	_ = route.Run(":8080")
}

func startPage2(c *gin.Context) {
	var person Person3
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	bytes, _ := json.Marshal(person)
	fmt.Println(string(bytes))
	c.String(200, "Success")
}
