package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 100, 20, 50}
	//fmt.Println(way2(2, 4, 4, 4))
	fmt.Println(win1(arr))
}
func way1(start, K, aim, N int) int {
	// 机器人当前位置start
	// 还有K步要走
	// 最终的目标是aim
	// 有哪些位置？ 1~N
	return process1(start, K, aim, N)
}
func process1(cur, rest, aim, N int) int {
	if rest == 0 {
		if cur == aim {
			return 1
		}
		return 0
	}
	if cur == 1 {
		return process1(cur+1, rest-1, aim, N)
	}
	if cur == N {
		return process1(cur-1, rest-1, aim, N)
	}
	// 中间位置上!
	return process1(cur-1, rest-1, aim, N) + process1(cur+1, rest-1, aim, N)
}

func way2(start, K, aim, N int) int {
	// 机器人当前位置start
	// 还有K步要走
	// 最终的目标是aim
	// 有哪些位置？ 1~N
	//var dp int[N+1][K+1] = {}
	// go初始化二维slice方法
	// process(cur, rest)之前没算过! dp[cur][rest] == -1
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = -1
		}
	}

	return process2(start, K, aim, N)
}
func process2(cur, rest, aim, N int) int {
	if rest == 0 {
		if cur == aim {
			return 1
		}
		return 0
	}
	if cur == 1 {
		return process1(cur+1, rest-1, aim, N)
	}
	if cur == N {
		return process1(cur-1, rest-1, aim, N)
	}
	// 中间位置上!
	return process1(cur-1, rest-1, aim, N) + process1(cur+1, rest-1, aim, N)
}

/**
给定一个整型数组arr,代表数值不同的纸牌排成一条线
玩家A和玩家B依次拿走每张纸牌
规定玩家A先拿，玩家B后拿
但是每个玩家每次只能拿走最左或最右的纸牌
玩家A和玩家B都绝顶聪明
请返回最后获胜者的分数
*/
func win(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	f := first(arr, 0, len(arr)-1)
	l := last(arr, 0, len(arr)-1)
	if f > l {
		return f
	}
	return l
}

// 先手函数-> arr[L, R]先手获得最好的分数返回
func first(arr []int, L, R int) int {
	if L == R {
		return arr[L]
	}
	p1 := arr[L] + last(arr, L+1, R)
	p2 := arr[R] + last(arr, L, R-1)
	if p1 > p2 {
		return p1
	}
	return p2
}

func last(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	p1 := first(arr, L+1, R) //
	p2 := first(arr, L, R-1)
	if p1 < p2 {
		return p1
	}
	return p2
}

func win1(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	var N = len(arr)
	fmap := make([][]int, N)
	gmap := make([][]int, N)
	for i := 0; i < N; i++ {
		fmap[i] = make([]int, N)
		gmap[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmap[i][j] = -1
			gmap[i][j] = -1
		}
	}

	f := first1(arr, 0, len(arr)-1, fmap, gmap)
	l := last1(arr, 0, len(arr)-1, fmap, gmap)
	if f > l {
		return f
	}
	return l
}

// 	arr := []int{ 50, 100, 20, 10}

// 先手函数-> arr[L, R]先手获得最好的分数返回
func first1(arr []int, L, R int, fmap, gmap [][]int) int {
	if fmap[L][R] != -1 {
		return fmap[L][R]
	}
	ans := 0
	if L == R {
		ans = arr[L]
	}
	p1 := arr[L] + last1(arr, L+1, R, fmap, gmap)
	p2 := arr[R] + last1(arr, L, R-1, fmap, gmap)
	if p1 > p2 {
		ans = p1
	}
	ans = p2
	fmap[L][R] = ans
	return ans
}

func last1(arr []int, L, R int, fmap, gmap [][]int) int {
	fmt.Println("L", L, R, gmap[L][R])

	if gmap[L][R] != -1 {
		return gmap[L][R]
	}
	ans := 0
	if L != R {
		p1 := first1(arr, L+1, R, fmap, gmap) //
		p2 := first1(arr, L, R-1, fmap, gmap)
		if p1 < p2 {
			ans = p1
		}
		ans = p2
	}
	gmap[L][R] = ans
	return ans
}
