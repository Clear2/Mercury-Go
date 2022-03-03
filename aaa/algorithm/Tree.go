package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraverSal(root *TreeNode) []int {
	var result []int
	inorderRecursive(root, &result)
	return result
}

func inorderRecursive(root *TreeNode, output *[]int) {
	if root != nil {
		inorderRecursive(root.Left, output)
		*output = append(*output, root.Val)
		inorderRecursive(root.Right, output)
	}
}

func main() {
	node4 := TreeNode{Val: 4}
	node5 := TreeNode{Val: 5}
	node6 := TreeNode{Val: 6}
	node7 := TreeNode{Val: 7}
	node2 := TreeNode{Val: 2, Left: &node4, Right: &node5}
	node3 := TreeNode{Val: 3, Left: &node6, Right: &node7}
	node1 := TreeNode{Val: 1, Left: &node2, Right: &node3}

	nodeb := TreeNode{Val: 1, Left: &node2, Right: &node4}
	//inorder(&node1)
	//result := inorderTraverSal(&node1)
	//fmt.Println(result)
	fmt.Println(isSameTree(&node1, &nodeb))
	arr := []int{0, 1, 2}
	m := map[int]int{2: 2, 3: 3}
	Change(arr, m)
	fmt.Println(arr, m)
}

func Change(arr []int, m map[int]int) {
	arr = append(arr, 4)
	m[0] = 1
}

/*
	1
   /  \
  2    3
 / \   / \
4   5 6   7
*/
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println("a---->>", root.Val) // 1 2 4
	inorder(root.Left)
	fmt.Println("b---->>", root.Val) // 4
	inorder(root.Right)
	fmt.Println("c---->>", root.Val) // 4
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
