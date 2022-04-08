package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 8, 10}
	//a := search(nums, 4)
	fmt.Printf("byte-->%d\n", searchRange(nums, 8))
}

func search(nums []int, target int) int {
	return searchTarget(nums, target, 0, len(nums)-1)
}

func searchTarget(nums []int, target, L, R int) int {
	if L >= R {
		return L
	}
	fmt.Println(L, R)
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

// 34--->https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	var L, R = 0, len(nums) - 1
	ret := []int{-1, -1}
	if R == -1 {
		return ret
	}

	for L < R {
		mid := L + (R-L)>>1
		if nums[mid] < target {
			L = mid + 1
		} else {
			R = mid
		}
	}
	if nums[L] != target {
		return ret
	} else {
		ret[0] = L
	}

	R = len(nums) - 1
	for L < R {

		mid := L + (R-L)>>1 + 1
		fmt.Println(L, R, mid)
		if nums[mid] > target {
			R = mid - 1
		} else {
			L = mid
		}
	}
	ret[1] = R

	return ret
}
func searchRangeTarget(num []int, target, L, R int) []int {
	mid := L + (R-L)>>1
	if L >= R {
		return []int{-1, -1}
	}
	if num[mid] == target {
		return []int{mid, mid + 1}
	}
	if num[mid] > target {
		return searchRangeTarget(num, target, L, mid-1)
	}
	return searchRangeTarget(num, target, mid+1, R)
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
