package learn_gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func Test_Set_get_cookie(t *testing.T) {
	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	_ = router.Run()
}
