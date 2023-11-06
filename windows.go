package main

import "github.com/electricbubble/go-toast"

func windowsOpen(path string) {
	//windows通过系统弹出通知 当前操作系统暂未支持
	_ = toast.Push("当前操作系统暂未支持", toast.WithTitle("go2terminal"))
}
