/**
 * @Author: liubaoshuai3
 * @Date: 2020/3/20 22:26
 * @File: LinkedQueue
 * @Description:
 */

package tools

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
//func (l *LinkedQueue) Enqueue(v interface{}) {
//	
//}