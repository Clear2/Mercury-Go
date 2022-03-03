package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Value == cur.Next.Value {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
func PrintNode(head *ListNode) {
	for head != nil {
		fmt.Printf("->%d", head.Value)
		head = head.Next
	}
	fmt.Print("->", head)
}
func main() {
	node1 := new(ListNode)
	node1.Value = 1
	node2 := new(ListNode)
	node2.Value = 1
	node3 := new(ListNode)
	node3.Value = 3
	node4 := new(ListNode)
	node4.Value = 5
	node5 := new(ListNode)
	node5.Value = 5

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node := deleteDuplicates(node1)
	PrintNode(node)
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
		if l1.Value <= l2.Value {
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
