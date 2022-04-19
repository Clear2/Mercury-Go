package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2.https://leetcode.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c1 := l1
	c2 := l2
	n1, n2, ca := 0, 0, 0
	pre := &ListNode{}
	node := &ListNode{}

	for c1 != nil || c2 != nil {
		if c1.Val > 0 {
			n1 = c1.Val
		}
		if c2.Val > 0 {
			n2 = c2.Val
		}
		n := n1 + n2 + ca
		ca = n / 10 // 1
		pre = node
		node = &ListNode{Val: n % 10}
		node.Next = pre
		if c1 != nil {
			c1 = c1.Next
		}
		if c2 != nil {
			c2 = c2.Next
		}
	}

	if ca == 1 {
		pre = node
		node = &ListNode{Val: 1}
		node.Next = pre
	}
	return node
}
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil || l2 != nil {
		var sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
func PrintNode(head *ListNode) {
	for head != nil {
		fmt.Printf("->%d", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")

}
func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	PrintNode(addTwoNumbers1(l1, l2))
	//addTwoNumbers(l1, l2)

	//node := deleteDuplicates(node1)
	//PrintNode(node)
}
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode //指向当前节点的上一个节点，初始化为nil因为刚好头节点的上一个节点应该视为nil
	cur := head       //从头节点开始遍历

	//cur.Next, cur, pre = pre, cur.Next, cur //将当前节点指向上一个节点，同时将当前节点向前进，下一节点的上一个节点即当前节点
	pre = cur
	PrintNode(cur)
	fmt.Print("\n")
	PrintNode(cur.Next)
	fmt.Print("\n")
	PrintNode(pre)
	fmt.Print("\n")
	return pre
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	prev := prehead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 == nil {
		prev.Next = l2
	}
	if l2 == nil {
		prev.Next = l1
	}
	return prehead.Next
}
