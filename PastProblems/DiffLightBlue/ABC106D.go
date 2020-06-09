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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m, q := nextInt(), nextInt(), nextInt()
	rail := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		rail[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		l, r := nextInt(), nextInt()
		rail[l][r]++
	}
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			rail[i][j] += rail[i][j-1]
		}
	}

	for i := 0; i < q; i++ {
		p, q := nextInt(), nextInt()
		ans := 0
		for j := p; j <= q; j++ {
			ans += rail[j][q]
		}
		puts(ans)
	}
}
