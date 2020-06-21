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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func divFloor(a, b int) int {
	return (a + (b - 1)) / b
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

	n, d, a := nextInt(), nextInt(), nextInt()
	mons := make(Mons, n)
	for i := 0; i < n; i++ {
		x, h := nextInt(), nextInt()
		mons[i] = Mon{x, h}
	}
	sort.Sort(mons)
	x, h := make([]int, n), make([]int, n)
	for i := range mons {
		x[i], h[i] = mons[i].x, mons[i].h
	}

	ans := 0
	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		if s[i] < h[i] {
			// 攻撃回数
			need := (h[i] - s[i] + a - 1) / a
			// 攻撃範囲
			r := upperBound(x, x[i]+2*d)
			s[i] += need * a
			s[r] -= need * a
			ans += need
		}
		s[i+1] += s[i]
	}

	puts(ans)
}
