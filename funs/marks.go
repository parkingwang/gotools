package funs

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// RequestID 生成request_id
func RequestID() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.FormatInt(time.Now().UnixNano(), 10) + fmt.Sprintf(":%d", r.Intn(10000))
}
