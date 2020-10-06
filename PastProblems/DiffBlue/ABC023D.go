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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	h, s := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		h[i], s[i] = nextInt(), nextInt()
	}

	ok := func(x int) bool {
		// 制限時間を求める
		lims := make([]int, n)
		for i := 0; i < n; i++ {
			if h[i] > x {
				return false
			}
			lims[i] = (x - h[i]) / s[i]
		}
		sort.Sort(sort.IntSlice(lims))
		// 制限時間が短い順に風船を割る
		for i, lim := range lims {
			// 時刻が制限時間を上回るとNG
			if i > lim {
				return false
			}
		}
		return true
	}

	// highの最小値を求める
	low, high := 0, 1000000000*1000000000
	for low < high-1 {
		mid := (low + high) / 2
		if ok(mid) {
			high = mid
		} else {
			low = mid
		}
	}

	puts(high)
}
