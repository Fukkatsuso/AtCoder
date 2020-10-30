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

	n := nextInt()
	a := nextInts(n)

	ans := 0
	l, r := 0, n-1
	for r-l > 1 {
		// 黒を置いたのと同じ側から白を置く必要がある
		right := a[r] + 1 + a[r-1]
		left := a[l] + 1 + a[l+1]
		if right < left {
			ans += a[r] + right
			r -= 2
			a[r] += right + 1
		} else {
			ans += a[l] + left
			l += 2
			a[l] += left + 1
		}
	}

	puts(ans)
}
