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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, a, b := getInt(), getInt(), getInt()
	v := getInts(n)
	sort.Sort(sort.Reverse(sort.IntSlice(v)))

	puts(float64(sum(v[:a]...)) / float64(a))

	com := [51][51]int{}
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				com[i][j] = 1
			} else {
				com[i][j] = com[i-1][j-1] + com[i-1][j]
			}
		}
	}

	// cnt[i]: 0~iまでのvでv[a-1]と同じものの個数
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		if v[i] == v[a-1] {
			cnt[i]++
		}
		if i > 0 {
			cnt[i] += cnt[i-1]
		}
	}

	var ans int
	if v[0] != v[a-1] {
		ans = com[cnt[n-1]][cnt[a-1]]
	} else {
		for i := a; i <= b; i++ {
			ans += com[cnt[n-1]][i]
		}
	}
	puts(ans)
}
