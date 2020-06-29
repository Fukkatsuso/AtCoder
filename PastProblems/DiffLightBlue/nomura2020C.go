// 解説AC

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

	n := nextInt()
	a := nextInts(n + 1)

	l, r := make([]int, n+1), make([]int, n+1)
	l[n], r[n] = a[n], a[n]
	for i := n - 1; i >= 0; i-- {
		l[i] = a[i] + (l[i+1]+1)/2
		r[i] = a[i] + r[i+1]
	}
	if l[0] > 1 {
		puts(-1)
		return
	}

	ans, cur := 0, 1
	for i := 0; i <= n; i++ {
		ans += cur
		if i < n {
			cur = min((cur-a[i])*2, r[i+1])
		}
	}
	puts(ans)
}
