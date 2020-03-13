/**
 * @Author: liubaoshuai3
 * @Date: 2020/3/13 20:59
 * @File: binarytree
 * @Description:
 */

package tools

import "fmt"

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
func preOrder(r *TreeNode) []int64 {
	if r == nil {
		return nil
	}
	if r.L == nil && r.R == nil {
		return []int64{r.Val.(int64)}
	}
	var res []int64
	var stack []*TreeNode
	stack = append(stack, r)
	for len(stack) != 0 {
		e := stack[len(stack) - 1]
		res = append(res, e.Val.(int64))
		stack = stack[:len(stack) - 1]
		if e.R != nil {
			stack = append(stack, e.R)
		}
		if r.L != nil {
			stack = append(stack, e.L)
		}
	}
	return res
}
