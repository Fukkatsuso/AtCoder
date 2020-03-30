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

func sum(a []int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
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

	n, k, m, r := nextInt(), nextInt(), nextInt(), nextInt()
	s := nextInts(n - 1)
	sort.Sort(sort.IntSlice(s))

	need := r * k
	// 最終試験も強制的に点数に含まれる
	if k == n {
		if d := max(0, need-sum(s)); d > m {
			puts(-1)
		} else {
			puts(d)
		}
		return
	}
	// 0でも通る
	if sum(s[n-1-k:]) >= need {
		puts(0)
		return
	}
	// 最高点を取ってもダメ
	if sum(s[n-1-k+1:])+m < need {
		puts(-1)
		return
	}
	puts(max(0, need-sum(s[n-1-k+1:])))
}
