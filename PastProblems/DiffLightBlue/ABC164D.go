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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := next()
	n := len(s)

	dp := make([]int, n+1)
	for i, tens := n-1, 1; i >= 0; i-- {
		dp[i] = (dp[i+1] + int(s[i]-'0')*tens) % 2019
		tens = (tens * 10) % 2019
	}

	mod := map[int]int{}
	for i := range dp {
		mod[dp[i]]++
	}

	ans := 0
	for _, v := range mod {
		ans += (v * (v - 1)) / 2
	}
	puts(ans)
}
