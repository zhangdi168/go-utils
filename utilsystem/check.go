// Package utilsystem
// @Author: zhangdi
// @File: check
// @Version: 1.0.0
// @Date: 2022/9/13 15:03
package utilsystem

import (
	"runtime"
)

func IsWindowsSystem() bool {
	//runtime.GOARCH 返回当前的系统架构；runtime.GOOS 返回当前的操作系统。
	sysType := runtime.GOOS

	if sysType == "windows" {
		return true
	} else {
		// LINUX系统 或 MAC系统
		return false
	}

	return false
}

func IsMacOs() bool {
	sysType := runtime.GOOS

	if sysType == "darwin" {
		// LMAC系统
		return true
	} else {

		return false
	}

	return false
}

func IsLinuxSystem() bool {
	sysType := runtime.GOOS

	if sysType == "linux" {
		// LINUX系统
		return true
	} else {

		return false
	}

	return false
}
