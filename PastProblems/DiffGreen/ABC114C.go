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

func is753(n int) bool {
	m := map[int]bool{}
	for n > 0 {
		m[n%10] = true
		n /= 10
	}
	return m[3] && m[5] && m[7]
}

func dfs(max, n int) int {
	if n > max {
		return 0
	}
	if is753(n) {
		return 1 + dfs(max, n*10+3) + dfs(max, n*10+5) + dfs(max, n*10+7)
	}
	return dfs(max, n*10+3) + dfs(max, n*10+5) + dfs(max, n*10+7)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	puts(dfs(n, 0))
}
