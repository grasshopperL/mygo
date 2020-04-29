/**
 * @Author: liubaoshuai3
 * @Date: 2020/3/18 20:57
 * @File: singlelinked
 * @Description:
 */

package tools

import "fmt"

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
func (l *LinkedList) InsertBefore(n *Node, v interface{}) bool {
	if n == nil {
		return false
	}
	h := l.H
	for h != nil {
		if h.GetNext() == n {
			break
		}
		h = h.GetNext()
	}
	if h == nil {
		return false
	}
	temp := NewLinkNode(v)
	temp.N = n
	h.N = temp
	l.L++
	return true
}

// insert a node to head
func (l *LinkedList) InsertToHead(v interface{}) bool {
	n := NewLinkNode(v)
	if l.H == nil {
		l.H = n
	} else {
		n.N = l.H.N
		l.H.N = n
	}
	l.L++
	return true
}

// insert a node to tail
func (l *LinkedList) InsertToTail(v interface{}) bool {
	n := l.H
	for nil != n.N {
		n = n.GetNext()
	}
	return l.InsertAfter(n, v)
}

// find the node by index
func (l *LinkedList) FindNodeByIndex(i int) *Node {
	if i > l.L {
		return nil
	}
	n := l.H
	for i != 0 {
		i--
		n = n.GetNext()
	}
	return n
}

// delete the node by node
func (l *LinkedList) DeleteNode(n *Node) bool {
	h := l.H
	for h != nil {
		if h.GetNext() == n {
			break
		}
		h = h.GetNext()
	}
	if h == nil {
		return false
	}
	h.N = n.GetNext()
	l.L--
	return true
}

// print the list
func (l *LinkedList) Print() {
	cur := l.H.N
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetVal())
		cur = cur.GetNext()
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

type ListNode struct {
	Val int
	Next *ListNode
}
func merge(l1, l2 *ListNode) *ListNode {
	r := &ListNode{}
	c := r
	if l1 == nil {
		c.Next = l2
	}
	if l2 == nil {
		c.Next = l1
	}
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			c.Next = l2
			l2 = l2.Next
		} else {
			c.Next = l1
			l1 = l1.Next
		}
		c = c.Next
	}
	return c.Next
}


