// Package tools
/**@Author: zhangdi
 * @File: build
 * @Version: 1.0.0
 * @Date: 2022/8/31 13:04
 * @Software : GoLand
 */
package tools

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestBuild(t *testing.T) {
	cmd := exec.Command("go", "version")

	stdout, err := os.OpenFile("stdout.log", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln(err)

	}

	defer stdout.Close()
	cmd.Stdout = stdout // 重定向标准输出到文件

	// 执行命令

	if err := cmd.Start(); err != nil {
		log.Println(err)

	}

}
