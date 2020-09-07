package utils

import (
	"fmt"
	"path"
	"strings"
	"testing"
)

//获取文件的名称、前缀、后缀
func TestFile(t *testing.T) {
	filename := "http://cyad.d.zhangyue01.com/dsp-barnd/CREATIVE/20200826/8097aebd5791a6a4c02b9de5a6b608ca.mp4"
	filenameall := path.Base(filename)                              //获取不包含目录的文件名

	filesuffix := path.Ext(filename)                                //获取文件后缀
	filesuffix2 := filename[strings.LastIndex(filename, ".")+1:]    //获取文件后缀

	fileprefix := filenameall[0 : len(filenameall)-len(filesuffix)] //获取文件名
	fileprefix2 := strings.TrimSuffix(filenameall, filesuffix)      //获取文件名

	fmt.Println("file name:", filenameall)    // 8097aebd5791a6a4c02b9de5a6b608ca.mp4
	fmt.Println("file prefix:", fileprefix)   // 8097aebd5791a6a4c02b9de5a6b608ca
	fmt.Println("file prefix2:", fileprefix2) // 8097aebd5791a6a4c02b9de5a6b608ca
	fmt.Println("file suffix:", filesuffix)   // .mp4
	fmt.Println("file suffix2:", filesuffix2) // mp4
}
