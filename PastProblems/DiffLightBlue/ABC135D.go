package main

import (
	"bufio"
	"fmt"
	"os"
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func reverseString(s string) string {
	n := len(s)
	bytes := []byte(s)
	for i := 0; i < n/2; i++ {
		bytes[i], bytes[n-1-i] = bytes[n-1-i], bytes[i]
	}
	return string(bytes)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const mod = 1000000007

	s := reverseString(next())
	n := len(s)

	// dp[i][j]: 下からi桁目まで見て13で割ったあまりがjとなる数の個数
	dp := make([][13]int, n+1)
	dp[0][0] = 1
	pow := 1
	for i := 0; i < n; i, pow = i+1, (pow*10)%13 {
		if s[i] != '?' {
			x := int(s[i] - '0')
			for j := 0; j < 13; j++ {
				dp[i+1][(j+x*pow)%13] = dp[i][j]
			}
			continue
		}
		for x := 0; x < 10; x++ {
			for j := 0; j < 13; j++ {
				dp[i+1][(j+x*pow)%13] += dp[i][j]
				dp[i+1][(j+x*pow)%13] %= mod
			}
		}
	}

	puts(dp[n][5])
}
