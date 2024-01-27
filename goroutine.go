package logger

import (
	"bytes"
	"runtime"
	"strings"
)

const (
	l       = 10 // l = len("goroutine ")
	goIDKey = "GoID"
)

func GoID() string {
	var buf [32]byte
	n := runtime.Stack(buf[:], false)

	b := bytes.NewBuffer(buf[l:n])
	s, _ := b.ReadString(' ')

	return strings.TrimSpace(s)
}
