// 解説AC
// 考察はおしかった

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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func indivterms(L, x int) int {
	for x > 0 {
		puts(L, L+1)
		L += 3
		x--
	}
	return L
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()

	if m < 0 || n == m {
		puts(-1)
		return
	}

	// 交わる区間がない
	if m == 0 {
		indivterms(1, n)
		return
	}

	if n >= 2 && m == n-1 {
		puts(-1)
		return
	}

	R := indivterms(2, n-1)
	puts(R-3*(m+1)-1, R)
}
