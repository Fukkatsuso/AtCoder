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

type Query struct {
	l, r, idx int
}

type Queries []Query

func (a Queries) Len() int           { return len(a) }
func (a Queries) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Queries) Less(i, j int) bool { return a[i].r < a[j].r }

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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, q := nextInt(), nextInt()
	c := nextInts(n)
	queries := make(Queries, q)
	for i := 0; i < q; i++ {
		l, r := nextInt(), nextInt()
		queries[i] = Query{l, r, i}
	}
	sort.Sort(queries)

	// right[i]: 色iの玉のうち最も右にある玉=良い玉の場所
	right := make([]int, n+1)
	// 良い玉のある場所:1, それ以外:0
	// とし，区間和をとる
	bit := NewBIT(n + 10)
	ans := make([]int, q)
	// 1-indexed
	i := 1
	for _, query := range queries {
		for i <= query.r {
			col := c[i-1]
			if right[col] > 0 {
				bit.add(right[col], right[col]+1, -1)
			}
			right[col] = i
			bit.add(i, i+1, 1)
			i++
		}
		ans[query.idx] = bit.sum(query.l, query.r+1)
	}

	for i := range ans {
		puts(ans[i])
	}
}
