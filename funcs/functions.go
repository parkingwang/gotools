package funcs

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// RequestID 生成request_id; 获取堆栈信息会影响性能，建议debug时用。
func RequestID() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.FormatInt(time.Now().UnixNano(), 10) + fmt.Sprintf(":%d", r.Intn(10000))
}

// GID 获取 goroutine ID
func GID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
