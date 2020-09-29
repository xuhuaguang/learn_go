package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gomodule/task"
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
