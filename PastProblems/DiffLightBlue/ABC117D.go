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

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func max(ones ...int) int {
	ret := ones[0]
	for _, v := range ones {
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

	n, k := nextInt(), nextInt()
	a := nextInts(n)

	// dp[i][j]: ibit目まで見て，確定値がk未満ならj=1，kと同じならj=0
	dp := make([][2]int, 50)
	for i := range dp {
		dp[i][0], dp[i][1] = -1, -1
	}
	dp[49][0] = 0
	for bit := 50 - 2; bit >= 0; bit-- {
		one := 0
		for i := range a {
			if (uint(a[i])>>uint(bit))&1 == 1 {
				one++
			}
		}

		digit := int(uint(1) << uint(bit))
		if dp[bit+1][1] >= 0 {
			// 0か1か，最大を取る方を使えば良い
			dp[bit][1] = max(dp[bit][1], dp[bit+1][1]+digit*max(one, n-one))
		}
		if dp[bit+1][0] >= 0 {
			if uint(k)&uint(digit) > 0 {
				// xのbit[bit]に0を使う
				dp[bit][1] = max(dp[bit][1], dp[bit+1][0]+digit*one)
				// 1を使う
				dp[bit][0] = max(dp[bit][0], dp[bit+1][0]+digit*(n-one))
			} else {
				// 0を使う
				dp[bit][0] = max(dp[bit][0], dp[bit+1][0]+digit*one)
			}
		}
	}

	puts(max(dp[0][0], dp[0][1]))
}
