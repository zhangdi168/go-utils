// Package tools /**

package tools

import (
	"log"
	"os"
	"sync"
)

// 创建一个读写锁
var lock sync.RWMutex

func logWrite(content1 any) {
	go func(content any) {
		lock.Lock()         //加锁
		defer lock.Unlock() //解锁

		logFile, err := os.OpenFile("logs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 755)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(logFile)
		log.SetPrefix("[WechatSend]")
		log.Print(content)
	}(content1)

}
