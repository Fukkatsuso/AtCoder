// bit全探索だとTLE
// 解説ACした

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isAllEven(a []int) bool {
	for i := range a {
		if a[i]%2 == 1 {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	n, p := nextInt(), nextInt()
	a := nextInts(n)

	ans := 0
	if isAllEven(a) {
		if p == 0 {
			ans = (1 << uint(n))
		} else {
			ans = 0
		}
	} else {
		ans = (1 << uint(n-1))
	}

	puts(ans)

	wt.Flush()
}

var (
	sc  = bufio.NewScanner(os.Stdin)
	rdr = bufio.NewReaderSize(os.Stdin, 1000000)
	wt  = bufio.NewWriter(os.Stdout)
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
