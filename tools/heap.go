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
func (h *Heap) RemoveMax() bool {
	if h.c == 0 {
		return false
	}
	h.h[1], h.h[h.c] = h.h[h.c], h.h[1]
	h.c--
	hUpToDown(h.h, h.c)
	return true
}

// heapify
func hUpToDown(h []int, c int) {
	for i := 1; i < c / 2;  {
		mI := i
		if h[i] < h[i * 2] {
			mI = i * 2
		}
		if i * 2 + 1 <= c && h[mI] < h[i * 2 + 1] {
			mI = i * 2 + 1
		}
		if mI == i {
			break
		}
		h[i], h[mI] = h[mI], h[i]
		i = mI
	}
}