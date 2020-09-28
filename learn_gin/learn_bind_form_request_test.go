package learn_gin

import (
	"github.com/gin-gonic/gin"
	"testing"
)

//Bind form-data request with custom struct

//$ curl "http://localhost:8080/getb?field_a=hello&field_b=world"
//		{"a":{"FieldA":"hello"},"b":"world"}
//$ curl "http://localhost:8080/getc?field_a=hello&field_c=world"
//		{"a":{"FieldA":"hello"},"c":"world"}
//$ curl "http://localhost:8080/getd?field_x=hello&field_d=world"
//		{"d":"world","x":{"FieldX":"hello"}}

func Test_Bind_FormData(t *testing.T) {
	r := gin.Default()
	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getd", GetDataD)

	_ = r.Run()
}

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	_ = c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var b StructC
	_ = c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var b StructD
	_ = c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}
