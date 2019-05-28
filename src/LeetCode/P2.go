package main

import (
	. "GLeetCode/src/Entity"
)
//
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	res := new(ListNode)
	point := res
	for l1 != nil || l2 != nil{
		v1 := 0
		v2 := 0
		if l1 != nil{
			v1 = l1.Val
		}
		if l2 != nil{
			v2 = l2.Val
		}
		point.Next = &ListNode{Val:(v1+v2 + carry)%10}
		carry = (v1+v2 + carry) / 10
		if l1 != nil{
			l1 = l1.Next
		}
		if l2 != nil{
			l2 = l2.Next
		}

		point = point.Next
	}
	if carry != 0{
		point.Next = &ListNode{Val:1}
	}
	return res.Next
}
