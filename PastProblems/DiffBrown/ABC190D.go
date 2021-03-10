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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()

	ans := 0

	// m項の数列
	for m := 1; m*m <= 2*n; m++ {
		if 2*n%m != 0 {
			continue
		}
		// 初項a, aa=2*a
		aa := 2*n/m - m + 1
		if aa < 0 || aa%2 != 0 {
			continue
		}
		ans++
		if aa/2 > 0 {
			ans++
		}
	}

	puts(ans)
}
