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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w, a, b := nextInt(), nextInt(), nextInt(), nextInt()

	s := make([][]int, h)
	for i := range s {
		s[i] = make([]int, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			val := 1
			if i < b && j < a {
				val = 0
			} else if i >= b && j >= a {
				val = 0
			}
			s[i][j] = val
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			putf("%d", s[i][j])
		}
		puts()
	}
}
