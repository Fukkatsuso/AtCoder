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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = nextInt(), nextInt()
	}
	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.IntSlice(b))

	var ans int
	if n%2 == 1 {
		ca, cb := a[n/2], b[n/2]
		ans = cb - ca + 1
	} else {
		ca, cb := a[(n-1)/2]+a[n/2], b[(n-1)/2]+b[n/2]
		ans = cb - ca + 1
	}

	puts(ans)
}
