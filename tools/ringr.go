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