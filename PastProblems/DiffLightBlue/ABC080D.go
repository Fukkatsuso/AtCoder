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

	const MAX = 100000 + 2

	N, C := nextInt(), nextInt()
	pp := make([][MAX]int, C)
	for i := 0; i < N; i++ {
		s, t, c := nextInt(), nextInt(), nextInt()-1
		pp[c][s]++
		pp[c][t+1]--
	}

	ans := 0
	for i := 1; i < MAX; i++ {
		a := 0
		for j := range pp {
			pp[j][i] += pp[j][i-1]
			if pp[j][i] > 0 {
				a++
			}
		}
		ans = max(ans, a)
	}

	puts(ans)
}
