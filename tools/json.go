// Package tools /**
package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// MapToJson
// @Description: map转json字符串
// @param map1
// @return string
func MapToJson(map1 map[string]string) []byte {
	bytesData, err := json.Marshal(map1)
	if len(bytesData) == 2 && bytesData[0] == 124 && bytesData[1] == 125 {
		return nil
	}
	if err != nil {
		return nil
	}
	return bytesData
}

// JsonToMap
// @Description: json字节流转map
// @param jsonBytes
func JsonToMap(jsonBytes []byte) map[string]string {
	map2 := make(map[string]string)
	err := json.Unmarshal(jsonBytes, &map2)
	if err != nil {
		return nil
	}
	return map2
}

// StructToMap
// @Description: 将结构体转map
// @param content
// @return map[string]string
func StructToMap(content string) map[string]string {
	var name map[string]string
	if marshalContent, err := json.Marshal(content); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber() // 设置将float64转为一个number
		if err := d.Decode(&name); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range name {
				name[k] = v
			}
		}
	}
	return name
}
