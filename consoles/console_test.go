// Package consoles
// @Author: zhangdi
// @File: console_test
// @Version: 1.0.0
// @Date: 2022/9/13 13:03
package consoles

import (
	"github.com/zhangdi168/go-utils/consts"
	"testing"
)

func TestColor(t *testing.T) {
	error := PrintlnColor("你好", consts.ColorRed, "世界", consts.ColorBlue)
	if error != nil {
		println(error.Error())
	}
}
