package problem746

func minCostClimbingStairs(cost []int) int {
	costLen := len(cost)
	if costLen == 1 {
		return cost[0]
	}
	dp := make([]int, costLen+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < costLen+1; i++ {
		if cost[i-1]+dp[i-1] > dp[i-2]+cost[i-2] {
			dp[i] = dp[i-2] + cost[i-2]
		} else {
			dp[i] = cost[i-1] + dp[i-1]
		}
	}
	return dp[costLen]
}
