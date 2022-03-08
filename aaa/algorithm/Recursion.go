package main

import "fmt"

func main() {
	fmt.Println(climbStairs1(3))
}

// Climbing Stairs
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	count := climbStairs(n-1) + climbStairs(n-2)
	return count
}

// Memorization
func climbStairs1(n int) int {
	memo := map[int]int{1: 1, 2: 2}
	return memoClimb(n, memo)
}
func memoClimb(n int, memo map[int]int) int {
	if v, ok := memo[n]; ok {
		return v
	}
	memo[n] = memoClimb(n-1, memo) + memoClimb(n-2, memo)
	return memo[n]
}
