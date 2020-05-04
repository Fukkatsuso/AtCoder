// ABC031C

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

func point(a []int, b, c int) (int, int) {
	x, y := 0, 0
	for i := b; i <= c; i++ {
		if (i-b)%2 == 0 {
			x += a[i]
		} else {
			y += a[i]
		}
	}
	return x, y
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a := nextInts(n)

	ans := -5000
	for i := 0; i < n; i++ {
		taka, aoki := -5000, -5000
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			var x, y int
			if i < j {
				x, y = point(a, i, j)
			} else {
				x, y = point(a, j, i)
			}
			if y > aoki {
				taka, aoki = x, y
			}
		}
		ans = max(ans, taka)
	}

	puts(ans)
}
