//  Package go-utils
// @Author: zhangdi
// @Version: 1.0.0
// @Date: 2022/9/13 10:05

package utillog

import (
	"log"
	"os"
	"sync"
)

// 创建一个读写锁
var lock sync.RWMutex

var prefix = ""
var path = "logs.log"

// SetPrefix
//
//	@Description: 设置全局日志前缀
//	@param prefix_
func SetPrefix(prefix_ string) {
	prefix = prefix_
}

// SetPath
//
//	@Description: 设置全局日志保存路径
//	@param path_
func SetPath(path_ string) {
	path = path_
}

// LogFile
//
//	@Description: 将日志写入到文件中去
//	@param content1
func LogFile(content1 any) {
	go func(content any) {
		lock.Lock()         //加锁
		defer lock.Unlock() //解锁
		logFile, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 755)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(logFile)
		log.SetPrefix(prefix)
	}(content1)
}
