// Package console
// @Author: zhangdi
// @File: color
// @Version: 1.0.0
// @Date: 2022/9/13 10:39
package console

import (
	"github.com/vua/vfmt"
	"go-utils/consts"
)

// PrintlnRed
// / @Description: 输出颜色
// / @param title 标题
// / @param colorContent 红色内容
func PrintlnRed(title string, RedContent string) {
	vfmt.Printf("%s @[::%s]", title, RedContent, consts.ColorRed)
}
