package main

/*
You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor.



Example 1:

Input: cost = [10,15,20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.
Example 2:

Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Explanation: You will start at index 0.
- Pay 1 and climb two steps to reach index 2.
- Pay 1 and climb two steps to reach index 4.
- Pay 1 and climb two steps to reach index 6.
- Pay 1 and climb one step to reach index 7.
- Pay 1 and climb two steps to reach index 9.
- Pay 1 and climb one step to reach the top.
The total cost is 6.


Constraints:

2 <= cost.length <= 1000
0 <= cost[i] <= 999
*/

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)

	// 最初の2段は0からでも1からでも開始できる
	dp[0] = 0
	dp[1] = 0

	// 各段でのコストの最小値を計算
	for i := 2; i <= n; i++ {
		// i番目の段に到達するには、i-1番目かi-2番目の段から来る必要がある
		// それぞれの段でのコストと、そこまでの最小コストを足した値の小さい方を選ぶ
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[n]
}
