package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// 1-indexed
type BIT [2][]int

func NewBIT(n int) *BIT {
	var b BIT
	for i := range b {
		b[i] = make([]int, n)
	}
	return &b
}

// [l,r)にxを加算
func (b *BIT) add(l, r, x int) {
	add_sub := func(p, idx, x int) {
		for idx < len(b[p]) {
			b[p][idx] += x
			idx += idx & (-idx)
		}
	}
	add_sub(0, l, -x*(l-1))
	add_sub(0, r, x*(r-1))
	add_sub(1, l, x)
	add_sub(1, r, -x)
}

// [l,r)の和
func (b *BIT) sum(l, r int) int {
	sum_sub := func(p, idx int) int {
		s := 0
		for idx > 0 {
			s += b[p][idx]
			idx -= idx & (-idx)
		}
		return s
	}
	return sum_sub(0, r-1) + sum_sub(1, r-1)*(r-1) - (sum_sub(0, l-1) + sum_sub(1, l-1)*(l-1))
}

//
// example: ABC153F
//
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

type Mon struct {
	x, h int
}

type Mons []Mon

func (a Mons) Len() int      { return len(a) }
func (a Mons) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Mons) Less(i, j int) bool {
	return a[i].x < a[j].x
}

func upperBound(a []int, key int) int {
	l, r := 0, len(a)-1
	for l <= r {
		mid := (l + r) / 2
		if a[mid] <= key {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, d, a := getInt(), getInt(), getInt()
	mons := make(Mons, n)
	for i := 0; i < n; i++ {
		x, h := getInt(), getInt()
		mons[i] = Mon{x, h}
	}
	sort.Sort(mons)
	x, h := make([]int, n), make([]int, n)
	for i := range mons {
		x[i], h[i] = mons[i].x, mons[i].h
	}

	bit := NewBIT(n + 10)
	for i := 0; i < n; i++ {
		bit.add(i+1, i+2, h[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		cur := bit.sum(i+1, i+2)
		if cur <= 0 {
			continue
		}
		need := (cur + a - 1) / a
		r := upperBound(x, x[i]+2*d)
		bit.add(i+1, r+1, -need*a)
		ans += need
	}

	puts(ans)
}
