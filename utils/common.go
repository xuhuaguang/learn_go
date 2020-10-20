package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/segmentio/ksuid"
	"io"
	"learn_go/models"
	"log"
	"math/rand"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func ToJsonStr(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}

func GenKsuid() ksuid.KSUID {
	id := ksuid.New()
	return id
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}

func EnsureDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(path, 0755)
	}
}

func StrListContains(list []string, val string) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}
func IntListContains(list []int, val int) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func StrListValveContains(clickTrackers []string, value string) bool {
	for _, v := range clickTrackers {
		if strings.Contains(v, value) {
			return true
		}
	}
	return false
}

//数组去重
func DuplicateIntArray(m []string) []string {
	s := make([]string, 0)
	samp := make(map[string]int)
	for _, value := range m {
		//计算map长度
		length := len(samp)
		samp[value] = 1
		//比较map长度, 如果map长度不相等， 说明key不存在
		if len(samp) != length {
			s = append(s, value)
		}
	}
	return s
}

// 对字符串进行md5哈希,
// 返回32位小写md5结果
func Md5(in string) string {
	has := md5.Sum([]byte(in))
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// 对字符串进行md5哈希,
// 返回32位大写写md5结果
func MD5ToUpper(s string) string {
	return strings.ToUpper(Md5(s))
}

// 对字符串进行md5哈希,
// 返回16位小写md5结果
func MD5_16(s string) string {
	return Md5(s)[8:24]
}

func Sha1(str string) string {
	s := sha1.New()
	_, _ = io.WriteString(s, str)
	return hex.EncodeToString(s.Sum(nil))
}

func ToString(e interface{}) string {
	res := ""
	if e == nil {
		return res
	}
	switch v := e.(type) {
	case string:
		res = e.(string)
		break
	case int:
		res = strconv.FormatInt(int64(v), 10)
		break
	case int32:
		res = strconv.FormatInt(int64(v), 10)
		break
	case int64:
		res = strconv.FormatInt(v, 10)
		break
	case float32:
		res = strconv.FormatFloat(float64(v), 'f', -1, 64)
		break
	case float64:
		res = strconv.FormatFloat(v, 'f', -1, 64)
		break
	}
	return res
}

func GetCurrentDateAsInt() int {
	i, _ := strconv.Atoi(time.Now().Format("20060102"))
	return i
}

func TransferTime(time int32) string {
	msg := ""
	h := time / 3600
	m := (time % 3600) / 60
	s := (time % 3600) % 60
	if time < 60 {
		msg = fmt.Sprintf("%d秒", s)
	} else if time == 60 {
		msg = fmt.Sprintf("%d分钟", m)
	} else if time < 3600 && time > 60 {
		if s == 0 {
			msg = fmt.Sprintf("%d分钟", m)
		} else {
			msg = fmt.Sprintf("%d分钟%d秒", m, s)
		}
	} else if time == 3600 {
		msg = fmt.Sprintf("%d小时", h)
	} else {
		if s == 0 && m == 0 {
			msg = fmt.Sprintf("%d小时", h)
		} else if s == 0 && m != 0 {
			msg = fmt.Sprintf("%d小时%d分钟", h, m)
		} else if s != 0 && m == 0 {
			msg = fmt.Sprintf("%d小时%d秒", h, s)
		} else {
			msg = fmt.Sprintf("%d小时%d分钟%d秒", h, m, s)
		}
	}
	return msg
}

//提前确认值为空
func Strings(key string) []string {
	return strings.Split(key, ",")
}

func GetAroundDaysTime(days int) string {
	unix := time.Now().AddDate(0, 0, days).Unix()
	return GetUnixToString(unix)
}

func GetUnixToString(timestamp int64) string {
	return GetUnixToFormatString(timestamp, models.DateFormat_yyyy_MM_dd)
}

func GetUnixToFormatString(timestamp int64, f string) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(f)
}

func GetCurrentDate() string {
	return GetUnixToFormatString(time.Now().Unix(), models.DateFormat_yyyy_MM_dd)
}

//获取当天剩余的秒数
func GetCurrentSurplusSeconds() int64 {
	timeCur := time.Now()                                                                          //当前时间
	timeNow := timeCur.Unix()                                                                      //当前时间戳
	timeSet := time.Date(timeCur.Year(), timeCur.Month(), timeCur.Day()+1, 0, 0, 0, 0, time.Local) //下一天0点的时间
	timeSetUnix := timeSet.Unix() - timeNow                                                        //到下一天0点剩余秒数(用于设置KEY有效时间)
	return timeSetUnix
}

func GetIsNewUser(timestamp int64) bool { //true-->新用户  false --> 老用户
	currentDate := GetCurrentDate()
	dataUserSignDate := GetUnixToFormatString(timestamp/models.SecondMs, models.DateFormat_yyyy_MM_dd)
	if currentDate == dataUserSignDate { //第一次签到时间（观看视频）在当天范围内，都认为是新用户
		return true
	}
	return false
}

func GetTimeArr(start, end string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	// 转成时间戳
	startUnix, _ := time.ParseInLocation(timeLayout, start, loc)
	endUnix, _ := time.ParseInLocation(timeLayout, end, loc)
	startTime := startUnix.Unix()
	endTime := endUnix.Unix()
	// 求相差天数
	date := (endTime - startTime) / 86400
	//times()
	return date
}

func GetNowMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//解码
func UrlDecode(urlStr string) string {
	deEscapeUrl, _ := url.QueryUnescape(urlStr)
	return deEscapeUrl
}

//编码
func UrlEncode(urlStr string) string {
	escapeUrl := url.QueryEscape(urlStr)
	return escapeUrl
}

func GetPlanMappingId(planId string) int {
	if len(planId) > 2 {
		plan := planId[len(planId)-2:]
		planId := planMap[plan]
		return StrToInt(planId)
	}
	return 0
}

func IntArrayToString(array []int) string {
	if len(array) == 0 {
		return ""
	}
	trim := strings.Trim(fmt.Sprint(array), "[]")
	return strings.Replace(trim, " ", ",", -1)
}

//范围之内的数字
func RandInt(min, max int) int {
	if min >= max || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

//[0,100)随机值，在[0,70)返回true，反之返回false
func RandShowRule(min, max, divValue int) bool {
	if RandInt(min, max) < divValue {
		return true
	} else {
		return false
	}
}

//获取minute分钟后的时间
func AddMinuteTimestamp(minute int) int64 {
	now := time.Now() //获取当前时间
	unix := now.Add(time.Minute * time.Duration(minute)).Unix()
	return unix
}

//获取增加minute分钟后的时间
func GetAddAfterTime(min, max, addMinute int) int64 {
	return AddMinuteTimestamp(RandInt(min, max) + addMinute)
}

func StringTransArray(s string, separator string) []string {
	return strings.Split(s, separator)
}

func DupStringArrayTransString(m []string, separator string) string {
	s := ""
	samp := make(map[string]int)
	for _, value := range m {
		//计算map长度
		length := len(samp)
		samp[value] = 1
		//比较map长度, 如果map长度不相等， 说明key不存在
		if len(samp) != length {
			if IsEmpty(s) {
				s = fmt.Sprintf("%s", value)
			} else {
				s = fmt.Sprintf("%s%s%s", s, separator, value)
			}
		}
	}
	return s
}

func StringArrayTransString(m []string, separator string) string {
	return strings.Join(m, separator)
}

func RemoveRepByLoop(slc []int) []int {
	var result []int // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

//方案对应关系
var planMap = map[string]string{
	"59": "73",
	"61": "73",
	"54": "46",
	"66": "46",
	"64": "76",
	"65": "77",
	"03": "41",
	"07": "44",
	"68": "34",
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
