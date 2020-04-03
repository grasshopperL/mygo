/**
 * @Author: liubaoshuai3
 * @Date: 2020/4/2 21:26
 * @File: array
 * @Description:
 */

package tools

import "errors"

// just for fun
type Array struct {
	d []int
	l int
} 

// init an array
func InitArray(l int) *Array {
	if l == 0 {
		return nil
	}
	return &Array{
		d: make([]int, l),
		l: l,
	}
}

// the length of array
func (a *Array) Len() int {
	return a.l
}

// jude index out of range or not
func (a *Array) IsIndexOutOfRange(i int) bool {
	if i > cap(a.d) {
		return true
	}
	return false
}

// find value by
func (a *Array) FindByIndex(i int) (error, int) {
	if a.IsIndexOutOfRange(i) {
		return errors.New("out of range"), 0
	}
	return nil, a.d[i]
}

// insert value by index
func (a *Array) InsertByIndex(i int, v int) bool {
	
	return true
}
