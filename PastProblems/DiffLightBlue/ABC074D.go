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

	const inf = 1 << 60

	n := nextInt()
	a := make([][]int, n)
	for i := range a {
		a[i] = nextInts(n)
		a[i][i] = inf
	}

	sum := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 経路i-k-j
			// 直通の道がなくても事足りるか
			eq := false
			for k := 0; k < n; k++ {
				if a[i][j] > a[i][k]+a[k][j] {
					puts(-1)
					return
				}
				eq = eq || (a[i][j] == a[i][k]+a[k][j])
			}
			if !eq {
				sum += a[i][j]
			}
		}
	}

	puts(sum)
}
