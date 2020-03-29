// 解説AC

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := nextInt(), nextInt()
	s := make([]int, n)
	for i := range s {
		a := nextInt()
		if i == 0 {
			s[i] = a
		} else {
			s[i] = s[i-1] + a
		}
	}

	ans := 0
	for i := 0; i < n-k+1; i++ {
		if i == 0 {
			ans += s[i+k-1]
		} else {
			ans += s[i+k-1] - s[i-1]
		}
	}

	puts(ans)
}

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
