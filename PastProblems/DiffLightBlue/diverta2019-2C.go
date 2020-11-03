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

// key以上の要素の最小インデックス
func lowerBound(a []int, key int) int {
	l, r := 0, len(a)-1
	for l <= r {
		mid := (l + r) / 2
		if a[mid] < key {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

type xy struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a := nextInts(n)
	sort.Sort(sort.IntSlice(a))

	if n == 2 {
		d := a[1] - a[0]
		if d > -d {
			puts(d)
			puts(a[1], a[0])
		} else {
			puts(-d)
			puts(a[0], a[1])
		}
		return
	}

	// 0未満の要素の個数
	q := lowerBound(a, 0)
	// 0以上の要素の個数
	p := n - q
	if p == 0 {
		p, q = 1, n-1
	} else if q == 0 {
		p, q = n-1, 1
	}

	seq := make([]xy, 0)
	for i := q; i < n-1; i++ {
		x, y := a[0], a[i]
		a[0] = x - y
		seq = append(seq, xy{x, y})
	}
	for i := 0; i < q; i++ {
		x, y := a[n-1], a[i]
		a[n-1] = x - y
		seq = append(seq, xy{x, y})
	}

	puts(a[n-1])
	for i := range seq {
		puts(seq[i].x, seq[i].y)
	}
}
