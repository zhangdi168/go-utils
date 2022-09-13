// Package utilcache
// @Author: zhangdi
// @Version: 1.0.0
// @Date: 2022/9/13 10:05
package utilcache

var _ ICache = &fileCache{}

type ICache interface {
	Set(k string, v string)
	TimeSet(k string, v string, exTime int64)
	Get(k string) string
}

var GlobalCache ICache
