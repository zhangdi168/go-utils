// Package utilconsole  go-utils
// @Author: zhangdi
// @Version: 1.0.0
// @Date: 2022/9/13 10:05
package utilconsole

import (
	"github.com/zhangdi168/go-utils/utilsystem"
	"os/exec"
)

// OpenBrowser
//
//	@Description: 跳转到内置浏览器
//	@param uri
//	@return error
func OpenBrowser(uri string) error {
	if utilsystem.IsWindowsSystem() {
		return jumpByWin(uri)
	} else if utilsystem.IsMacOs() {
		return jumpByMac(uri)
	} else {
		return jumpByLinux(uri)
	}
}

func jumpByWin(uri string) error {
	if utilsystem.IsWindowsSystem() {
		cmd := exec.Command("cmd", "/C", "start "+uri)
		return cmd.Run()
	}
	return nil
}

func jumpByMac(uri string) error {
	cmd := exec.Command("open", uri)
	return cmd.Start()
}

func jumpByLinux(uri string) error {
	cmd := exec.Command("xdg-open", uri)
	return cmd.Start()
}
