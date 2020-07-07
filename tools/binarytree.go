/**
 * @Author: liubaoshuai3
 * @Date: 2020/3/13 20:59
 * @File: binarytree
 * @Description:
 */

package tools

import (
	"fmt"
	"strconv"
)

// tree's node
type TreeNode struct {
	Val interface{}
	L *TreeNode
	R *TreeNode
}

// binary tree
type BinTree struct {
	H *TreeNode
}

// make a new node
func NewNode(v interface{}) *TreeNode {
	return &TreeNode{Val: v}
}

// return node as string
func (t *TreeNode) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", t.Val, t.L, t.R)
}

//make a new binary tree
func NewBinaryTree(v interface{}) *BinTree {
	return &BinTree{NewNode(v)}
}

// pre order
func PreOrder(t *BinTree) *ArrayStack {
	r := t.H
	if r == nil {
		return nil
	}
	s := NewArrayStack()
	res := NewArrayStack()
	_ = s.Push(r)
	for !s.IsEmpty() {
		res.Push(s.Pop())
		if r.R != nil {
			s.Push(r.L)
		}
		if r.R != nil {
			s.Push(r.R)
		}
	}
	return res
}

// in order
func InOrder(t *BinTree) *ArrayStack {
	r := t.H
	res := NewArrayStack()
	s := NewArrayStack()
	if r == nil {
		return nil
	}
	for r != nil || !s.IsEmpty() {
		if r != nil {
			s.Push(r)
			r = r.L
		} else {
			e := s.Pop().(*TreeNode)
			res.Push(e.Val)
			r = e.R
		}
	}
	return res
}

// pre order
func PreOrderTwo(t *BinTree) *ArrayStack {
	r := t.H
	res := NewArrayStack()
	s := NewArrayStack()
	if r == nil {
		return nil
	}
	for r != nil || !s.IsEmpty() {
		if r != nil {
			res.Push(r.Val)
			s.Push(r)
			r = r.L
		} else {
			r = s.Pop().(*TreeNode).R
		}
	}
	return res
}

// post order ok, if you say why two stack, you can down a little
func PostOrder(t *BinTree) *ArrayStack {
	r := t.H
	if r == nil {
		return nil
	}
	res := NewArrayStack()
	s1 := NewArrayStack()
	s2 := NewArrayStack()
	s1.Push(r)
	for !s1.IsEmpty() {
		e := s1.Pop().(*TreeNode)
		s2.Push(e)
		if e.L != nil {
			s1.Push(e.L)
		}
		if e.R != nil {
			s1.Push(e.R)
		}
	}
	for !s2.IsEmpty() {
		res.Push(s2.Pop().(*TreeNode).Val)
	}
	return res
}

// post order two
func PostOrderTwo(t *BinTree) *ArrayStack {
	r := t.H
	res := NewArrayStack()
	s := NewArrayStack()
	s.Push(r)
	var pre *TreeNode
	if !s.IsEmpty() {
		r = s.Top().(*TreeNode)
		if (r.L != nil && r.R != nil) || (pre != nil && (pre == r.R || pre == r.L)) {
			res.Push(r.Val)
			s.Pop()
			pre = r
		} else {
			if r.R != nil {
				s.Push(r.R)
			}
			if r.L != nil {
				s.Push(r.L)
			}
		}
	}
	return res
}

func decodeString(s string) string {
	stack, res := []string{}, ""
	for _, str := range s {
		if len(stack) == 0 || (len(stack) > 0 && str != ']') {
			stack = append(stack, string(str))
		} else {
			tmp := ""
			for stack[len(stack)-1] != "[" {
				tmp = stack[len(stack)-1] + tmp
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			index, repeat := 0, ""
			for index = len(stack) - 1; index >= 0; index-- {
				if stack[index] >= "0" && stack[index] <= "9" {
					repeat = stack[index] + repeat
				} else {
					break
				}
			}
			nums, _ := strconv.Atoi(repeat)
			copyTmp := tmp
			for i := 0; i < nums-1; i++ {
				tmp += copyTmp
			}
			res += tmp
			stack = []string{}
		}
	}
	return res
}