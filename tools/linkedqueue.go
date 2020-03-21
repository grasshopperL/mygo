/**
 * @Author: liubaoshuai3
 * @Date: 2020/3/20 22:26
 * @File: LinkedQueue
 * @Description:
 */

package tools

import "fmt"

// queue based on single list
type LinkedQueue struct {
	H *Node
	T *Node
	L int
}

// init a queue
func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{
		H: nil,
		T: nil,
		L: 0,
	}
}

// insert a node to queue
func (l *LinkedQueue) Enqueue(v interface{}) bool {
	temp := NewLinkNode(v)
	if l.H == nil {
		l.H = temp
		l.T = temp
	}
	l.T.N = temp
	l.T = temp
	return true
}

// pop a node from queue
func (l *LinkedQueue) Dequeue() interface{} {
	if l.H == nil {
		return nil
	}
	v := l.H.GetVal()
	l.H = l.H.GetNext()
	l.L--
	return v
}

// overload string
func (l *LinkedQueue) String() string {
	if l.H == nil {
		return "empty queue"
	}
	result := "head<-"
	for cur := l.H; cur != nil; cur = cur.N {
		result += fmt.Sprintf("<-%+v", cur.V)
	}
	result += "<-tail"
	return result
}