package learn_gin

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"testing"
	"time"
)

//$ curl -X GET "localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15"

func Test_Query_String_Post(t *testing.T) {
	route := gin.Default()
	route.GET("/testing", startPage)
	_ = route.Run(":8085")
}

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}

	c.String(200, "Success")
}

func TestPrice(t *testing.T) {
	zyPrice := int64(0)
	zyPrice = int64(2*float64(Ssp_Yuan)) / CPM
	fmt.Println(zyPrice)
	zhyFinalPrice := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", zyPrice)))
	fmt.Println(zhyFinalPrice)
}

const (
	CPM            = 1000
	Ssp_Yuan int64 = 100 * CPM
)
