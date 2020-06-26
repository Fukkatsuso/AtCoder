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

	n := nextInt()
	c := nextInts(n)

	ans := float64(0)
	for i := 0; i < n; i++ {
		// コインiの数字の約数の数
		s := 0
		for j := 0; j < n; j++ {
			if i == j || c[i]%c[j] != 0 {
				continue
			}
			s++
		}
		if s%2 == 1 {
			ans += 0.5
		} else {
			ans += (float64(s) + 2.0) / (2.0*float64(s) + 2.0)
		}
	}

	puts(ans)
}
