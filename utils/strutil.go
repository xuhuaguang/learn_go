package utils

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

// Bytes2String直接转换底层指针，两者指向的相同的内存，改一个另外一个也会变。
// 效率是string([]byte{})的百倍以上，且转换量越大效率优势越明显。
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// String2Bytes直接转换底层指针，两者指向的相同的内存，改一个另外一个也会变。
// 效率是string([]byte{})的百倍以上，且转换量越大效率优势越明显。
// 转换之后若没做其他操作直接改变里面的字符，则程序会崩溃。
// 如 b:=String2bytes("xxx"); b[1]='d'; 程序将panic。
func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Equal(s1, s2 string) bool {
	return 0 == strings.Compare(s1, s2)
}

func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func IntToStr(d int) string {
	return fmt.Sprintf("%d", d)
}

func Int32ToStr(d int32) string {
	return fmt.Sprintf("%d", d)
}

func IsNotEmpty(s string) bool {
	return len(s) > 0
}

func IsEmpty(s string) bool {
	return len(s) == 0
}

// JoinURL Creates a URL string from parts
func JoinURL(elem ...string) string {
	parts := []string{}
	for _, e := range elem {
		if strings.HasSuffix(e, "/") {
			e = strings.TrimSuffix(e, "/")
		}
		if strings.HasPrefix(e, "/") {
			e = strings.TrimPrefix(e, "/")
		}
		parts = append(parts, e)
	}
	res := strings.Join(parts, "/")
	return res
}

func RemoveEmptyStr(arr []string) (r []string) {
	for _, s := range arr {
		if IsNotEmpty(s) {
			r = append(r, s)
		}
	}
	return
}

func StrWithDefault(val, def string) string {
	if IsEmpty(val) {
		return def
	} else {
		return val
	}
}
