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
	v := make([]int, n)
	x := 0
	for i := 0; i < n; i++ {
		a, b := getInt(), getInt()
		v[i] = 2*a + b
		x -= a
	}
	sort.Sort(sort.Reverse(sort.IntSlice(v)))

	ans := 0
	for x <= 0 {
		x += v[ans]
		ans++
	}
	puts(ans)
}
