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
//	@Description: 跳转到内置浏览器，仅windows下生效
//	@param uri
//	@return error
func OpenBrowser(uri string) error {
	if utilsystem.IsWindowsSystem() {
		cmd := exec.Command("cmd", "/C", "start "+uri)
		return cmd.Run()
	}
	return nil
}
