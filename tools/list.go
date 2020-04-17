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

// new a list
func NewList() *List {
	return new(List).Init()
}

// the length of the list
func (l *List) Len() int {
	return l.len
}

// the front node of l
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// the prev node of l
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazy init
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert a element after at
func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	at.next = e
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// insert value
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{
		Value: v,
	}, at)
}

// remove a value
// why = nil can avoid memory leaks
func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
	return e
}

// move node to after to given node
func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	e.next = e.prev.next
	e.prev = e.next.prev
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	l.len++
	return e
}

// remove a node from list
func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

// why lazy init?
func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v ,&l.root)
}

// why lazy init
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// insert element before mark
func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

// insert element after mark
func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.next)
}
