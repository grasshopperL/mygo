/**
 * @Author: liubaoshuai3
 * @Date: 2021/5/7 18:36
 * @File: 2
 * @Description: https://leetcode-cn.com/problems/add-two-numbers/
 */

package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := l1
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		carry = sum / 10
	}
	return res.Next
}

// if don't malloc other space