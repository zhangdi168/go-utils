// Package consoles
// @Author: zhangdi
// @File: console_test
// @Version: 1.0.0
// @Date: 2022/9/13 13:03
package unittest

import (
	"github.com/zhangdi168/go-utils/utilcolor"
	"github.com/zhangdi168/go-utils/utilconsole"
	"testing"
)

func TestColor(t *testing.T) {
	error := utilconsole.PrintlnColor("你好", utilcolor.ColorRed, "世界", utilcolor.ColorBlue)
	if error != nil {
		println(error.Error())
	}
}
