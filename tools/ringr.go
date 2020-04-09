/**
 * @Author: liubaoshuai3
 * @Date: 2020/4/6 20:34
 * @File: ringr
 * @Description:
 */

package tools

// the ring struct called ringer
type Ringr struct {
	Value interface{}
	pre, next *Ringr
}

// init return self?
func (r *Ringr) init() *Ringr {
	r.pre = r;
	r.next = r
	return r
}

// get the next value
func (r *Ringr) Next() *Ringr {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// get the pre value
func (r *Ringr) Prev() *Ringr {
	if r.pre == nil {
		return r.init()
	}
	return r.pre
}

// move the value by given steps
func (r *Ringr) Move(n int) *Ringr {
	if r.next == nil {
		return r.init()
	}
	switch  {
	case n < 0:
		for ; n < 0; n++ {
			r = r.pre
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// new a ringr
func New(n int) *Ringr {
	if n <= 0 {
		return nil
	}
	r := new(Ringr)
	p := r
	for i := 0; i < n; i++ {
		p.next = &Ringr{pre:p}
		p = p.next
	}
	p.next = r
	r.pre = p
	return r
}

// link the ring
func (r *Ringr) Link(s *Ringr) *Ringr {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.pre = r
		n.pre = p
		p.next = n
	}
	return n
}

// remove some ringers from ring
// a little meng
func (r *Ringr) Unlink(n int) *Ringr {
	if n <= 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

// the length of the ring
func (r *Ringr) Len() int {
	n := 0
	if r != nil {
		n := 1
		for p := r.Next(); p != r; p = p.next {
			n++
		}
	}
	return n
}
