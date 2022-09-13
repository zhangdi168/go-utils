// Package cache /**
package cache

import (
	error2 "WechatSend/utils/tools"
	"fmt"
	"time"
)

// InitGoCache
// @Description: 初始化go-cache
// @param timeOutMin
// @param timeClearMin
func InitGoCache(timeOutMin int, timeClearMin int) {
	GlobalCache = &goCache{
		TimeOutMin:   time.Duration(timeOutMin),
		TimeClearMin: time.Duration(timeClearMin),
		cache:        cache.New(time.Duration(timeOutMin)*time.Minute, time.Duration(timeClearMin)*time.Minute),
	}
}

type goCache struct {
	TimeOutMin   time.Duration
	TimeClearMin time.Duration
	cache        *cache.Cache
}

// Get
// @Description: 读取缓存
// @receiver c
// @param k
// @return any
// @return bool
func (c goCache) Get(k string) string {
	defer error2.GetErrorRecover() //捕获异常
	r, flag := c.cache.Get(k)
	if !flag {
		panic("goCache get失败")
		return ""
	}
	str := fmt.Sprintf("%s", r)
	return str
}

// Set
// @Description: 写入缓存，并设置超时时间
// @receiver c
// @param k
// @param v
// @param exTime
func (c goCache) Set(k string, v string) {
	defer error2.GetErrorRecover() //捕获异常

	// 设置缓存值并使用默认时间
	c.cache.Set(k, v, cache.DefaultExpiration)
}

// TimeSet Set
// @Description: 写入缓存，并设置超时时间
// @receiver c
// @param k
// @param v
// @param exTime 过期时间-秒
func (c goCache) TimeSet(k string, v string, exTime int64) {
	defer error2.GetErrorRecover() //捕获异常

	// 设置缓存值并带上过期时间
	c.cache.Set(k, v, time.Duration(exTime)*time.Second)
}
