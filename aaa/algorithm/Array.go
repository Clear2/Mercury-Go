package main

import (
	"fmt"
	"sort"
)

// 原地修改数组
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var i = 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	for _, value := range nums {
		if m[value] {
			return value
		} else {
			m[value] = true
		}
	}
	return -1
}
func main() {
	//var num1 = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//var num2 = []int{1}
	//merge(num1, 1, num2, 1)
	//fmt.Println(num1, num2)
	//length := removeDuplicates(num1)
	fmt.Println(letterCasePermutation("a1b2"))
}

// Merge Sorted Array -> https://leetcode.com/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m <= 0 || nums2[n-1] >= nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n = n - 1
		} else {
			nums1[m+n-1] = nums1[m-1]
			m -= 1
		}
	}
}

// 数组不重复全排列 -> https://leetcode.com/problems/permutations-ii/
// https://leetcode.com/problems/permutations-ii/discuss/552551/Go-golang-solutions
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	dfs(&res, 0, nums)
	return res
}

func dfs(res *[][]int, index int, nums []int) {
	if index == len(nums)-1 {
		*res = append(*res, append([]int{}, nums...))
	}
	used := make(map[int]bool)
	for i := index; i < len(nums); i++ {
		if _, ok := used[nums[i]]; ok {
			continue
		}
		nums[i], nums[index] = nums[index], nums[i]
		dfs(res, index+1, nums)
		nums[i], nums[index] = nums[index], nums[i]
		used[nums[i]] = true
	}
}

//permutations => https://leetcode.com/problems/permutations/
func permute(nums []int) [][]int {
	var res [][]int
	backtrack(&res, []int{}, nums)
	return res
}

func backtrack(res *[][]int, ans []int, nums []int) {
	if len(nums) == 0 {
		*res = append(*res, ans)
	}
	for i, v := range nums {
		newNumes := append(append([]int{}, nums[:i]...), nums[i+1:]...)
		backtrack(res, append(ans, v), newNumes)
	}
}

// Letter Case Permutation => https://leetcode.com/problems/letter-case-permutation/
func letterCasePermutation(s string) []string {
	res := []string{}
	letterDFS(s, []byte{}, &res, 0)
	return res
}

func letterDFS(s string, b []byte, res *[]string, pos int) {
	if pos == len(s) {
		*res = append(*res, string(b))
		return
	}
	if s[pos] >= '0' && s[pos] <= '9' {
		letterDFS(s, append(b, s[pos]), res, pos+1)
		return
	}
	letterDFS(s, append(b, s[pos]), res, pos+1)
	letterDFS(s, append(b, switchCase(s[pos])), res, pos+1)

}
func switchCase(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return 'A' + (b - 'a')
	} else {
		return 'a' + (b - 'A')
	}
}
