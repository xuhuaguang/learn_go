package learn_go

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

//strings包常用方法

func Test_Strings_Method(t *testing.T) {
	//1、是否包含指定内容 返回bool类型
	s1 := "ok let's go"
	fmt.Println(strings.Contains(s1, "go")) //结果为true

	//2、是否包含指定的字符串中任意一个字符 有一个出现过 就返回true
	fmt.Println(strings.ContainsAny(s1, "glass")) //结果为true。因为字母g、l

	//3、返回指定字符出现的次数
	fmt.Println(strings.Count(s1, "g")) //1

	//4、文本的开头
	fmt.Println(strings.HasPrefix(s1, "ok")) //结果为true

	//5、文本的结尾
	fmt.Println(strings.HasSuffix(s1, ".txt")) //结果为false

	//6、查找指定字符在字符串中存在的位置 如果不存在返回-1
	fmt.Println(strings.Index(s1, "g")) //9

	//7、查找字符中任意一个字符出现在字符串中的位置
	fmt.Println(strings.IndexAny(s1, "s")) //7

	//8、查找指定字符出现在字符串中最后一个的位置
	fmt.Println(strings.LastIndex(s1, "s")) //7

	//9、字符串的拼接
	s2 := []string{"123n", "abc", "ss"}
	s3 := strings.Join(s2, "_")
	fmt.Println(s3) // 123n_abc_ss

	//10、字符串的切割
	s4 := strings.Split(s3, "_")
	fmt.Println(s4) // 返回切片[]string{"123n","abc","ss"}

	//11、字符串的替换
	s5 := "okoletsgo"
	s6 := strings.Replace(s5, "o", "*", 1)
	fmt.Println(s6) //*koletsgo
	s7 := strings.Replace(s5, "o", "*", -1)
	fmt.Println(s7) //*k*letsg*
	//TODO 1 只替换1次,  -1 全部替换

	//12、字符串的截取
	//str[start:end]包含start 不包含end
	fmt.Println(s1[2:7])

	//13、获取url的域名
	//https://fc.ele.me/a/ODE0NDg2MWI0YjFjMTFlYjlhM2QwMjQyMGI1OWUxMjQ=
	s8 := "https://fc.ele.me/a/ODE0NDg2MWI0YjFjMTFlYjlhM2QwMjQyMGI1OWUxMjQ="
	domain := strings.Split(strings.Split(s8,"//")[1],"/")[0]
	s, _ := url.Parse(s8)
	println(domain)
	println(s.Host)

}
