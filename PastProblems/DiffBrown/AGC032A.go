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

	n := nextInt()
	b := nextInts(n)

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		for j := len(b) - 1; j >= 0; j-- {
			if b[j] == j+1 {
				ans[n-1-i] = j + 1
				b = delete(b, j)
				break
			}
		}
		if ans[n-1-i] == 0 {
			break
		}
	}

	if ans[0] == 0 {
		puts(-1)
	} else {
		for i := range ans {
			puts(ans[i])
		}
	}
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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func delete(a []int, i int) []int {
	a = append(a[:i], a[i+1:]...)
	res := make([]int, len(a))
	copy(res, a)
	return res
}
