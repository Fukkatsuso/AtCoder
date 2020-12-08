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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	p := make([]int, n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = getInt() - 1
		d[i] = p[i] - i
	}

	// pm[i]: d[pm[i]]>0 かつ d[pm[i]+1]<0
	pm := make([]int, 0)
	for i := 0; i < n-1; i++ {
		if d[i] > 0 && d[i+1] < 0 {
			pm = append(pm, i)
		}
	}

	ans := make([]int, 0)
	used := make([]bool, n)
	for len(pm) > 0 {
		idx := pm[0]
		pm = pm[1:]
		if d[idx] == 0 && d[idx+1] == 0 {
			continue
		}
		if d[idx]*d[idx+1] == 0 {
			puts(-1)
			return
		}
		if d[idx]*d[idx+1] > 0 {
			continue
		}
		if used[idx] {
			puts(-1)
			return
		}
		ans = append(ans, idx+1)
		used[idx] = true
		p[idx], p[idx+1] = p[idx+1], p[idx]
		d[idx], d[idx+1] = d[idx+1]+1, d[idx]-1
		if idx > 0 && d[idx-1] > 0 && d[idx] < 0 {
			pm = append(pm, idx-1)
		}
		if idx < n-2 && d[idx+1] > 0 && d[idx+2] < 0 {
			pm = append(pm, idx+1)
		}
	}

	for i := 0; i < n-1; i++ {
		if p[i] > p[i+1] {
			puts(-1)
			return
		}
		if !used[i] {
			puts(-1)
			return
		}
	}

	for i := range ans {
		puts(ans[i])
	}
}
