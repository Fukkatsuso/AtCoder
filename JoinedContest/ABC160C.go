package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	k, n := nextInt(), nextInt()
	a := nextInts(n)
	sort.Sort(sort.IntSlice(a))

	ans := k + 1
	for i := 0; i < n; i++ {
		d := min(k-(a[i]-a[(i-1+n)%n]+k)%k, k-(a[(i+1)%n]-a[i]))
		ans = min(ans, d)
	}

	puts(ans)
}

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
