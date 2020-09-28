package learn_gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

func Test_Multiple_files(t *testing.T) {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// Upload the file to specific dst.
			_ = c.SaveUploadedFile(file, "D:\\Software")
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	_ = router.Run(":8080")
}
