package util

import (
	"bytes"
	"log"
	"runtime"
)

func PanicHandler() {
	trace := tracePanic(10)
	if len(trace) > 0 {
		log.Println(string(trace))
	}
}

func tracePanic(kb int) []byte {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine")
	line := []byte("\n")
	stack := make([]byte, kb<<10) // 4kb
	length := runtime.Stack(stack, false)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	return stack
}
