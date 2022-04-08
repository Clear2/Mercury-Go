package main

import "fmt"

func main() {
	fmt.Println(way2(2, 4, 4, 4))
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

//
