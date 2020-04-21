// 解説の方がスマートな解答
// 一応AC

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	a := nextInts(n)
	sort.Sort(sort.IntSlice(a))

	ans := n
	double := 0
	for l := 0; l < n; {
		r := l
		for r+1 < n && a[r+1] == a[l] {
			r++
		}
		cnt := r - l + 1
		if cnt%2 == 0 {
			double++
		}
		if cnt > 2 {
			ans -= ((cnt - 1) / 2) * 2
		}
		l = r + 1
	}

	ans -= ((double + 1) / 2) * 2

	puts(ans)
}
