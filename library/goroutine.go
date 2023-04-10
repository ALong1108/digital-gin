package library

import (
	"bytes"
	"runtime"
	"strings"
)

var l = len("goroutine ")

func GoroutineID() string {
	var buf [32]byte
	n := runtime.Stack(buf[:], false)

	b := bytes.NewBuffer(buf[l:n])
	s, _ := b.ReadString(' ')

	return strings.TrimSpace(s)
}
