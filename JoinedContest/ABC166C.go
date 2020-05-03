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

	n, m := nextInt(), nextInt()
	h := nextInts(n)

	top := make([]bool, n)
	for i := range top {
		top[i] = true
	}
	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		if h[a] <= h[b] {
			top[a] = false
		}
		if h[b] <= h[a] {
			top[b] = false
		}
	}

	ans := 0
	for i := range top {
		if top[i] {
			ans++
		}
	}
	puts(ans)
}
