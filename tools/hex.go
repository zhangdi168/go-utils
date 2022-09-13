// Package tools
/**@Author: zhangdi
 * @File: hex
 * @Version: 1.0.0
 * @Date: 2022/9/2 10:16
 * @Software : GoLand
 */
package tools

import "strings"

// HexToBytes  16进制转换
// / @param hexStr 16进制字符串
// / @return []byte
func HexToBytes(hexStr string) []byte {
	hexStr = strings.ToUpper(hexStr)
	var bytes []byte
	for i := 0; i < len(hexStr); i += 2 {
		bytes = append(bytes, HexCharToByte(hexStr[i:i+1])<<4|HexCharToByte(hexStr[i+1:i+2]))
	}
	return bytes
}

// HexCharToByte 单个字符转换为byte
// / @param hexChar 单个字符
// / @return byte
func HexCharToByte(hexChar string) byte {
	if len(hexChar) != 1 {
		return 0
	}
	if hexChar >= "0" && hexChar <= "9" {
		return byte(hexChar[0] - '0')
	}
	if hexChar >= "A" && hexChar <= "F" {
		return byte(hexChar[0] - 'A' + 10)
	}
	return 0
}
