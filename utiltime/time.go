// Package utiltime Package go-utils
// @Author: zhangdi
// @Version: 1.0.0
// @Date: 2022/9/13 10:05
package utiltime

import (
	"fmt"
	"github.com/zhangdi168/go-utils/utilerror"
	"math"
	"strconv"
	"strings"
	"time"
)

/*
时间常量
*/
const (

	// SecondsPerMinute 定义每分钟的秒数
	SecondsPerMinute = 60
	// SecondsPerHour 定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	// SecondsPerDay 定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

// ResolveTime
//
//	@Description: 将秒转换成具体的天数、时、分钟数
//	@param seconds
//	@return day
//	@return hour
//	@return minute
//	@return second
func ResolveTime(seconds int) (day int, hour int, minute int, second int) {
	second = ((seconds % 86400) % 3600) % 60
	//每分钟秒数
	minute = ((seconds % 86400) % 3600) / SecondsPerMinute
	//每小时秒数
	hour = (seconds % 86400) / SecondsPerHour
	//每天秒数
	day = seconds / SecondsPerDay
	return
}

// GetSubTime
//
//	@Description: 计算某个时间与当前的时间间隔天 时 分 秒
//	@param t1
//	@return string
//	@return bool 是否是倒计时，true表示传入时间大于当前时间，false表示传入时间小于当前时间
func GetSubTime(t1 time.Time) (string, bool) {

	now := time.Now()
	nowUnix := now.Unix()
	dateUnix := t1.Unix()
	if dateUnix < 0 {
		return "", false
	}
	isNextTime := true
	//计算两个时间间隔
	tmInvalid := int(dateUnix - nowUnix)
	if tmInvalid < 0 {
		//时间已经过了
		isNextTime = false
		tmInvalid = int(math.Abs(float64(tmInvalid)))
		//return fmt.Sprint("时间总在悄然流逝", t1.Format("2006-1-2 15:04:05"), "已经过去了\n但人生是一场没有退路的旅程,要活在当下，着眼未来，加油！")
	}

	days, hours, mins, secs := ResolveTime(tmInvalid)

	return fmt.Sprintf("%v天%v时%v分%v秒", days, hours, mins, secs), isNextTime
}

// DateStrToTime 将字符串转成go语言支持时间数据
//
//	@Description:支持格式：2022-07-04 或2022-07-04 12:00:00
//	@param dateStr
//	@return time.Time
func DateStrToTime(dateStr string) time.Time {
	defer utilerror.GetErrorRecover()
	var t1 time.Time
	if strings.Contains(dateStr, ":") {
		t1, _ = time.ParseInLocation("2006-01-02 15:04:05", dateStr, time.Local)
	} else {
		t1, _ = time.ParseInLocation("2006-01-02 15:04:05", dateStr+" 00:00:00", time.Local)
	}
	return t1
}

// StartTimer 定时任务
//
//		@Description: 定时任务每天的H点M分执行函数f
//		@param HM 时间，格式如 13:14表示每天13点14分执行函数funDoSomething
//		@param funDoSomething 定点执行的函数
//	   @return chan  返回一个chan,我们需要齐停止是 写入true : stopTicker <- true
func StartTimer(HM string, funDoSomething func()) chan bool {

	if !strings.Contains(HM, ":") {
		return nil
	}
	HMS := strings.Split(HM, ":")
	if len(HMS) < 2 {
		return nil
	}
	Hour, err := strconv.Atoi(HMS[0])
	if err != nil {
		return nil
	}
	Min, err := strconv.Atoi(HMS[1])
	if err != nil {
		return nil
	}

	//定义一个channel来接收是否停止
	stopTicker := make(chan bool)
	go func() {

		fmt.Printf("定时任务启动 每天%v执行", HM)
		now := time.Now()
		//格式化第一次运行的具体时间
		var firstRunDate time.Time
		if now.Hour() > Hour {
			firstRunDate = now.Add(24 * time.Hour)
		} else if now.Hour() == Hour && now.Minute() >= Min {
			//小时相同 则比较分钟
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

		//启动第一个和第二个ticker，第二个ticker固定时长为24h
		ticker1 := time.NewTicker(difference)
		ticker2 := time.NewTicker(24 * time.Hour)
		defer ticker2.Stop()
		defer ticker1.Stop()

		for {
			//select 在执行过程中，必须命中其中的某一分支。否则就会一直等待
			select {
			case date1 := <-ticker1.C:
				go funDoSomething()
				fmt.Printf("第一次执行于:%v", date1)
				ticker1.Stop()                //ticker1仅执行一次就停止
				ticker2.Reset(24 * time.Hour) //此刻开始ticker2进入循环
			case date2 := <-ticker2.C:
				go funDoSomething()
				fmt.Printf("执行定时任务于:%v", date2)
			case isStop := <-stopTicker: //停止定时器信号
				if isStop {
					println("收到定时器停止命令，即将停止定时器")
					ticker1.Stop()
					ticker2.Stop()
					return
				}

			}
		}
	}()

	return stopTicker

}
