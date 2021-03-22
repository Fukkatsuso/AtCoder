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

func getInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = getInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func sum(a ...int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

type Item struct {
	w, v int
}

type Box struct {
	size, id int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m, q := getInt(), getInt(), getInt()

	items := make([]Item, n)
	for i := 0; i < n; i++ {
		w, v := getInt(), getInt()
		items[i] = Item{w, v}
	}
	// 価値が高くて小さいもの順
	sort.Slice(items, func(i, j int) bool {
		if items[i].v == items[j].v {
			return items[i].w < items[j].w
		}
		return items[i].v > items[j].v
	})

	x := make([]Box, m)
	for i := 0; i < m; i++ {
		x[i] = Box{getInt(), i}
	}
	// 小さいもの順
	sort.Slice(x, func(i, j int) bool {
		return x[i].size < x[j].size
	})

	maxVal := func(l, r int) int {
		box := make([]int, m)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if l <= x[j].id && x[j].id <= r {
					continue
				}
				if box[j] == 0 && x[j].size >= items[i].w {
					box[j] = items[i].v
					break
				}
			}
		}
		val := sum(box...)
		return val
	}

	for i := 0; i < q; i++ {
		l, r := getInt()-1, getInt()-1
		puts(maxVal(l, r))
	}
}
