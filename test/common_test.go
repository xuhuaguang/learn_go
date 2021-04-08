package test

import (
	"fmt"
	"learn_go/entity"
	"learn_go/utils"
	"math"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var res []*ResponseInfo

var zyBidArray []*entity.ZyBid

func TestRands(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var winningBid *ResponseInfo
	index := rand.Intn(3) //012
	if len(res) >= 3 {
		winningBid = res[index]
	} else {
		if index < len(res) {
			winningBid = res[index]
		}
	}
	if winningBid == nil {
		fmt.Println("cpt 论空了,index=", index)
	} else {
		fmt.Println(winningBid.ID)
	}
}

func init() {
	for i := 1; i <= 2; i++ {
		info := &ResponseInfo{
			ID:         utils.IntToStr(i),
			CreativeId: utils.IntToStr(i),
			Price:      i,
			Adm:        fmt.Sprintf("adm%d", i),
		}
		res = append(res, info)
	}

	for i := 1; i <= 6; i++ {

		if i == 2 || i == 4 {
			for i := 1; i <= 2; i++ {
				zyBidArray = append(zyBidArray, getZyBid(i))
			}
		} else {
			zyBidArray = append(zyBidArray, getZyBid(i))
		}
	}
}

func getZyBid(i int) *entity.ZyBid {
	return &entity.ZyBid{
		ID:                "",
		AdverId:           "",
		CreativeId:        fmt.Sprintf("%d", i),
		Price:             i * 100,
		Adm:               "adm",
		LandingPageUrl:    "dp",
		LandingPageAction: i,
		DeepLinkUrl:       "www.baidu.com",
		Width:             i * 2,
		Height:            i * 4,
		TemplateId:        fmt.Sprintf("1_%d", i),
		AppPackage:        "app",
		AppName:           "app_name",
		AppVersion:        "app_version",
		From:              "SSP",
		CreativeIdType:    utils.RandInt(10, 12),
		BidType:           utils.RandInt(10, 12),
		CampaignStartDate: 0,
		CampaignEndDate:   0,
		CptRatio:          rand.Intn(3),
		CardPlace:         i,
	}
}

type ResponseInfo struct {
	ID         string `json:"id"`
	CreativeId string `json:"cid"`
	Price      int    `json:"price"`
	Adm        string `json:"adm"`
}

func TestRangeArray(t *testing.T) {
	var array []*ResponseInfo
	for _, v := range array {
		fmt.Println("---------", v)
	}
}

//测试卡片流轮播
func TestCardsVideo(t *testing.T) {
	bidGroupByPlace := make(map[int][]*entity.ZyBid)
	for _, res := range zyBidArray {
		bidderResponses := bidGroupByPlace[res.CardPlace]
		bidderResponses = append(bidderResponses, res)
		bidGroupByPlace[res.CardPlace] = bidderResponses
	}

	fmt.Println("以插入位置分组的大小==》", len(bidGroupByPlace))

	var matchBidsAll []*entity.ZyBid

	for k, data := range bidGroupByPlace {
		fmt.Println(fmt.Sprintf("插入位置=%d,数组大小%d", k, len(data)))

		//安装bidType类型分组
		bidGroupByType := make(map[int][]*entity.ZyBid)

		for _, v := range data {
			bidderResponses := bidGroupByPlace[v.BidType]
			bidderResponses = append(bidderResponses, v)
			bidGroupByType[v.BidType] = bidderResponses
		}

		cptBidsCards := bidGroupByType[BidBrandCpt]

		var matchBids []*entity.ZyBid //符合的响应素材
		matchBids = append(matchBids, bidGroupByType[BidBrandCpm]...)
		matchBids = append(matchBids, bidGroupByType[BidType_Cpm]...)

		// 按购买轮播数权重分配
		var weightedBids []*entity.ZyBid
		for _, bid := range cptBidsCards {
			for i := 0; i < bid.CptRatio; i++ {
				weightedBids = append(weightedBids, bid)
			}
		}
		if len(weightedBids) > 0 {
			rand.Shuffle(len(weightedBids), func(i, j int) { weightedBids[i], weightedBids[j] = weightedBids[j], weightedBids[i] })
		}

		var winningBid *entity.ZyBid
		index := rand.Intn(3) //012
		if len(weightedBids) >= 3 {
			winningBid = weightedBids[index]
		} else {
			if index < len(weightedBids) {
				winningBid = weightedBids[index]
			}
		}

		for _, bid := range cptBidsCards {
			if winningBid != nil && bid.CreativeId == winningBid.CreativeId {
				matchBids = append(matchBids, winningBid)
			} else {
				//todo 错误处理
			}
		}

		//todo 此位置符合的全部响应数据 matchBids
		matchBidsAll = append(matchBidsAll, matchBids...)
	}

	fmt.Println("原先的数据大小=", len(zyBidArray))
	fmt.Println("过滤后的数据大小=", len(matchBidsAll))
}

const (
	BidType_Cpm  int = 0
	BidType_Cpc      = 1
	BidBrandCpt      = 10
	BidBrandCpm      = 11
	BidBrandCard     = 12
)

func TestTimeDiff(t *testing.T) {
	//diff := getDateDiff(1608794098244, getCurrentTimeMs())
	diff := getDateDiff(1606476056000, getCurrentTimeMs())
	fmt.Println(diff)

	newOs := ""

	os := "7.0"
	if strings.Contains(os, ".") {
		split := strings.Split(os, ".")
		newOs = split[0]
	} else {
		newOs = os
	}
	fmt.Println(newOs)

	b := newOs == "1"
	fmt.Println(b)
}

// 向上取整相差天数
func getDateDiff(start, currentTime int64) int64 {
	date := float64(currentTime-start) / secondMs / oneDaySecond
	return int64(math.Ceil(date))
}

//获取int64 当前时间戳/输入time时间戳
func getCurrentTimeMs() int64 {
	return time.Now().UnixNano() / 1e6
}

const (
	oneDaySecond = 24 * 60 * 60
	secondMs     = 1000
)
