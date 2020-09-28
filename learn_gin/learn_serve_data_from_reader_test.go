package learn_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

//Serving data from reader

func Test_Serving_data_from_reader(t *testing.T) {
	router := gin.Default()
	router.GET("/someDataFromReader", server)
	_ = router.Run(":8080")
}

func server(c *gin.Context) {
	response, err := http.Get("https://dongfeng.alicdn.com/201708/2d51aec7defb4df4bdc9676fa6881d7b.jpg")
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}
