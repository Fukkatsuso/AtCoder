// 解説AC

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	r, c, k := nextInt(), nextInt(), nextInt()
	item := make([][]int, r)
	for i := 0; i < r; i++ {
		item[i] = make([]int, c)
	}
	for i := 0; i < k; i++ {
		ri, ci, v := nextInt()-1, nextInt()-1, nextInt()
		item[ri][ci] = v
	}

	dp := make([][][]int, r)
	for i := range dp {
		dp[i] = make([][]int, c)
		for j := range dp[i] {
			dp[i][j] = make([]int, 4)
		}
	}
	dp[0][0][1] = item[0][0]
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			for k := 0; k <= 3; k++ {
				if i-1 >= 0 {
					dp[i][j][0] = max(dp[i][j][0], dp[i-1][j][k])
					dp[i][j][1] = max(dp[i][j][1], dp[i-1][j][k]+item[i][j])
				}
				if j-1 >= 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[i][j-1][k])
					if k > 0 {
						dp[i][j][k] = max(dp[i][j][k], dp[i][j-1][k-1]+item[i][j])
					}
				}
			}
		}
	}

	ans := max(dp[r-1][c-1]...)
	puts(ans)
}
