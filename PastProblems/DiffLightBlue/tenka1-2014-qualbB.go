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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	s := next()
	t := make([]string, n)
	for i := range t {
		t[i] = next()
	}

	L := len(s)
	dp := make([]int, L+1)
	dp[0] = 1
	for i := 1; i <= L; i++ {
		for j := range t {
			lt := len(t[j])
			if i >= lt && s[i-lt:i] == t[j] {
				dp[i] += dp[i-lt]
				dp[i] %= mod
			}
		}
	}

	puts(dp[L])
}
