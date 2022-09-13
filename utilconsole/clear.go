// Package utilconsole
// @Author: zhangdi
// @File: clear
// @Version: 1.0.0
// @Date: 2022/9/13 15:02
package utilconsole

import (
	"github.com/zhangdi168/go-utils/utilsystem"
	"os"
	"os/exec"
)

// Clear
//
//	@Description: 清空控制台
func Clear() {
	if utilsystem.IsWindowsSystem() {
		clearWindows()
	} else {
		clearLinux()
	}
}

// clearWindows
//
//	@Description: 清空控制台-windows
func clearWindows() {
	//执行clear指令清除控制台
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// clearLinux
//
//	@Description: 清空控制台-linux或mac
func clearLinux() {
	//执行clear指令清除控制台
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
