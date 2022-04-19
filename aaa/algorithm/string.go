package main

import (
	"fmt"
)

// 14. Longest Common Prefix -> https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	mx := 0
	for {
		for i := 1; i < len(strs); i++ {
			if strs[i-1][mx] != strs[i][mx] {
				return string(strs[0][:mx])
			}
		}
		mx++
	}
	return string(strs[0][:mx])
}

// ICPrefixTwo 求 str1 与 str2 的最长公共前缀
func ICPrefixTwo(str1, str2 string) string {
	j := 0
	for ; j < len(str1) && j < len(str2); j++ {
		if str1[j] != str2[j] {
			break
		}
	}
	return str1[0:j]

}

// 48. Rotate Image => https://leetcode.com/problems/rotate-image/
func rotates(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 {
		return
	}
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[j][i], matrix[j][len(matrix)-1-i] = matrix[j][len(matrix)-1-i], matrix[j][i]
		}
	}
}
func main() {
	//var array = []string{"flower", "flow", "flight"}
	//a := longestCommonPrefix(array)
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotates(nums)
	fmt.Println(fn())
}

func fn() (int, string) {
	var a = 1
	var b = 2
	a, b = b, a

	return 1, "string"
}
