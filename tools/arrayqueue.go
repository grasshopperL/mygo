/**
 * @Author: liubaoshuai3
 * @Date: 2020/3/21 20:37
 * @File: arraylinked
 * @Description:
 */

package tools

import "fmt"

// array queue struct
type ArrayQueue struct {
	q []interface{}
	h, t int
	c int
}

// init an array queue
func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		q: make([]interface{}, n),
		h: 0,
		t: 0,
		c: n,
	}
}

// insert a node to queue
func (a *ArrayQueue) Enqueue(v interface{}) bool {
	if a.t == a.c {
		return false
	}
	a.q[a.t] = v
	a.t++
	return true
}

// pop a node from queue
func (a *ArrayQueue) Dequeue() interface{} {
	if a.h == a.t {
		return nil
	}
	v := a.q[a.h]
	a.h++
	return v
}

// overload string
func (a *ArrayQueue) String() string {
	if a.h == a.t {
		return "empty queue"
	}
	result := "head"
	for i := a.h; i <= a.t-1; i++ {
		result += fmt.Sprintf("<-%+v", a.q[i])
	}
	result += "<-tail"
	return result
}