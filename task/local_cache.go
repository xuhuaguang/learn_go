package task

import (
	"go.uber.org/atomic"
	"time"
)

const (
	OneDayTime       = 24 * time.Hour
	defaultCacheTime = 65 * time.Second
	DefaultMaxDspQps = float32(10000)
)

type SspDb struct {
	SspUsrLevelCount *atomic.Int32
}

func NewSspDb() *SspDb {
	return &SspDb{SspUsrLevelCount: atomic.NewInt32(0)}
}
