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

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) insert(e, at *Element) *Element {
	e.pre = at
	e.next = at.next
	e.pre.next = e
	e.next.pre = e
	e.list = l
	l.len++
	return e
}

func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{
		next:  nil,
		pre:   nil,
		list:  nil,
		Value: v,
	}, at)
}

func (l *List) remove(e *Element) *Element {
	e.pre.next = e.next
	e.next.pre = e.pre
	l.len--
	e.list = nil
	return e
}

// move e to after at
//func (l *List) move(e, at *Element) *Element {
//	if e == at {
//		return e
//	}
//	e.pre.next = at
//
//}