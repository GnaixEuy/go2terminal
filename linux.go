package main

import (
	"github.com/electricbubble/go-toast"
)

func linuxOpen(path string) {
	_ = toast.Push("当前操作系统暂未支持", toast.WithTitle("go2terminal"))
}
