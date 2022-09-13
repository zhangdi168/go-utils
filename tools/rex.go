// Package tools /**
package tools

import "regexp"

// ParseDate
//
//	@Description: 解析时间 支持2022-08-29 12:00:01 和 2022-08-29两种格式
//	@param dateStr
//	@return string
func ParseDate(dateStr string) string {

	rex := regexp.MustCompile(`\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}`)
	strs := rex.FindAllString(dateStr, -1)
	if len(strs) != 0 {
		return strs[0]
	}

	rex = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	strs = rex.FindAllString(dateStr, -1)
	if len(strs) != 0 {
		return strs[0]
	}
	return ""
}
