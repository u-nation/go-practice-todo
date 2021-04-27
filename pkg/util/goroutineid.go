package util

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// GoroutineID は呼び出し側goroutineのスタックトレースをbufにフォーマットし、bufに書き込まれたバイト数を取得する
func GoroutineID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]

	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutineid: %v", err))
	}

	return id
}
