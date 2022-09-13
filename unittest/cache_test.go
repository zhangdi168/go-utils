/**
 * @Author: zhangdi
 * @File: cache_test
 * @Version: 1.0.0
 * @Date: 2022/8/28 0:06
 * @Software : GoLand
 */
package unittest

import (
	"fmt"
	"testing"
)

func TestWriteTojson(t *testing.T) {

}

func TestSlice(t *testing.T) {
	s := make([]int, 10)
	s1 := make([]string, 10)
	s2 := make([]map[string]any, 10)

	fmt.Printf("s:%v %v %v \n", s, len(s), cap(s))

	fmt.Printf("s1:%v %v %v  \n", s, len(s1), cap(s1))

	fmt.Printf("s2:%v %v %v  \n", s, len(s2), cap(s2))
}

func TestSclice2(t *testing.T) {
	s := make([]int, 10, 11)
	fmt.Printf("扩容前:%v %v %v \n", s, len(s), cap(s))
	//添加21个
	s = append(s, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3)
	fmt.Printf("扩容后:%v %v %v \n", s, len(s), cap(s))
}
