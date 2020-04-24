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

	n, x := nextInt(), nextInt()

	// a[i]: レベルiのバーガーの長さ
	// p[i]: レベルiのバーガーのパティ
	a, p := make([]int, n+1), make([]int, n+1)
	a[0], p[0] = 1, 1
	for i := 1; i <= n; i++ {
		a[i] = a[i-1]*2 + 3
		p[i] = p[i-1]*2 + 1
	}

	var f func(level, x int) int
	f = func(level, x int) int {
		if level == 0 {
			if x == 1 {
				return 1
			}
			return 0
		}
		if x == 1 {
			return 0
		} else if x <= a[level-1]+1 {
			return f(level-1, x-1)
		} else if x == a[level-1]+2 {
			return p[level-1] + 1
		} else if x <= 2*a[level-1]+2 {
			return p[level-1] + 1 + f(level-1, x-a[level-1]-2)
		}
		return 1 + 2*p[level-1]
	}

	puts(f(n, x))
}
