/**
 * @Author: zhangdi
 * @File: tool_test
 * @Version: 1.0.0
 * @Date: 2022/8/28 0:01
 * @Software : GoLand
 */
package tools

import (
	"fmt"
	"github.com/gonutz/ide/w32"
	"os"
	"testing"
	"time"
)

// HideConsole 隐藏console
func HideConsole() {
	ShowConsoleAsync(w32.SW_HIDE)
}

// ShowConsole 显示console
func ShowConsole() {
	ShowConsoleAsync(w32.SW_SHOW)
}

func ShowConsoleAsync(commandShow uintptr) {
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, commandShow)
		}
	}
}
func TestMain1(t *testing.T) {
	WriteContent("1.txt", os.O_CREATE|os.O_WRONLY, []byte("hello123\n121213"))
}
func TestRead(t *testing.T) {
	s := ReadContent("1.txt", os.O_CREATE|os.O_RDONLY)
	println(string(s))
}

func TestMapTojson(t *testing.T) {
	data := make(map[string]string)
	data["a"] = "1"
	data["b"] = "1"
	data["c"] = "2"

	s := MapToJson(data)

	println(string(s))
}

func TestJsonToMap(t *testing.T) {
	json := `{"a":"1","b":"1","c":"2"}`
	m := JsonToMap([]byte(json))
	print(m)
}

func TestTime(t *testing.T) {
	dateStr := "2022-07-04"
	//dateStr := "2022-08-28 12:00:00"
	var t1 time.Time
	t1 = DateStrToTime(dateStr)
	s, isNextTime := GetSubTime(t1)
	if isNextTime {
		fmt.Printf("距离%v还有%v", dateStr, s)
	} else {
		fmt.Printf("%v已经过去了%v", dateStr, s)
	}

}

// TestTicker
//
//	@Description: 定时器
//	@param t
func TestTicker(t *testing.T) {
	var ch chan int
	//定时任务
	ticker := time.NewTicker(time.Second * 2)
	go func() {
		for {
			timers := <-ticker.C
			fmt.Println(timers)
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
	}()

	//用于停止ticker
	go func() {

	}()
	//读取管道，由于我们没往管道放数据，也就是会阻塞当前进程，进制其结束
	<-ch
	println(2)
}

// TestDayTimeDo
//
//	@Description: 测试每天定时执行
//	@param t
func TestDayTimeDo(t *testing.T) {
	StartTimer("20:33", doSomething)
}

func doSomething() {
	println("我被执行了")
}

// TestSub
//
//	@Description: 时间差值计算
//	@param t
func TestSub(t *testing.T) {
	//每天的12点执行
	Hour, Min := 20, 26
	now := time.Now()

	//格式化第一次运行的具体时间
	var firstRunDate time.Time
	if now.Hour() > Hour {
		firstRunDate = now.Add(24 * time.Hour)
	} else {
		firstRunDate = now
	}
	FirstRunTimeStr := fmt.Sprintf("%v-%02d-%02d %02d:%02d:00", firstRunDate.Year(), firstRunDate.Month(), firstRunDate.Day(), Hour, Min)
	FirstRunTime, err := time.ParseInLocation("2006-01-02 15:04:05", FirstRunTimeStr, time.Local)
	if err != nil {
		println("转换失败")
	}
	//计算时间差值
	difference := FirstRunTime.Sub(now)

	//启动第一个ticker
	ticker1 := time.NewTicker(difference)

	<-ticker1.C   //读取ticker的管道，但是时间没到会处于阻塞状态
	doSomething() //执行我们的定时任务业务逻辑，能执行到说明ticker1往管道写入数据了
	println("第一次执行完毕，即将开始循环执行...")
	for {
		ticker := time.NewTicker(24 * time.Hour)
		<-ticker.C    //读取ticker的管道，但是时间没到会处于阻塞状态
		doSomething() //执行我们的定时任务业务逻辑，能执行到说明ticker1往管道写入数据了
	}
	print("执行结束")
}

func TestHex(t *testing.T) {
	s := "zhnagdi"

	r := HexToBytes(s)
	println(r)
	t.Logf("16进制 %v", r)

}
