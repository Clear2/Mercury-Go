package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//s1 := []int{2, 3, 4}
	//x := sortedSquares(s1)
	//x := twoSum(s1, 6)
	s := "OllEH DLROW"
	fmt.Println(reverseWords(s))
}

func sortedSquares(nums []int) []int {
	var L, R = 0, len(nums) - 1
	result := make([]int, len(nums), len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if math.Abs(float64(nums[L])) > math.Abs(float64(nums[R])) {
			result[i] = nums[L] * nums[L]
			L++
		} else {
			result[i] = nums[R] * nums[R]
			R--
		}
	}
	return result
}

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, L, R int) {
	for L < R {
		nums[L], nums[R] = nums[R], nums[L]
		L++
		R--
	}
}

func moveZeroes(nums []int) {
	var L = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[L], nums[i] = nums[i], nums[L]
			L++
		}
	}
}

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		j--
		i++
	}
	//
	// 只有最后一个元素了，返回
	//if len(s) <= 1 {
	//	return
	//}
	// 第一个和最后一个元素交换
	//s[i], s[j] = s[j], s[i]
	// 递归调用
	//reverseString(s[1:j])
}

func reverseWords(s string) string {
	// 根据空格切割字符串
	word := strings.Split(s, " ")
	for k, v := range word {
		rev := []byte(v)
		// 反转字符串
		reverseString(rev)
		// 修改字符串数组
		word[k] = string(rev)
	}
	return strings.Join(word, " ")
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head} // dummy head
	slow, fast := dummy, head
	for fast != nil {
		fast = fast.Next
		if n == 0 {
			slow = slow.Next
		} else {
			n--
		}
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
