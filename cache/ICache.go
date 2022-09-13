// Package cache /**
package cache

var _ ICache = &goCache{}
var _ ICache = &fileCache{}

type ICache interface {
	Set(k string, v string)
	TimeSet(k string, v string, exTime int64)
	Get(k string) string
}

var GlobalCache ICache
