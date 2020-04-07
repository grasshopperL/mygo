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