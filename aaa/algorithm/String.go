package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]

	for _, k := range strs {
		fmt.Println(k)

		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
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

func main() {
	var array = []string{"flower", "flow", "flight"}
	a := longestCommonPrefix(array)
	fmt.Println(a)
}
