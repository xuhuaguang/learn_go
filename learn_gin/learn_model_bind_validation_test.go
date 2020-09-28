package learn_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

//Model binding and validation

//Also, Gin provides two sets of methods for binding:
//
//Type - Must bind
//	Methods - Bind, BindJSON, BindXML, BindQuery, BindYAML
//	Behavior - These methods use MustBindWith under the hood. If there is a binding error,
//   	the request is aborted with c.AbortWithError(400, err).SetType(ErrorTypeBind).
//   	This sets the response status code to 400 and the Content-Type header is set to text/plain;
//   	charset=utf-8. Note that if you try to set the response code after this,
//   	it will result in a warning [GIN-debug] [WARNING] Headers were already written.
//   	Wanted to override status code 400 with 422. If you wish to have greater control over the behavior, consider using the ShouldBind equivalent method.
//Type - Should bind
//	Methods - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYAML
//	Behavior - These methods use ShouldBindWith under the hood. If there is a binding error,
//     the error is returned and it is the developer’s responsibility to handle the request and error appropriately.

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func Test_model_bind(t *testing.T) {
	router := gin.Default()

	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>user</user>
	//		<password>123</password>
	//	</root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding a HTML form (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Listen and serve on 0.0.0.0:8080
	_ = router.Run(":8080")
}

//$ curl -v -X POST \
//  http://localhost:8080/loginJSON \
//  -H 'content-type: application/json' \
//  -d '{ "user": "manu" }'
//> POST /loginJSON HTTP/1.1
//> Host: localhost:8080
//> User-Agent: curl/7.51.0
//> Accept: */*
//> content-type: application/json
//> Content-Length: 18
//>
//* upload completely sent off: 18 out of 18 bytes
//< HTTP/1.1 400 Bad Request
//< Content-Type: application/json; charset=utf-8
//< Date: Fri, 04 Aug 2017 03:51:31 GMT
//< Content-Length: 100
//<
//{"error":"Key: 'Login.Password' Error:Field validation for 'Password' failed on the 'required' tag"}