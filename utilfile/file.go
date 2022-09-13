// Package utilfile  go-utils
// @Author: zhangdi
// @Version: 1.0.0
// @Date: 2022/9/13 10:05
package utilfile

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ReadContent(path string, flag int) []byte {
	f, err := os.OpenFile(path, flag, 755)
	if err != nil {
		fmt.Println("read fail")
		return make([]byte, 0)
	}
	defer f.Close()

	// 把file读取到缓冲区中
	var chunk []byte
	buf2 := make([]byte, 1024)

	for {
		// 从file读取到buf中
		n, err := f.Read(buf2)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return make([]byte, 0)
		}
		// n=0说明Read不到东西了
		if n == 0 {
			break
		}
		// 读取到最终的缓冲区中
		chunk = append(chunk, buf2[:n]...)
	}

	return chunk
}

// WriteContent
// @Description: 使用 File(Write,WriteString) 写入文件
// @param path
// @param flag
// @param content
// @return int
func WriteContent(path string, flag int, content []byte) int {
	file, err := os.OpenFile(path, flag, 755)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	n, err := file.Write(content) //写入文件(字节数组)
	if err != nil {
		log.Fatal(err.Error())
	}
	return n

}
