package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 5, 6, 7, 9, 11}
	fmt.Println(maxPoint2(arr, 10))
}

func maxPoint2(arr []int, L int) int {
	left, right, max, N := 0, 0, 0, len(arr)

	for left < N {
		for right < N && arr[right]-arr[left] <= L {
			right = right + 1
		}
		left = left + 1
		if max < right-left {
			max = right - left
		}
	}
	return max
}
