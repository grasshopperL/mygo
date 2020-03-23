/**
 * @Author: liubaoshuai3
 * @Date: 2020/3/21 21:01
 * @File: circularqueue
 * @Description:
 */

package tools

import "fmt"

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
func (c *CircularQueue) Enqueue(v interface{}) bool {
	if t := c.IsFull(); t {
		return false
	}
	c.q[c.t] = v
	c.t = (c.t + 1) % c.c
	c.c++
	return true
}

// pop a value from queue
func (c *CircularQueue) Dequeue() interface{} {
	if t := c.IsEmpty(); t {
		return false
	}
	v := c.q[c.h]
	c.h = (c.h + 1) % c.c
	return v
}

// over write String()
func (c *CircularQueue) String() string {
	if t := c.IsEmpty(); t {
		return "empty queue"
	}
	result := "head"
	var i = c.h
	for true {
		result += fmt.Sprintf("<-%+v", c.q[i])
		i = (i + 1) % c.c
		if i == c.t {
			break
		}
	}
	result += "<-tail"
	return result
}