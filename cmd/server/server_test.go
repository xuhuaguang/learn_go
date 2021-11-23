package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn_go/task"
	"math/rand"
	"testing"
	"time"
)

func TestLoader(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	router := gin.New()
	task.StartAllDataLoader(10)
	_ = router.Run(fmt.Sprintf(":%d", 8888))
}

type respData struct {
	RequestID string `json:"requestId,omitempty"` //每⼀次请求的 UUID,由服务器 响应⽣成
	Ads       []*Ads `json:"ads,omitempty"`       //⼴告数组
}

type Ads struct {
	Advertiserid  string         `json:"advertiserid,omitempty"`  //⼴告ID
	Adid          string         `json:"adid,omitempty"`          //创意ID
	IsApp         bool           `json:"isApp,omitempty"`         //判断⼴告是否是 下载类⼴告 true: 下载类 false：⾮下载类
	OpenExternal  bool           `json:"openExternal,omitempty"`  //判断是否是需要 调起三⽅App true:是  false:不是
	TargetUrl     string         `json:"targetUrl,omitempty"`     //落地⻚链接。落地⻚的url 或者 App的下载地址【如 果clickTracker中包 含“c.gdt.qq.com”域 名的Tracker链接且 isApp&openExtern al为True时，此参数 值⽆效，需要将该 Tracker链接进⾏ 302跳转获取实际的 URL；详⻅宏替换 需求⼀】
	Lp            string         `json:"lp,omitempty"`            //落地⻚地址，⾮原⽣⼴告.例如 ：banner、interstitial
	Adtype        string         `json:"adtype,omitempty"`        //⼴告类型，⾮原 ⽣⼴告
	DealId        string         `json:"dealId,omitempty"`        //PMP模式下的竞 价ID
	PackageName   string         `json:"packageName,omitempty"`   //Android 包名
	Bid           int            `json:"bid,omitempty"`           //竞价的价格，单 位为：分
	PubContent    *PubContent    `json:"pubContent,omitempty"`    //返回单个⼴告具 体内容
	EventTracking *EventTracking `json:"eventTracking,omitempty"` //上报链接集合
}

type EventTracking struct {
	ClickTrackers      []string `json:"clickTrackers,omitempty"`      //展示上报URL
	ImpressionTrackers []string `json:"impressionTrackers,omitempty"` //点击上报URL
	DplAttempt         []string `json:"dplAttempt,omitempty"`         //deeplink被调⽤ 上报URL
	DplSuccess         []string `json:"dplSuccess,omitempty"`         //应⽤已安装且成功调⽤ deeplink 上报URL
	DownloadStart      []string `json:"downloadStart,omitempty"`      //⽤户点击下载按钮上报 URL （仅适⽤于安卓
	DownloadFinis      []string `json:"downloadFinis,omitempty"`      //⽤户下载完成上报URL （下载类⼴告，仅适⽤于 安卓）
	InstallFinish      []string `json:"installFinish,omitempty"`      //⽤户安装成功上报URL （下载类⼴告，仅适⽤于 安卓）
}

type PubContent struct {
	Title       string       `json:"title,omitempty"`       //⼴告标题
	Description string       `json:"description,omitempty"` //⼴告描述
	Icon        *Icon        `json:"icon,omitempty"`        //⼴告ICON
	Screenshots *Screenshots `json:"screenshots,omitempty"` //图⽚⼴告资源实体
	LandingURL  string       `json:"landingURL,omitempty"`  //deeplink⼴告对 应调起地址.openExternal字 段为true时，先 调⽤此字段尝试 调起三⽅app， 例如京 东;openExternal 字段为false时,此 字段不需要处理
	Cta         string       `json:"cta,omitempty"`         //⼴告按钮⽂字
	Rating      float64      `json:"rating,omitempty"`      //评分等级
}

type Icon struct {
	Width       float32 `json:"width,omitempty"`       //icon的宽
	Height      float32 `json:"height,omitempty"`      //icon的⾼
	URL         string  `json:"url,omitempty"`         //⼴告ICON的URL地址
	AspectRatio float64 `json:"aspectRatio,omitempty"` //图像宽⾼⽐例
}

type Screenshots struct {
	Width       float32 `json:"width,omitempty"`
	Height      float32 `json:"height,omitempty"`
	URL         string  `json:"url,omitempty"`
	AspectRatio float64 `json:"aspectRatio,omitempty"`
}
