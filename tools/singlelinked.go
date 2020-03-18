/**
 * @Author: liubaoshuai3
 * @Date: 2020/3/18 20:57
 * @File: singlelinked
 * @Description:
 */

package tools

// node struct
type Node struct {
	N *Node
	V interface{}
}

// link struct
type LinkedList struct {
	H *Node
	L int
}

// create a new node
func NewLinkNode(v interface{}) *Node {
	return &Node{
		N: nil,
		V: v,
	}
}

// create a new list
func NewLinkedList() *LinkedList {
	return &LinkedList{NewLinkNode(0), 0}
}

// get the next node
func (n *Node) GetNext() *Node {
	return n.N
}

// get the node value
func (n *Node) GetVal() interface{} {
	return n.V
}

// insert after the given node
func (l *LinkedList) InsertAfter(n *Node, v interface{}) bool {
	if n == nil {
		return false
	}
	h := l.H
	for h != nil {
		if h == n {
			break
		}
		h = h.GetNext()
	}
	if h == nil {
		return false
	}
	temp := NewLinkNode(v)
	if h.GetNext() == nil {
		h.N = temp
	} else {
		temp.N = h.N
		h.N = temp
	}
	l.L++
	return true
}

// insert before the given node
//func (l *LinkedList) InsertBefore(n *Node, v interface{}) bool {
//	if n == nil {
//		return false
//	}
//	h := l.H
//	for h != nil {
//		if h == n {
//			break
//		}
//		h = h.GetNext()
//	}
//	if h == nil {
//		return false
//	}
//}

