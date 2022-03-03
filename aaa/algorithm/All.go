package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))
}

//1.Two Sum
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if p, ok := m[target-nums[i]]; ok {
			return []int{p, i}
		}
		m[nums[i]] = i
	}
	return nil
}
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var reverse = 0
	// x % 10 得到最后一位数字
	// 得到第二位数字 / 10 % 10
	for x > reverse {
		reverse = reverse*10 + x%10
		x = x / 10
	}
	fmt.Println(reverse)
	return x == reverse || x == reverse/10
}
