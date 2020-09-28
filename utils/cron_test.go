package utils

import (
	"fmt"
	"github.com/robfig/cron"
	"testing"
	"time"
)

func Test_Cron(t *testing.T) {
	c := cron.New()
	// 1分钟加载一次
	_ = c.AddFunc("*/10 * * * * *", func() {
		fmt.Println(time.Now())
		fmt.Println("------")
	})
	c.Start()
	select {}
}
