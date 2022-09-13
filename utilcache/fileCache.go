// Package utilcache
// @Author: zhangdi
// @Version: 1.0.0
// @Date: 2022/9/13 10:05
package utilcache

import (
	"encoding/json"
	"github.com/zhangdi168/go-utils/utilerror"
	"github.com/zhangdi168/go-utils/utilfile"
	"github.com/zhangdi168/go-utils/utiljson"
	"os"
	"time"
)

func InitFileCache() {
	data := readJson()
	if data == nil {
		data = make(map[string]string)
	}

	GlobalCache = &fileCache{
		Data:           data,
		DefaultSeconds: 7200,
	}
}

type FileCacheItem struct {
	ExpirationTime int64
	Value          string
}

type fileCache struct {
	Data           map[string]string
	DefaultSeconds int64 // 默认过期秒数
}

func (f fileCache) Set(k string, v string) {
	f.TimeSet(k, v, f.DefaultSeconds)
}

func (f fileCache) TimeSet(k string, v string, exTime int64) {
	item := FileCacheItem{
		ExpirationTime: time.Now().Unix() + exTime,
		Value:          v,
	}
	jsonStr, err := json.Marshal(item)
	if err != nil {
		return
	}
	f.Data[k] = string(jsonStr)
	writeJson(f.Data)
}

func (f fileCache) Get(k string) string {
	defer utilerror.GetErrorRecover()
	item, ok := f.Data[k]
	if !ok {
		//不存在key
		return ""
	}
	info := &FileCacheItem{}
	err := json.Unmarshal([]byte(item), info)
	if err != nil {
		return ""
	}
	//缓存过期 删除缓存
	if time.Now().Unix() > info.ExpirationTime {
		delete(f.Data, k)
		writeJson(f.Data)
		return ""
	}
	return info.Value
}

// readJson
// @Description: 读取文件缓存中的所有数据
// @return map[string]string
func readJson() map[string]string {
	jsonStr := utilfile.ReadContent("cache.json", os.O_CREATE|os.O_RDONLY)
	if jsonStr == nil {
		return nil
	}
	data := utiljson.JsonToMap(jsonStr)
	return data
}

// writeJson
// @Description: 将map的数据写入文件缓存
// @return int 返回写入成功的字节数
func writeJson(data map[string]string) int {
	jsonBytes := utiljson.MapToJson(data)
	if jsonBytes == nil {
		return 0
	}
	//O_TRUNC覆盖写入
	n := utilfile.WriteContent("cache.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, jsonBytes)
	return n

}
