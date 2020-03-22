/**
 * @Author: liubaoshuai3
 * @Date: 2020/3/21 21:01
 * @File: circularqueue
 * @Description:
 */

package tools

// struct circular queue
type CircularQueue struct {
	q []interface{}
	h int
	t int
	c int
}

// new a circular queue
func NewCircularQueue(n int) *CircularQueue {
	if n <= 0 {
		return nil
	}
	return &CircularQueue{
		q: make([]interface{}, n),
		h: 0,
		t: 0,
		c: n,
	}
}

// if the queue is empty or not
func (c *CircularQueue) IsEmpty() bool {
	if c.h == c.t {
		return true
	}
	return false
}

// if the queue is full or not
// (tail+1)%capacity==head
func (c *CircularQueue) IsFull() bool {
	if (c.t + 1) % c.c == c.h {
		return true
	}
	return false
}

// push a value to queue
//func (c *CircularQueue)  {
//
//}