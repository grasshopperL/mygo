/**
 * @Author: liubaoshuai3
 * @Date: 2020/3/24 20:24
 * @File: heap
 * @Description:
 */

package tools

// the struct of heap
// top-max heap and heapify from down to up
type Heap struct {
	h []int
	c int
	n int
}

// new a heap
func NewHeap(c int) *Heap {
	return &Heap{
		h: make([]int, c),
		c: c,
		n: 0,
	}
}

// insert a value to heap
func (h *Heap) Insert(v int) bool {
	if h.c == h.n {
		return false
	}
	h.c++
	h.h[h.c] = v
	i := h.c
	p := i / 2
	for p > 0 && h.h[p] < h.h[i]{
		h.h[p], h.h[i] = h.h[i], h.h[p]
		i = p
		p = i / 2
	}
	return true
}

// hwo to use two return value to finish this methods
//func (h *Heap) RemoveMax() int {
//	if h.c == 0 {
//
//	}
//}