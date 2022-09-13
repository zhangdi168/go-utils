// Package consoles
// @Author: zhangdi
// @File: color
// @Version: 1.0.0
// @Date: 2022/9/13 10:39
package consoles

import (
	"errors"
	"github.com/vua/vfmt"
	"github.com/zhangdi168/go-utils/consts"
)

// PrintlnRed
// / @Description: 输出颜色
// / @param title 标题
// / @param colorContent 红色内容
func PrintlnRed(title string, RedContent string) {
	vfmt.Printf("%s @[%s::%s]", title, RedContent, consts.ColorRed)
}

// PrintlnGreen
// / @Description: 输出颜色
// / @param title 标题
// / @param colorContent 绿色内容
func PrintlnGreen(title string, RedContent string) {
	vfmt.Printf("%s @[%s::%s]", title, RedContent, consts.ColorGreen)
}

// PrintlnBlue
// / @Description: 输出颜色
// / @param title 标题
// / @param colorContent 蓝色内容
func PrintlnBlue(title string, RedContent string) {
	vfmt.Printf("%s @[%s::%s]", title, RedContent, consts.ColorBlue)
}

// PrintlnColor
//
//	@Description: 控制台带颜色文本打印通用方法
//	@param params ：格式如下：PrintlnColor("内容",颜色代码，"内容2"，颜色代码)
func PrintlnColor(params ...string) error {
	if len(params)%2 != 0 {
		return errors.New("参数数量必须是偶数个")
	}
	line := ""
	color := ""
	content := ""
	for i, param := range params {
		if i%2 == 0 {
			//颜色代码
			content = param

		} else {
			color = param
			line += "@[" + content + "::" + color + "]"
			color = ""
			content = ""
		}
	}
	vfmt.Println(line)
	return nil
}
