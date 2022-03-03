package main

import "fmt"

func main() {
	//nums := []int{1, 3, 5, 6}
	//a := search(nums, 4)
	fmt.Printf("byte-->%d\n", (3-2)>>1)
}

func search(nums []int, target int) int {
	return searchTarget(nums, target, 0, len(nums)-1)
}

func searchTarget(nums []int, target, L, R int) int {
	if L > R {
		return L
	}
	mid := L + (R-L)>>1
	fmt.Println(L, R, mid)
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return searchTarget(nums, target, L, mid-1)
	}
	return searchTarget(nums, target, mid+1, R)

}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
//func firstBadVersion(n int) int {
//	var L, R, Mid = 0, n, 0
//	for L < R {
//		Mid = L + (R-L)>>1
//		if isBadVersion(Mid) {
//			R = Mid
//		} else {
//			L = Mid + 1
//		}
//	}
//	return L
//}
