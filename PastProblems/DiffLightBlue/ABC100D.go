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

func sum(a ...int) int {
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

	n, m := nextInt(), nextInt()
	x, y, z := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i], z[i] = nextInt(), nextInt(), nextInt()
	}

	ans := 0
	sums := make([]int, n)
	for bit := 0; bit < (1 << 3); bit++ {
		for i := 0; i < n; i++ {
			sum := 0
			sum += x[i] * (1 - 2*((bit>>0)&1))
			sum += y[i] * (1 - 2*((bit>>1)&1))
			sum += z[i] * (1 - 2*((bit>>2)&1))
			sums[i] = sum
		}
		sort.Sort(sort.Reverse(sort.IntSlice(sums)))
		ans = max(ans, sum(sums[:m]...))
	}

	puts(ans)
}
