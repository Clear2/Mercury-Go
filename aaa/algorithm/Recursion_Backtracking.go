package main

import "fmt"

func main() {
	fmt.Println(fib2(2))
}

// 0 1 1 2 3 5
func fib1(n int) int {
	if n < 2 {
		return n
	}

	return fib1(n-1) + fib1(n-2)
}

func fib2(n int) int {
	cache := map[int]int{0: 0, 1: 1}
	return calculate(n, cache)
}

func calculate(N int, cache map[int]int) int {
	val, ok := cache[N]
	if ok {
		return val
	} else {
		cache[N] = calculate(N-1, cache) + calculate(N-2, cache)
		return cache[N]
	}
}
