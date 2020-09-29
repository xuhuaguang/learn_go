package task

import (
	"fmt"
	"github.com/robfig/cron"
)

type Loader interface {
	Run()
}

func StartAllDataLoader(loaderIntervalSecond int) {
	var loaders = []Loader{
		&StrategyIdLoader{interval: loaderIntervalSecond},
		&AdSlotLoader{interval: 15},
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
	_ = c.AddFunc(fmt.Sprintf("*/%d * * * * *", l.interval), func() {
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
	_ = c.AddFunc(fmt.Sprintf("*/%d * * * * *", l.interval), func() {
		l.Load()
	})
	c.Start()
}

func (l *AdSlotLoader) Load() {
	fmt.Println("AdSlotLoader start ---,time is :", l.interval)
}
