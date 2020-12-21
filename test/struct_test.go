package test

import (
	"encoding/json"
	"fmt"
	"learn_go/utils"
	"sort"
	"testing"
	"time"
)

type VideoRewardCps struct {
	Id            int      `json:"id"`            //
	Limit         int      `json:"limit"`         //观看次数
	Interval      int      `json:"interval"`      //观看间隔时间, 单位秒
	GiftType      string   `json:"gift_type"`     //礼物类型, voucher: 代金券, book:书籍章节, adless: 无广告, vouchers: 阶梯代金券, wheel: 大转盘
	GiftId        int32    `json:"gift_id"`       //礼物id, 赠送类型为代金券时的字段
	WheelId       int32    `json:"wheel_id"`      //大转盘id
	UserGroupId   int      `json:"user_group_id"` // 用户集合id
	SendType      int      `json:"send_type"`     //赠送代金券的类型, 用于代金券详情的来源
	SendDesc      string   `json:"send_desc"`     //赠送描述
	Amount        int32    `json:"amount"`        //代金券金额,  动态代金券
	Amounts       []string `json:"amounts"`       //阶梯代金券金额,  列表
	ChapterCount  int32    `json:"chapter_count"` //赠送的章节数量  (赠送类型为书籍章节时的字段)
	Duration      int32    `json:"duration"`      //无广告时长,单位秒  (赠送类型为无广告的字段)
	OvRewardVideo OvRewardVideo
}

//兼容cps策略，ov章尾激励视频
type OvRewardVideo struct {
	FirstCopy  string `json:"first_copy"`  //主文案（ov章尾激励视频）
	SecondCopy string `json:"second_copy"` //副文案（ov章尾激励视频）
	ImgUrl     string `json:"img_url"`     //图片（ov章尾激励视频）
	ButtonUrl  string `json:"button_url"`  //button图片（ov章尾激励视频）
	ButtonCopy string `json:"button_copy"` //button文案（ov章尾激励视频）
}

func TestStructData(t *testing.T) {
	video := &VideoRewardCps{
		Id:           1,
		Limit:        2,
		Interval:     3,
		GiftType:     "",
		GiftId:       4,
		WheelId:      5,
		UserGroupId:  5,
		SendType:     6,
		SendDesc:     "",
		Amount:       7,
		Amounts:      nil,
		ChapterCount: 8,
		Duration:     9,
	}

	rewardVideo := OvRewardVideo{
		FirstCopy:  "主文案",
		SecondCopy: "副文案",
		ImgUrl:     "www.baidu.com",
		ButtonUrl:  "www.baidu.com",
		ButtonCopy: "www.baidu.com",
	}
	video.OvRewardVideo = rewardVideo
	bytes, _ := json.Marshal(video)
	fmt.Println(string(bytes))
}

func TestNumSort(t *testing.T) {
	numSort := getNumSort()
	for _, v := range numSort {
		fmt.Println(v)
	}
}

func getNumSort() []string {
	nums := []string{"2.3", "1.2", "3.4", "9.7", "5.0"}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}

func TestParseJson(t *testing.T) {
	jsonMsg := `{"first_copy":"主文案","second_copy":"副文案","img_url":"www.baidu.com","button_url":"www.baidu.com","button_copy":"www.baidu.com"}`
	jsonMsg2 := `{"first_copy":"主文案2","second_copy":"副文案2","img_url":"www.baidu2.com","button_url":"www.baidu2.com","button_copy":"www.baidu2.com"}`
	var reward *OvRewardVideo
	_ = json.Unmarshal([]byte(jsonMsg), &reward)
	if reward == nil {
		reward = &OvRewardVideo{}
	}

	var reward2 *OvRewardVideo
	_ = json.Unmarshal([]byte(jsonMsg2), &reward2)
	if reward2 == nil {
		reward = &OvRewardVideo{}
	}

	fmt.Println(reward.FirstCopy)
	fmt.Println(reward2.FirstCopy)
}

func TestEfficiency(t *testing.T) {
	channelId := 0
	start := time.Now().Unix()
	for i := 0; i <= 10000000; i++ {
		//if oppoChannelMap[channelId] {
		//	break
		//}
		if utils.IntListContains(oppoChannels, channelId) {
			break
		}
	}

	end := time.Now().Unix()
	fmt.Println((end - start) / 1000)
}

var oppoChannelMap = map[int]bool{
	164001: true, 141001: true, 107625: true, 141002: true, 107626: true, 141003: true, 107627: true, 119837: true, 119838: true, 119839: true, 107605: true, 107604: true,
	107615: true, 119152: true, 119511: true, 119761: true, 119762: true, 119763: true, 119764: true, 107621: true, 119727: true, 107616: true, 119728: true, 107617: true,
	119725: true, 119693: true, 107614: true, 119691: true, 107612: true, 119507: true, 107610: true, 119545: true, 107611: true, 107607: true, 107150: true, 107136: true,
	116138: true, 116484: true, 116488: true, 116702: true, 116895: true, 119029: true, 109186: true, 116704: true, 116850: true, 108620: true, 116485: true, 107148: true,
	107159: true, 107164: true, 116948: true, 119153: true, 107180: true, 119501: true, 107198: true, 119510: true, 107199: true, 107157: true, 119729: true, 107618: true,
}

var oppoChannels = []int{
	164001, 141001, 107625, 141002, 107626, 141003, 107627, 119837, 119838, 119839, 107605, 107604, 107607, 107618,
	107615, 119152, 119725, 119511, 119761, 119762, 119763, 119764, 107621, 119727, 107616, 119728, 107617, 119729,
	119725, 107615, 119693, 107614, 119691, 107612, 119507, 107610, 119545, 107611, 107604, 107607, 107180, 107604,
	116138, 116484, 116488, 116702, 116895, 119029, 109186, 116704, 116850, 108620, 116485, 107148, 107136, 107150,
	107159, 107164, 116948, 119153, 107180, 119501, 107198, 119510, 107199, 119152, 107157,
}
