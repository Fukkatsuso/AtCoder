package main

// 文字列sとtの最長共通部分列
// O(|s||t|)
func LCS(s, t string) int {
	n, m := len(s), len(t)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			x := dp[i-1][j-1]
			if s[i-1] == t[j-1] {
				x++
			}
			dp[i][j] = max(dp[i-1][j], dp[i][j-1], x)
		}
	}
	return dp[n][m]
}
