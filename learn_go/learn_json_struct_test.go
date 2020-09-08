package learn_go

import (
	"encoding/json"
	"fmt"
	"testing"
)

//结构体转为json字符串
//需要注意的是将结构体转换为Json数据时候，定义结构体的字段必须首字母大写。否则无法正常解析。

//结构体一
type Prescript struct {
	Name     string
	Unit     string
	Additive *Additive
}

//结构体二
type Additive struct {
	Name string
	Unit string
}

//1、结构体转为json字符串
func Test_Json_Pre(t *testing.T) {
	p := Prescript{}
	p.Name = "鹤顶红"
	p.Unit = "1.2kg"
	p.Additive = &Additive{
		Name: "砒霜",
		Unit: "0.5kg",
	}

	buf, err := json.Marshal(p) //转换为json返回两个结果
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("json = ", string(buf)) //json =  {"Name":"鹤顶红","Unit":"1.2kg","Additive":{"Name":"砒霜","Unit":"0.5kg"}}
}

//可以看出其中json字符中每一个key的首字母也是大写，最后一个没有设置数据的字段的结果为null。
// 那么如何强制将他变为小写的。并且将不需要显示的字段隐藏掉。就需要在结构体上添加标记。
// Name string `json:"name"`

type Prescript2 struct {
	Name     string     `json:"name"` //重新指定json字段为小写输出
	Unit     string     `json:"unit"`
	Additive *Additive2 `json:"additive,omitempty"`
}

type Additive2 struct {
	Name string `json:"name"`
	Unit string `json:"unit"`
}

func Test_Json_Pre2(t *testing.T) {
	p := Prescript2{}
	p.Name = "鹤顶红"
	p.Unit = "1.2kg"
	p.Additive = &Additive2{
		Name: "砒霜",
		Unit: "0.5kg",
	}

	buf, err := json.Marshal(p) //转换为json返回两个结果
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//在结构体加上标记之后再转换之后的json字符串就会舒服多了
	fmt.Println("json = ", string(buf)) //json =  {"name":"鹤顶红","unit":"1.2kg","additive":{"name":"砒霜","unit":"0.5kg"}}
}

//2、json字符串转为结构体
func Test_json_Struct(t *testing.T) {
	jsonstr := `{"name":"鹤顶红","unit":"1.2kg","additive":{"name":"砒霜","unit":"0.5kg"}}`
	var p Prescript
	if err := json.Unmarshal([]byte(jsonstr), &p); err != nil { //字符串转换成byte, 将字符串转为这个&p这个结构体
		fmt.Println(err)
	}
	fmt.Println(p)
}
