/**
 * @Author: liubaoshuai3
 * @Date: 2020/7/7 20:27
 * @File: listr
 * @Description:
 */

package container

type Element struct {
	next, pre *Element
	list *List
	Value interface{}
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && &e.list.root != p {
		return p
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.pre; e.list != nil && &e.list.root != p {
		return p
	}
	return nil
}

type List struct {
	root Element
	len int
}

func (l *List) Init() *List {
	l.root.pre = &l.root
	l.root.next = &l.root
	l.len = 0
	return l
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.pre
}
