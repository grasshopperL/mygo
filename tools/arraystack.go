/**
 * @Author: liubaoshuai3
 * @Date: 2020/3/13 20:58
 * @File: arraystack
 * @Description:
 */

package tools

import "fmt"

// stack with array
type ArrayStack struct {
	top int
	data []interface{}
}

// new a stack with array
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		top:  -1,
		data: make([]interface{}, 0, 32),
	}
}

// stack is empty or not
func (a *ArrayStack) IsEmpty() bool {
	//if len(a.data) == 0 {
	//	return true
	//}
	//return false
	if a.top < 0 {
		return true
	}
	return false
}

//push a value to stack
func (a *ArrayStack) Push(v interface{}) int {
	if a.top < 0 {
		a.top = 0
	} else {
		a.top += 1
	}
	if a.top < len(a.data) {
		a.data[a.top] = v
	} else {
		a.data = append(a.data, v)
	}
	return a.top
}

//Pop
func (a *ArrayStack) Pop() interface{} {
	if a.top < 0 {
		return nil
	}
	e := a.data[a.top]
	if a.top == 0 {
		a.top = -1
	} else {
		a.top -= 1
	}
	return e
}

//top
func (a *ArrayStack) Top() interface{} {
	if a.top < 0 {
		return nil
	}
	t := a.data[a.top]
	return t
}

// Flush
func (a *ArrayStack) FlushStack() {
	a.top = -1
}
//print
func (a *ArrayStack) Print() []interface{} {
	if a.top < 0 {
		fmt.Printf("The stack is empty")
		return nil
	} else {
		var res []interface{}
		for i := a.top; i >= 0; i-- {
			res = append(res, a.data[i])
		}
		return res
	}
}