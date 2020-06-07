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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

type Edge struct {
	x, y int
}

type Edges []Edge

func (a Edges) Len() int      { return len(a) }
func (a Edges) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Edges) Less(i, j int) bool {
	if a[i].y == a[j].y {
		return a[i].x < a[j].x
	}
	return a[i].y < a[j].y
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const inf = 1000000

	_, m := nextInt(), nextInt()
	edges := make(Edges, m)
	for i := 0; i < m; i++ {
		edges[i] = Edge{nextInt() - 1, nextInt() - 1}
	}
	sort.Sort(edges)

	ans := 0
	end := 0
	for i := 0; i < m; i++ {
		if edges[i].x >= end {
			end = edges[i].y
			ans++
		}
	}

	puts(ans)
}
