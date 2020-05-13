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

func nextInt2() (int, int) {
	return nextInt(), nextInt()
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

	n, m := nextInt2()
	rel := make([][]bool, n)
	for i := range rel {
		rel[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		x, y := nextInt2()
		x--
		y--
		rel[x][y], rel[y][x] = true, true
	}

	ans := 0
	for bit := 1; bit < (1 << uint(n)); bit++ {
		group := make([]int, 0)
		ok := true
		for i := 0; ok && i < n; i++ {
			if (bit>>uint(i))&1 == 1 {
				for _, j := range group {
					ok = ok && rel[i][j]
				}
				group = append(group, i)
			}
		}
		if ok {
			ans = max(ans, len(group))
		}
	}

	puts(ans)
}
