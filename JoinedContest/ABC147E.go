// WA

package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(nums []int) int {
	ret := nums[0]
	index := 0
	for i, v := range nums {
		if v < ret {
			ret = v
			index = i
		}
	}
	return index
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	a := make([][]int, h)
	b := make([][]int, h)
	for i := range a {
		a[i] = make([]int, w)
		for j := range a[i] {
			fmt.Scan(&a[i][j])
		}
	}
	for i := range b {
		b[i] = make([]int, w)
		for j := range b[i] {
			fmt.Scan(&b[i][j])
		}
	}

	dp := make([][]int, h)
	for i := range dp {
		dp[i] = make([]int, w)
	}
	dp[0][0] = a[0][0] - b[0][0]
	for i := 1; i < w; i++ {
		sum1 := dp[0][i-1] + a[0][i] - b[0][i]
		sum2 := dp[0][i-1] + b[0][i] - a[0][i]
		if abs(sum1) < abs(sum2) {
			dp[0][i] = sum1
		} else {
			dp[0][i] = sum2
		}
	}
	for j := 1; j < h; j++ {
		sum1 := dp[j-1][0] + a[j][0] - b[j][0]
		sum2 := dp[j-1][0] + b[j][0] - a[j][0]
		if abs(sum1) < abs(sum2) {
			dp[j][0] = sum1
		} else {
			dp[j][0] = sum2
		}
	}

	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			sum1 := dp[i-1][j] + a[i][j] - b[i][j]
			sum2 := dp[i-1][j] + b[i][j] - a[i][j]
			sum3 := dp[i][j-1] + a[i][j] - b[i][j]
			sum4 := dp[i][j-1] + b[i][j] - a[i][j]
			sums := []int{sum1, sum2, sum3, sum4}
			abss := []int{abs(sum1), abs(sum2), abs(sum3), abs(sum4)}
			dp[i][j] = sums[min(abss)]
		}
	}
	fmt.Println(dp[h-1][w-1])
}
