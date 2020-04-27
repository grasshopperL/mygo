/**
 * @Author: liubaoshuai3
 * @Date: 2020/4/20 22:48
 * @File: heapr
 * @Description:
 */

package tools

import (
	"sort"
)

// implement the interface do a heap
type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
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

//  pop a value
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

// remove one value
func Remove(h Interface, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

// I don't know why do like this
func Fix(h Interface, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

// exchange from j to j's parent
func up(h Interface, j int) {
	for  {
		i := (j - 1) / 2
		if i == j || !h.Less(i, j) {
			break
		}
		h.Swap(i, j)
		j = i
	}

}

// exchange from i to i's child
func down(h Interface, i0, n int) bool {
	i := i0
	for  {
		j1 := 2 * i + 1
		if j1 < 0 || j1 > n {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && !h.Less(j2, j1) {
			j = j2
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}