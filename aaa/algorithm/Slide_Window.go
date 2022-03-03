package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLongestSubstring(s string) int {
	n, ans := len(s), 0
	m := make(map[byte]int)
	for j, i := 0, 0; j < n; j++ {
		if m[s[j]] != 0 {
			i = maxInt(m[s[j]], i)
		}
		ans = maxInt(ans, j-i+1)
		m[s[j]] = j + 1
	}
	return ans
}

func checkInclusion(s1 string, s2 string) bool {
	return false
}
