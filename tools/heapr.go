/**
 * @Author: liubaoshuai3
 * @Date: 2020/4/20 22:48
 * @File: heapr
 * @Description:
 */

package tools

import (
	"hash/crc32"
	"sort"
)

// implement the interface do a heap
type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop()
}

// init a heap ???
func Init(h Interface)  {
	n := h.Len()
	for i := n/2 -1; i >= 0; i-- {
		down(h, i, n)
	}
}

// push a value to heap
func Push(h Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len() - 1)
}