package main

import (
	"fmt"
	"github.com/electricbubble/go-toast"
	"os/exec"
	"strings"
)

func macOSOpen(path string) {
	iterm2Exit := exitIterm2App()
	if iterm2Exit {
		_ = toast.Push("iTerm.app 进入 路径:"+path, toast.WithTitle("go2terminal"))
		// 走iterm2的逻辑
		command := exec.Command("open", "-a", "iTerm.app", path)
		err := command.Run()
		if err != nil {
			return
		}
	} else {
		_ = toast.Push("Terminal.app 进入 路径:"+path, toast.WithTitle("go2terminal"))
		// 原生terminal
		command := exec.Command("open", "-a", "Terminal.app", path)
		err := command.Run()
		if err != nil {
			return
		}
	}
}

func exitIterm2App() bool {
	appBundleID := "com.googlecode.iterm2"
	cmd := exec.Command("mdfind", fmt.Sprintf("kMDItemCFBundleIdentifier == '%s'", appBundleID))
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("未安装应用程序 Iterm2\n")
	} else {
		if len(strings.TrimSpace(string(output))) > 0 {
			fmt.Printf("已安装应用程序 Iterm2\n")
			return true
		} else {
			fmt.Printf("未安装应用程序 Iterm2\n")
		}
	}
	return false
}
