// 1問REでドハマり
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

func reachable(f []int, from, to int) bool {
	n := len(f)
	dp := make([][16005]bool, n+1)
	dp[0][from+8000] = true
	for i := 0; i < n; i++ {
		for j := 0; j < 16005; j++ {
			if dp[i][j] {
				dp[i+1][j-f[i]], dp[i+1][j+f[i]] = true, true
			}
		}
	}
	return dp[n][to+8000]
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s, x, y := next(), nextInt(), nextInt()

	ff := make([][]int, 2)
	ff[0], ff[1] = []int{0}, []int{}
	for i, t := 0, 0; i < len(s); i++ {
		if s[i] == 'T' {
			t++
			ff[t%2] = append(ff[t%2], 0)
		} else {
			ff[t%2][len(ff[t%2])-1]++
		}
	}

	if !reachable(ff[0][1:], ff[0][0], x) {
		puts("No")
		return
	}
	if !reachable(ff[1], 0, y) {
		puts("No")
		return
	}
	puts("Yes")
}
