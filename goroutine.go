package logger

import (
	"bytes"
	"runtime"
	"strings"
)

// l = len("goroutine ")
const l = 10

func GoID() string {
	var buf [32]byte
	n := runtime.Stack(buf[:], false)

	b := bytes.NewBuffer(buf[l:n])
	s, _ := b.ReadString(' ')

	return strings.TrimSpace(s)
}
