// Package utilerror  go-utils
// @Author: zhangdi
// @Version: 1.0.0
// @Date: 2022/9/13 10:05
package utilerror

import "github.com/zhangdi168/go-utils/utillog"

// GetErrorRecover
// @Description: 捕获异常
func GetErrorRecover() {
	if err := recover(); err != nil {
		println(err)
		utillog.LogFile(err)
	}
}
