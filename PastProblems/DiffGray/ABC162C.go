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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	K := nextInt()

	g := make([][]int, K+1)
	for i := range g {
		g[i] = make([]int, K+1)
		for j := range g[i] {
			g[i][j] = gcd(i, j)
		}
	}

	ans := 0
	for i := 1; i <= K; i++ {
		for j := 1; j <= K; j++ {
			for k := 1; k <= K; k++ {
				ans += g[g[i][j]][k]
			}
		}
	}
	puts(ans)
}
