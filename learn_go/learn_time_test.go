package learn_go

import (
	"fmt"
	"testing"
	"time"
)

//time包操作的都是时间，时间的单位都包括年，月，日，时，分，秒，毫秒，微妙，纳秒，皮秒。
//需要注意的室Go语言中时间的格式化，需要指定格式化时间的模板, 不管年月日的类型格式怎么写，
// 但是具体的数值必须写成2006-01-02 15:04:05， 如果不是这个日期就不能够格式化，这个时间也是为了纪念Go语言诞生的时间。

//格式都是time.Time类型的数据
func Test_Time(t *testing.T) {
	//获取当前时间
	time2 := time.Now()
	fmt.Println(time2) //2020-03-31 21:26:01.7307507 +0800 CST m=+0.001999001
	//获取的时间后面的信息是时区

	//上面的时间看起来不是很方便 于是需要格式化时间
	s := time2.Format("2006年1月2日 15:04:05")
	fmt.Println(s) //2020年9月8日 12:10:26

	s2 := time2.Format("2006-1-2 15:04:05")
	fmt.Println(s2) //打印出的格式就是当前的时间 2020-3-31 23:08:35

	s3 := time2.Format("2006/1/2")
	fmt.Println(s3) //打印出的格式就是当前的年月日 2020/3/31
}

//将string类型的字符串时间转为具体时间格式则用time包下的parse函数。
func Test_Str_Time(t *testing.T) {
	//字符串类型的时间
	str := "2020年3月31日"
	//第一个参数是模板,第二个是要转换的时间字符串
	s, _ := time.Parse("2006年1月2日", str)
	fmt.Println(s) //打印出的格式就是2020-03-31 00:00:00 +0000 UTC

	//获取年月日信息
	year, month, day := time.Now().Date()
	fmt.Println(year, month, day) //2020 March 31

	//获取时分秒信息
	hour, minute, second := time.Now().Clock()
	fmt.Println(hour, minute, second) //23 23 54

	//获取今年过了多少天了
	tday := time.Now().YearDay()
	fmt.Println(tday) //91  (今年已经过了91天了)

	//获取今天是星期几
	weekday := time.Now().Weekday()
	fmt.Println(weekday) //Tuesday
}

//时间戳
func Test_Time_Stamp(t *testing.T) {
	//获取指定日期的时间戳
	time3 := time.Date(2020, 3, 31, 23, 30, 0, 0, time.UTC)
	timestamp := time3.Unix()
	fmt.Println(timestamp) //1585697400

	//获取当前时间的时间戳
	timestamp2 := time.Now().Unix()
	fmt.Println(timestamp2) //1585669151

	//当前时间的以纳秒为单位的时间戳
	timestamp3 := time.Now().UnixNano()
	fmt.Println(timestamp3) //1585669151296330900

	fmt.Println("---------华丽的分割线----------")
	//时间间隔 相加
	now := time.Now()
	//当前时间加上一分钟
	time4 := now.Add(time.Minute)
	fmt.Println(now)   //2020-03-31 23:43:35.0004791 +0800 CST m=+0.002999201
	fmt.Println(time4) //2020-03-31 23:44:35.0004791 +0800 CST m=+60.002999201

	//计算两个时间的间隔
	d := time4.Sub(now)
	fmt.Println(d) //1m0s  相差一分钟

	fmt.Println("---------时间戳与时间格式互转----------")
	//将指定时间转为时间戳格式
	beforetime := "2020-04-08 00:00:00"                             //待转化为时间戳的字符串
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc := time.Now().Location()                                    //获取时区
	theTime, _ := time.ParseInLocation(timeLayout, beforetime, loc) //使用模板在对应时区转化为time.time类型
	aftertime := theTime.Unix()                                     //转化为时间戳 类型是int64
	fmt.Println(theTime)                                            //打印输出theTime 2020-04-08 00:00:00 +0800 CST
	fmt.Println(aftertime)                                          //打印输出时间戳 1586275200

	//再将时间戳转换为日期
	dataTimeStr := time.Unix(aftertime, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println(dataTimeStr)
}
