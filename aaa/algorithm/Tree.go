package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
	fmt.Println(mergeTrees(&node1, &nodeb))

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

/*
	1
   /  \
  2    3
 / \   / \
4   5 6   7
*/
// 二叉树的递归遍历
// 1.中序遍历-> https://leetcode.com/problems/binary-tree-inorder-traversal/
// 2.前序遍历-> https://leetcode.com/problems/binary-tree-preorder-traversal/
// 3.后序遍历-> https://leetcode.com/problems/binary-tree-postorder-traversal/
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

// 相同的树-> https://leetcode.com/problems/same-tree/
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

// 合并两个二叉树-> https://leetcode.com/problems/merge-two-binary-trees/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return root1
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	newTree := &TreeNode{Val: root1.Val + root2.Val}
	newTree.Left = mergeTrees(root1.Left, root2.Left)
	newTree.Right = mergeTrees(root1.Right, root2.Right)
	return newTree

}

//1 +2 =  3
// 3 + 1 = 4
// 5 + 0 = 5
