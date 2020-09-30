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
	mod            = 1000000007
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, _ := nextInt(), nextInt()
	f := make([]int, n)
	for i := range f {
		f[i] = nextInt() - 1
	}

	// dp[i]: i番目までのサプリを摂取する組み合わせ, 1-indexed
	// sum[i]: sum(dp[0:i])
	dp, sum := make([]int, n+1), make([]int, n+1)
	dp[0] = 1
	l := 0
	// cnt[k]: fの区間[l,i)に数字kが出現している回数
	cnt := map[int]int{}
	for i := 1; i <= n; i++ {
		// しゃくとり法
		cnt[f[i-1]]++
		for l < i && cnt[f[i-1]] > 1 {
			cnt[f[l]]--
			l++
		}
		// 累積和
		sum[i] = (sum[i-1] + dp[i-1]) % mod
		// サプリ(l~i-1:i]を1日で一気に摂取
		dp[i] = (sum[i] - sum[l] + mod) % mod
	}

	puts(dp[n])
}
