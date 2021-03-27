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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func primeFactor(n int) map[int]int {
	m := map[int]int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			m[i]++
			n /= i
		}
	}
	if n != 1 {
		m[n] = 1
	}
	return m
}

func factorCount(n int) int {
	pf := primeFactor(n)
	res := 1
	for _, v := range pf {
		res *= (v + 1)
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	k := getInt()

	// dp[i]: x*y <= iとなるx,yの組み合わせ
	dp := make([]int, k+1)
	for i := 1; i <= k; i++ {
		// x*y == i
		dp[i] = factorCount(i)
		// x*y < iの分を加算
		dp[i] += dp[i-1]
	}

	ans := 0
	for i := 1; i <= k; i++ {
		ans += dp[k/i]
	}

	puts(ans)
}
