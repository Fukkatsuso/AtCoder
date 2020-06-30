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
	p, q := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = nextInt()
		q[i] = p[i]
	}
	sort.Sort(sort.IntSlice(q))

	d := 0
	for i := range p {
		if p[i] != q[i] {
			d++
		}
	}

	if d <= 2 {
		puts("YES")
	} else {
		puts("NO")
	}
}
