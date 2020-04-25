// 解説ブログAC

package main

import (
	"bufio"
	"fmt"
	"os"
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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	r := nextInts(n)

	graph := func(g []int, up bool) []int {
		for i, l := 1, 1; i < n; i++ {
			if up {
				if g[l-1] < r[i] {
					g = append(g, r[i])
					l++
					up = false
				} else {
					g[l-1] = r[i]
				}
			} else {
				if g[l-1] > r[i] {
					g = append(g, r[i])
					l++
					up = true
				} else {
					g[l-1] = r[i]
				}
			}
		}
		return g
	}

	// 上下
	ud := make([]int, 0)
	ud = append(ud, r[0])
	ud = graph(ud, true)
	// puts(ud)

	// 下上
	du := make([]int, 0)
	du = append(du, r[0])
	du = graph(du, false)
	// puts(du)

	ans := max(len(ud), len(du))
	if ans < 3 {
		ans = 0
	}
	puts(ans)
}
