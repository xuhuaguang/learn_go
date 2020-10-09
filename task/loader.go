package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

type Loader interface {
	Run()
}

func StartAllDataLoader(loaderIntervalSecond int) {
	var loaders = []Loader{
		&StrategyIdLoader{interval: loaderIntervalSecond},
		&AdSlotLoader{interval: 15},
		&TestCronTask{interval: 1},
	}

	for _, v := range loaders {
		go v.Run()
	}
}

type StrategyIdLoader struct {
	interval int
}

func (l *StrategyIdLoader) Run() {
	l.Load()
	c := cron.New()
	_, _ = c.AddFunc(fmt.Sprintf("*/%d * * * * *", l.interval), func() {
		l.Load()
	})
	c.Start()
}

func (l *StrategyIdLoader) Load() {
	fmt.Println("StrategyIdLoader start ---,time is ", l.interval)
}

type AdSlotLoader struct {
	interval int
}

func (l *AdSlotLoader) Run() {
	l.Load()
	c := cron.New()
	_, _ = c.AddFunc(fmt.Sprintf("*/%d * * * * *", l.interval), func() {
		l.Load()
	})
	c.Start()
}

func (l *AdSlotLoader) Load() {
	fmt.Println("AdSlotLoader start ---,time is :", l.interval)
}

type TestCronTask struct {
	interval int
}

//https://crontab.guru/#*/1_*_*_*_*
func (l *TestCronTask) Run() {
	l.Load()
	c := cron.New()
	_, _ = c.AddFunc("*/1 * * * *", func() { //每一分钟执行一次
		l.Load()
	})
	c.Start()
}

func (l *TestCronTask) Load() {
	fmt.Println("TestCronTask start ---,time is :", l.interval, time.Now())
}
