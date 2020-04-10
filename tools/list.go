/**
 * @Author: liubaoshuai3
 * @Date: 2020/4/10 20:34
 * @File: list
 * @Description:
 */

package tools

// the element struct
type Element struct {
	next, prev *Element
	list *List
	Value interface{}
}

// the list struct
type List struct {
	root Element
	len int
}

// the next element
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// the prev element
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// init a new list
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}