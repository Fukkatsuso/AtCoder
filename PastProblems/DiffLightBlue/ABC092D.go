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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func NewSlice(h, w int) [][]byte {
	s := make([][]byte, h)
	for i := range s {
		s[i] = make([]byte, w)
	}
	return s
}

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, b := nextInt(), nextInt()
	a--
	b--

	k := 50
	h, w := 2*k, 2*k
	ans := NewSlice(h, w)
	for i := range ans {
		for j := range ans[i] {
			if i < k-1 {
				ans[i][j] = '#'
			} else {
				ans[i][j] = '.'
			}
		}
	}

	for i := 0; i < k-1 && a > 0; i += 2 {
		for j := 0; j < w && a > 0; j += 2 {
			ans[i][j] = '.'
			a--
		}
	}

	for i := k; i < h && b > 0; i += 2 {
		for j := 0; j < w && b > 0; j += 2 {
			ans[i][j] = '#'
			b--
		}
	}

	puts(h, w)
	for i := range ans {
		for j := range ans[i] {
			putf("%c", ans[i][j])
		}
		puts()
	}
}
