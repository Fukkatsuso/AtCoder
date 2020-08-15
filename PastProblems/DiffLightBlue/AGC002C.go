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

	n, l := nextInt(), nextInt()
	a := nextInts(n)

	last := -1
	for i := 0; i < n-1; i++ {
		if a[i]+a[i+1] >= l {
			last = i
		}
	}
	if last < 0 {
		puts("Impossible")
		return
	}

	puts("Possible")
	for i := 0; i < last; i++ {
		puts(i + 1)
	}
	for i := n - 2; i > last; i-- {
		puts(i + 1)
	}
	puts(last + 1)
}
