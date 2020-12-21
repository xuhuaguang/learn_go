package test

import (
	"fmt"
	"learn_go/utils"
	"math/rand"
	"testing"
	"time"
)

var res []*ResponseInfo

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
}

type ResponseInfo struct {
	ID         string `json:"id"`
	CreativeId string `json:"cid"`
	Price      int    `json:"price"`
	Adm        string `json:"adm"`
}
