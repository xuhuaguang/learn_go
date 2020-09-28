package learn_gin

import (
	"github.com/gin-gonic/gin"
	"testing"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//curl -v --form user=user --form password=password http://localhost:8080/login

func Test_Urlencoded_binding(t *testing.T) {
	router := gin.Default()
	router.POST("/login", login)
	_ = router.Run(":8080")
}

func login(c *gin.Context) {
	// you can bind multipart form with explicit binding declaration:
	// c.ShouldBindWith(&form, binding.Form)
	// or you can simply use autobinding with ShouldBind method:
	var form LoginForm
	// in this case proper binding will be automatically selected
	if c.ShouldBind(&form) == nil {
		if form.User == "user" && form.Password == "password" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	}
}

func Test_Urlencoded_binding2(t *testing.T) {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	_ = router.Run(":8080")
}
