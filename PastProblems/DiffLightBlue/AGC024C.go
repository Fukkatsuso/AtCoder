// いくつかWAで解説AC

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
	a := nextInts(n)

	if a[0] > 0 {
		puts(-1)
		return
	}

	ans := 0
	b := make([]int, n)
	for i := n - 1; i > 0; i-- {
		if a[i] == b[i] {
			b[i-1] = b[i] - 1
			continue
		}
		if a[i] > i || b[i] > a[i] {
			puts(-1)
			return
		}
		ans += a[i]
		b[i-1] = a[i] - 1
	}

	puts(ans)
}
