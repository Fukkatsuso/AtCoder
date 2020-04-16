// 解説AC

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

func cmb(n int) int {
	return (n * (n - 1)) / 2
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a := nextInts(n)

	s := make([]int, n)
	s[0] = a[0]
	for i := 1; i < n; i++ {
		s[i] = s[i-1] + a[i]
	}
	sort.Sort(sort.IntSlice(s))

	ans := 0
	for l := 0; l < n; l++ {
		r := l
		for r+1 < n && s[r+1] == s[l] {
			r++
		}
		ans += cmb(r - l + 1)
		if s[l] == 0 {
			ans += r - l + 1
		}
		l = r
	}

	puts(ans)
}
