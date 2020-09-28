package learn_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func Test_AsciiJson(t *testing.T) {
	r := gin.Default()
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	_ = r.Run(":8080")
}
