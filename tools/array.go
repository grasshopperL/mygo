/**
 * @Author: liubaoshuai3
 * @Date: 2020/4/2 21:26
 * @File: array
 * @Description:
 */

package tools

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

func (a *Array) IsIndexOutOfRange() bool {
	
}
